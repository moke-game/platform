package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	auth "github.com/gstones/platform/services/auth/pkg/module"
	profile "github.com/gstones/platform/services/profile/pkg/module"
)

func main() {
	fxmain.Main(
		ofx.RedisCacheModule,
		profile.ProfileModule,
		mfx.NatsModule,
		auth.AuthMiddlewareModule,
	)
}
