package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/moke-game/platform.git/services/auth/pkg/module"
	knapsack "github.com/moke-game/platform.git/services/knapsack/pkg/module"
)

func main() {
	fxmain.Main(
		ofx.RedisCacheModule,
		knapsack.KnapsackModule,
		mfx.NatsModule,
		auth.AuthMiddlewareModule,
	)
}
