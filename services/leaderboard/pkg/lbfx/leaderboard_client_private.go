package lbfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	leaderboard "github.com/moke-game/platform/api/gen/leaderboard/api"
	"github.com/moke-game/platform/pkg/platformfx"
)

type LeaderboardClientPrivateParams struct {
	fx.In

	Client leaderboard.LeaderboardPrivateServiceClient `name:"LeaderboardClientPrivate"`
}

type LeaderboardClientPrivateResult struct {
	fx.Out

	Client leaderboard.LeaderboardPrivateServiceClient `name:"LeaderboardClientPrivate"`
}

func CreateLeaderboardPrivateClient(host string, sSetting sfx.SecuritySettingsParams) (leaderboard.LeaderboardPrivateServiceClient, error) {
	return platformfx.NewClient(host, sSetting, leaderboard.NewLeaderboardPrivateServiceClient)
}

var LeaderboardClientPrivateModule = fx.Provide(
	func(
		setting LeaderboardSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out LeaderboardClientPrivateResult, err error) {
		out.Client, err = CreateLeaderboardPrivateClient(setting.Url, sSetting)
		return
	},
)
