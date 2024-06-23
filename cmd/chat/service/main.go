package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"

	auth "github.com/moke-game/platform.git/services/auth/pkg/module"

	"github.com/moke-game/platform.git/services/chat/pkg/module"
)

func main() {
	fxmain.Main(
		mfx.NatsModule,
		module.ChatModule,
		auth.AuthMiddlewareModule,
	)
}
