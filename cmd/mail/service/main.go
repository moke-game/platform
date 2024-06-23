package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
	"github.com/moke-game/platform/services/mail/pkg/module"
)

func main() {
	fxmain.Main(
		mfx.NatsModule,
		module.MailModule,
		ofx.RedisCacheModule,
		auth.AuthMiddlewareModule,
	)
}
