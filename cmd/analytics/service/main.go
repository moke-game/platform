package main

import (
	"github.com/gstones/moke-kit/fxmain"

	"github.com/moke-game/platform/services/analytics/pkg/module"
)

func main() {
	fxmain.Main(
		module.AnalyticsModule,
	)
}
