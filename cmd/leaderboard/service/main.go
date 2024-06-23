package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/moke-game/platform.git/services/auth/pkg/module"
	"github.com/moke-game/platform.git/services/leaderboard/pkg/module"
)

func main() {
	fxmain.Main(
		ofx.RedisCacheModule,
		module.LeaderboardService,
		auth.AuthMiddlewareModule,
	)
}
