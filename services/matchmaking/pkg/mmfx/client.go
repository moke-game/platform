package mmfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
	"go.uber.org/fx"

	matchmaking "github.com/moke-game/platform/api/gen/matchmaking/api"
)

type ClientParams struct {
	fx.In

	Client matchmaking.MatchServiceClient `name:"MatchServiceClient"`
}

type ClientResult struct {
	fx.Out

	Client matchmaking.MatchServiceClient `name:"MatchServiceClient"`
}

func NewClient(host string, setting sfx.SecuritySettingsParams) (matchmaking.MatchServiceClient, error) {
	if setting.MTLSEnable {
		if conn, err := tools.DialWithSecurity(
			host,
			setting.ClientCert,
			setting.ClientKey,
			setting.ServerName,
			setting.ServerCaCert,
		); err != nil {
			return nil, err
		} else {
			return matchmaking.NewMatchServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return matchmaking.NewMatchServiceClient(conn), nil
		}
	}
}

var ClientModule = fx.Provide(
	func(
		setting MatchmakingSettingParams,
		security sfx.SecuritySettingsParams,
	) (out ClientResult, err error) {
		if client, e := NewClient(setting.URL, security); e != nil {
			err = e
		} else {
			out.Client = client
		}
		return
	})
