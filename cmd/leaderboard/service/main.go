package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/gstones/platform/services/auth/pkg/module"
	"github.com/gstones/platform/services/leaderboard/pkg/module"
)

func main() {
	fxmain.Main(
		ofx.RedisCacheModule,
		module.LeaderboardService,
		auth.AuthMiddlewareModule,
	)
}
