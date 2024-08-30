package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/matchmaking/internal"
	"github.com/moke-game/platform/services/matchmaking/pkg/mmfx"
)

var MatchmakingModule = fx.Module("matchmaking",
	mmfx.MatchmakingSettingModule,
	internal.Module,
)

var MatchmakingClientModule = fx.Module("matchmaking_client",
	mmfx.ClientModule,
	mmfx.MatchmakingSettingModule,
)
