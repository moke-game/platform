package main

import (
	"github.com/gstones/moke-kit/fxmain"

	analytics "github.com/moke-game/platform/services/analytics/pkg/module"
	auth "github.com/moke-game/platform/services/auth/pkg/module"
)

func main() {
	fxmain.Main(
		analytics.AnalyticsModule,
		// Analytics embeds WithoutAuth; pass-through middleware satisfies kit
		// binder prod checks (#224+) without calling ValidateToken.
		auth.PrivateServiceAuthModule,
	)
}
