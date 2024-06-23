package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform.git/services/analytics/internal/service"
	"github.com/moke-game/platform.git/services/analytics/pkg/analyfx"
)

// AnalyticsModule provides service for analytics
var AnalyticsModule = fx.Module("analytics",
	service.ServiceModule,
	analyfx.SettingsModule,
)

// AnalyticsClientModule provides client for analytics
var AnalyticsClientModule = fx.Module("analytics-client",
	analyfx.AnalyticsClientModule,
	analyfx.SettingsModule,
)

// AnalyticsAllModule provides client, service and middleware for analytics
var AnalyticsAllModule = fx.Module("analytics-all",
	service.ServiceModule,
	analyfx.AnalyticsClientModule,
	analyfx.SettingsModule,
)
