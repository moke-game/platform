package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/moke-game/platform.git/services/auth/pkg/module"
	"github.com/moke-game/platform.git/services/buddy/pkg/module"
)

func main() {
	fxmain.Main(
		module.BuddyModule,
		ofx.RedisCacheModule,
		mfx.NatsModule,
		auth.AuthMiddlewareModule,
	)
}
