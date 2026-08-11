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
// public gRPC service is registered without AuthMiddleware.
//
// Private services that embed utility.WithoutAuth are exempt; they are expected
// to stay on internal networks only.
func CheckProdPublicAuthFailClosed(
	app mfx.AppParams,
	auth sfx.AuthMiddlewareParams,
	grpc sfx.GrpcServiceParams,
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
	return nil
}

// ProdAuthGuardModule runs the production public-auth startup check.
var ProdAuthGuardModule = fx.Invoke(CheckProdPublicAuthFailClosed)
