package module

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/3rd/agones/pkg/module"

	"github.com/moke-game/platform.git/services/matchmaking/internal/service"
	"github.com/moke-game/platform.git/services/matchmaking/pkg/matchfx"
)

// MatchModule Provides match service
var MatchModule = fx.Module("match",
	service.MatchService,
	matchfx.SettingsModule,
	module.AgonesAllocateClientModule,
)

// MatchClientModule Provides match client for grpc
var MatchClientModule = fx.Module("match_client",
	matchfx.MatchClientModule,
	matchfx.SettingsModule,
)

// MatchAllModule Provides client, service for match
var MatchAllModule = fx.Module("match_all",
	service.MatchService,
	matchfx.MatchClientModule,
	matchfx.SettingsModule,
	module.AgonesAllocateClientModule,
)
