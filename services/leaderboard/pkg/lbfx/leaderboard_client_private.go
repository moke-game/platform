package lbfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
	"go.uber.org/fx"

	leaderboard "github.com/moke-game/platform/api/gen/leaderboard/api"
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
			return leaderboard.NewLeaderboardPrivateServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return leaderboard.NewLeaderboardPrivateServiceClient(conn), nil
		}
	}
}

var LeaderboardClientPrivateModule = fx.Provide(
	func(
		setting LeaderboardSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out LeaderboardClientPrivateResult, err error) {
		if cli, e := CreateLeaderboardPrivateClient(setting.Url, sSetting); e != nil {
			err = e
		} else {
			out.Client = cli
		}
		return
	},
)
