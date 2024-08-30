package afx

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/utility"
	"github.com/supabase-community/supabase-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// supabase auth middleware
// https://supabase.com/docs/guides/auth

// SupabaseAuthor is auth for grpc middleware
type SupabaseAuthor struct {
	client        *supabase.Client
	unAuthMethods map[string]struct{}
}

// Auth will auth every grpc request with supabase
func (d *SupabaseAuthor) Auth(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)
	if _, ok := d.unAuthMethods[method]; ok {
		return context.WithValue(ctx, utility.WithOutTag, true), nil
	} else if token, err := auth.AuthFromMD(ctx, string(utility.TokenContextKey)); err != nil {
		return ctx, err
	} else if resp, err := d.client.Auth.WithToken(token).GetUser(); err != nil {
		return ctx, err
	} else {
		ctx = context.WithValue(ctx, utility.UIDContextKey, resp.ID.String())
		return ctx, nil
	}
}

// AddUnAuthMethod add unauth method
func (d *SupabaseAuthor) AddUnAuthMethod(method string) {
	if d.unAuthMethods == nil {
		d.unAuthMethods = make(map[string]struct{})
	}
	d.unAuthMethods[method] = struct{}{}
}

// SupabaseCheckModule is the supabase Auth module for grpc middleware
var SupabaseCheckModule = fx.Provide(
	func(
		l *zap.Logger,
		sSetting SupabaseSettingParams,
	) (out sfx.AuthMiddlewareResult, err error) {
		c, err := supabase.NewClient(sSetting.URL, sSetting.Key, nil)
		if err != nil {
			return
		}
		out.AuthMiddleware = &SupabaseAuthor{
			client:        c,
			unAuthMethods: map[string]struct{}{},
		}
		return
	},
)
