package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	"github.com/gstones/platform/services/auth/pkg/module"
)

func main() {
	fxmain.Main(
		ofx.RedisCacheModule,
		module.AuthModule,
	)
}
