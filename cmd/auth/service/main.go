package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	"github.com/moke-game/platform/services/auth/pkg/module"
)

func main() {
	fxmain.Main(
		ofx.RedisCacheModule,
		// AuthAllModule = AuthService (WithoutAuth) + AuthMiddleware.
		// Middleware satisfies kit binder prod checks (#224+) while login RPCs
		// remain reachable via utility.WithoutAuth on the auth service.
		module.AuthAllModule,
	)
}
