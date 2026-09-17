package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/leaderboard/internal/service/private"
	"github.com/moke-game/platform/services/leaderboard/internal/service/public"
	"github.com/moke-game/platform/services/leaderboard/pkg/lbfx"
)

var settings = lbfx.LeaderboardSettingsModule

var LeaderboardModule = fx.Module("leaderboard",
	settings,
	public.Module,
	private.Module,
)

var LeaderboardClientPublic = fx.Module("leaderboardClientPublic",
	settings,
	lbfx.LeaderboardClientModule,
)

var LeaderboardClientPrivate = fx.Module("leaderboardClientPrivate",
	settings,
	lbfx.LeaderboardClientPrivateModule,
)

var LeaderboardClientAll = fx.Module("leaderboardClientAll",
	settings,
	lbfx.LeaderboardClientModule,
	lbfx.LeaderboardClientPrivateModule,
)

var LeaderboardAll = fx.Module("leaderboardAll",
	settings,
	public.Module,
	private.Module,
	lbfx.LeaderboardClientModule,
	lbfx.LeaderboardClientPrivateModule,
)
