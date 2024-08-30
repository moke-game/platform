package afx

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/moke-game/platform/api/gen/auth/api"
)

// Author is auth for grpc middleware
type Author struct {
	client        pb.AuthServiceClient
	unAuthMethods map[string]struct{}
}

// Auth will  auth  every grpc request
func (d *Author) Auth(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)
	if _, ok := d.unAuthMethods[method]; ok {
		return context.WithValue(ctx, utility.WithOutTag, true), nil
	} else if token, err := auth.AuthFromMD(ctx, string(utility.TokenContextKey)); err != nil {
		return ctx, err
	} else if resp, err := d.client.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	}); err != nil {
		return ctx, err
	} else {
		ctx = context.WithValue(ctx, utility.UIDContextKey, resp.GetUid())
		return ctx, nil
	}
}

func (d *Author) AddUnAuthMethod(method string) {
	if d.unAuthMethods == nil {
		d.unAuthMethods = make(map[string]struct{})
	}
	d.unAuthMethods[method] = struct{}{}
}

// AuthCheckModule is the module for grpc middleware
var AuthCheckModule = fx.Provide(
	func(
		l *zap.Logger,
		sSetting sfx.SecuritySettingsParams,
		params AuthClientParams,
	) (out sfx.AuthMiddlewareResult, err error) {
		out.AuthMiddleware = &Author{
			client:        params.AuthClient,
			unAuthMethods: map[string]struct{}{},
		}
		return
	},
)
