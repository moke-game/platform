package module

import (
	agones "github.com/gstones/moke-kit/3rd/agones/pkg/module"
	awsConfig "github.com/gstones/moke-kit/3rd/cloud/pkg/module"
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/matchmaking/internal"
	"github.com/moke-game/platform/services/matchmaking/pkg/mmfx"
)

var settings = mmfx.MatchmakingSettingsModule

var MatchmakingModule = fx.Module("matchmaking",
	settings,
	agones.AgonesAllocateClientModule,
	awsConfig.AWSConfigModule,
	internal.Module,
)

var MatchmakingClientModule = fx.Module("matchmaking_client",
	settings,
	mmfx.ClientModule,
)
