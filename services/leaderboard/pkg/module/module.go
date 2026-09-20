package module

import (
	"go.uber.org/fx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
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

// LeaderboardClientModule is the catalog name used by pkg/assembly.
var LeaderboardClientModule = LeaderboardClientPublic

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

// LeaderboardAllModule is the catalog name used by pkg/assembly.
var LeaderboardAllModule = LeaderboardAll

// App is the standalone leaderboard process. Keep infra in sync with assembly recipe "leaderboard".
var App = fx.Module("leaderboard_app",
	auth.AuthMiddlewareModule,
	LeaderboardModule,
)
