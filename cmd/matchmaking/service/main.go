package main

import (
	"github.com/gstones/moke-kit/fxmain"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
	mm "github.com/moke-game/platform/services/matchmaking/pkg/module"
)

func main() {
	fxmain.Main(
		// MatchService is client-facing (public); require AuthMiddleware.
		auth.AuthMiddlewareModule,
		mm.MatchmakingModule,
	)
}
