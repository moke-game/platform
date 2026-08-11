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

func hasPublicGrpc(grpc sfx.GrpcServiceParams) bool {
	for _, svc := range grpc.GrpcServices {
		if _, private := svc.(authFuncOverrider); !private {
			return true
		}
	}
	return false
}

func hasPublicGateway(gateway sfx.GatewayServiceParams) bool {
	for _, svc := range gateway.GatewayServices {
		if _, private := svc.(authFuncOverrider); !private {
			return true
		}
	}
	return false
}

// CheckProdPublicAuthFailClosed fails process startup in production when any
// public gRPC or gateway service is registered without real AuthMiddleware.
//
// Private services that embed utility.WithoutAuth are exempt (gRPC and gateway);
// they are expected to stay on internal networks only.
//
// Pass-through middleware (PrivateServiceAuthModule) is only allowed when every
// registered gRPC/gateway service is private (implements AuthFuncOverride).
//
// Note: this Invoke is wired into AuthMiddlewareModule / AuthAllModule /
// PrivateServiceAuthModule. If a main forgets those modules entirely, this check
// does not run — rely on moke-kit #221 request-path fail-closed, or kit #224
// binder startup check.
func CheckProdPublicAuthFailClosed(
	app mfx.AppParams,
	auth sfx.AuthMiddlewareParams,
	grpc sfx.GrpcServiceParams,
	gateway sfx.GatewayServiceParams,
) error {
	if !utility.ParseDeployments(app.Deployment).IsProd() {
		return nil
	}

	publicGrpc := hasPublicGrpc(grpc)
	publicGateway := hasPublicGateway(gateway)

	if _, passThrough := auth.AuthMiddleware.(*passThroughAuthor); passThrough {
		if publicGrpc || publicGateway {
			return fmt.Errorf(
				"prod: PrivateServiceAuthModule (pass-through) cannot be used with public gRPC or gateway services",
			)
		}
		return nil
	}

	if auth.AuthMiddleware != nil {
		return nil
	}
	if publicGrpc {
		for _, svc := range grpc.GrpcServices {
			if _, private := svc.(authFuncOverrider); private {
				continue
			}
			return fmt.Errorf(
				"prod: public gRPC service %T registered without AuthMiddleware (fail-closed)",
				svc,
			)
		}
	}
	if publicGateway {
		for _, svc := range gateway.GatewayServices {
			if _, private := svc.(authFuncOverrider); private {
				continue
			}
			return fmt.Errorf(
				"prod: public gateway service %T registered without AuthMiddleware (fail-closed)",
				svc,
			)
		}
	}
	return nil
}

// ProdAuthGuardModule runs the production public-auth startup check.
var ProdAuthGuardModule = fx.Invoke(CheckProdPublicAuthFailClosed)
