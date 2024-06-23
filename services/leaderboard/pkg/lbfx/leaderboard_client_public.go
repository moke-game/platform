package lbfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
	"go.uber.org/fx"

	"github.com/moke-game/platform/api/gen/leaderboard"
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
	if sSetting.MTLSEnable {
		if conn, err := tools.DialWithSecurity(
			host,
			sSetting.ClientCert,
			sSetting.ClientKey,
			sSetting.ServerName,
			sSetting.ServerCaCert,
		); err != nil {
			return nil, err
		} else {
			return leaderboard.NewLeaderboardServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return leaderboard.NewLeaderboardServiceClient(conn), nil
		}
	}
}

var LeaderboardClientModule = fx.Provide(
	func(
		setting LeaderboardSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out LeaderboardClientResult, err error) {
		if cli, e := CreateLeaderboardClient(setting.Url, sSetting); e != nil {
			err = e
		} else {
			out.Client = cli
		}
		return
	},
)
