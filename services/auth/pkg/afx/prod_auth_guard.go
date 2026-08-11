package afx

import (
	"context"
	"fmt"

	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

// authFuncOverrider matches grpc-ecosystem AuthFuncOverride (e.g. utility.WithoutAuth).
type authFuncOverrider interface {
	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
}

// CheckProdPublicAuthFailClosed fails process startup in production when any
// public gRPC or gateway service is registered without AuthMiddleware.
//
// Private services that embed utility.WithoutAuth are exempt; they are expected
// to stay on internal networks only.
//
// Note: this Invoke is wired into AuthMiddlewareModule / AuthAllModule. If a
// main forgets those modules entirely, this check does not run — rely on
// moke-kit #221 request-path fail-closed, or kit #224 binder startup check.
func CheckProdPublicAuthFailClosed(
	app mfx.AppParams,
	auth sfx.AuthMiddlewareParams,
	grpc sfx.GrpcServiceParams,
	gateway sfx.GatewayServiceParams,
) error {
	if !utility.ParseDeployments(app.Deployment).IsProd() {
		return nil
	}
	if auth.AuthMiddleware != nil {
		return nil
	}
	for _, svc := range grpc.GrpcServices {
		if _, private := svc.(authFuncOverrider); private {
			continue
		}
		return fmt.Errorf(
			"prod: public gRPC service %T registered without AuthMiddleware (fail-closed)",
			svc,
		)
	}
	if len(gateway.GatewayServices) > 0 {
		return fmt.Errorf(
			"prod: gateway service registered without AuthMiddleware (fail-closed)",
		)
	}
	return nil
}

// ProdAuthGuardModule runs the production public-auth startup check.
var ProdAuthGuardModule = fx.Invoke(CheckProdPublicAuthFailClosed)
