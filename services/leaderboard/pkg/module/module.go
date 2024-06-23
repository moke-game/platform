package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/leaderboard/internal/service/private"
	"github.com/moke-game/platform/services/leaderboard/internal/service/public"
	"github.com/moke-game/platform/services/leaderboard/pkg/lbfx"
)

var LeaderboardService = fx.Module("leaderboard",
	lbfx.LeaderboardSettingModule,
	public.Module,
	private.Module,
)

var LeaderboardClientPublic = fx.Module("leaderboardClientPublic",
	lbfx.LeaderboardClientModule,
	lbfx.LeaderboardSettingModule,
)

var LeaderboardClientPrivate = fx.Module("leaderboardClientPrivate",
	lbfx.LeaderboardClientPrivateModule,
	lbfx.LeaderboardSettingModule,
)

var LeaderboardClientAll = fx.Module("leaderboardClientAll",
	lbfx.LeaderboardClientPrivateModule,
	lbfx.LeaderboardClientModule,
	lbfx.LeaderboardSettingModule,
)

var LeaderboardAll = fx.Module("leaderboardAll",
	public.Module,
	private.Module,
	lbfx.LeaderboardClientModule,
	lbfx.LeaderboardClientPrivateModule,
	lbfx.LeaderboardSettingModule,
)
