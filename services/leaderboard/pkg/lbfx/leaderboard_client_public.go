package lbfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	leaderboard "github.com/moke-game/platform/api/gen/leaderboard/api"
	"github.com/moke-game/platform/pkg/platformfx"
)

type LeaderboardClientParams struct {
	fx.In

	Client leaderboard.LeaderboardServiceClient `name:"LeaderboardClient"`
}

type LeaderboardClientResult struct {
	fx.Out

	Client leaderboard.LeaderboardServiceClient `name:"LeaderboardClient"`
}

func CreateLeaderboardClient(host string, sSetting sfx.SecuritySettingsParams) (leaderboard.LeaderboardServiceClient, error) {
	return platformfx.NewClient(host, sSetting, leaderboard.NewLeaderboardServiceClient)
}

var LeaderboardClientModule = fx.Provide(
	func(
		setting LeaderboardSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out LeaderboardClientResult, err error) {
		out.Client, err = CreateLeaderboardClient(setting.Url, sSetting)
		return
	},
)
