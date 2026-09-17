package main

import (
	"github.com/gstones/moke-kit/fxmain"

	analytics "github.com/moke-game/platform/services/analytics/pkg/module"
	auth "github.com/moke-game/platform/services/auth/pkg/module"
)

func main() {
	fxmain.Main(
		analytics.AnalyticsModule,
		auth.PrivateServiceAuthModule, // pass-through; analytics is WithoutAuth-only
	)
}
