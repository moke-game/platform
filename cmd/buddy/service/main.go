package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/gstones/platform/services/auth/pkg/module"
	"github.com/gstones/platform/services/buddy/pkg/module"
)

func main() {
	fxmain.Main(
		module.BuddyModule,
		ofx.RedisCacheModule,
		mfx.NatsModule,
		auth.AuthMiddlewareModule,
	)
}
