package afx

import (
	"context"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

// passThroughAuthor satisfies siface.IAuthMiddleware without validating tokens.
// Use only for processes that exclusively host utility.WithoutAuth services
// (e.g. analytics) so kit binder prod checks (#224+) see AuthMiddleware != nil.
type passThroughAuthor struct{}

func (p *passThroughAuthor) Auth(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, utility.WithOutTag, true), nil
}

func (p *passThroughAuthor) AddUnAuthMethod(string) {}

// PassThroughAuthModule provides a no-op AuthMiddleware for private-only processes.
var PassThroughAuthModule = fx.Provide(
	func() (out sfx.AuthMiddlewareResult, err error) {
		out.AuthMiddleware = &passThroughAuthor{}
		return
	},
)
