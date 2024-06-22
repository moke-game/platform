package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"

	auth "github.com/gstones/platform/services/auth/pkg/module"

	"github.com/gstones/platform/services/chat/pkg/module"
)

func main() {
	fxmain.Main(
		mfx.NatsModule,
		module.ChatModule,
		auth.AuthMiddlewareModule,
	)
}
