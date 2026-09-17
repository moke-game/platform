package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/analytics/internal/service"
	"github.com/moke-game/platform/services/analytics/pkg/analyfx"
)

var settings = analyfx.SettingsModule

// AnalyticsModule provides service for analytics
var AnalyticsModule = fx.Module("analytics",
	settings,
	service.ServiceModule,
)

// AnalyticsClientModule provides client for analytics
var AnalyticsClientModule = fx.Module("analytics-client",
	settings,
	analyfx.AnalyticsClientModule,
)

// AnalyticsAllModule provides client and service for analytics
var AnalyticsAllModule = fx.Module("analytics-all",
	settings,
	service.ServiceModule,
	analyfx.AnalyticsClientModule,
)
