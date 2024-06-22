package main

import (
	"github.com/gstones/moke-kit/fxmain"

	"github.com/gstones/platform/services/analytics/pkg/module"
)

func main() {
	fxmain.Main(
		module.AnalyticsModule,
	)
}
