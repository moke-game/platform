package matchfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/moke-game/platform.git/api/gen/matchmaking"
)

type MatchClientParams struct {
	fx.In
	MatchClient pb.MatchServiceClient `name:"MatchClient"`
}

type MatchClientResult struct {
	fx.Out

	MatchClient pb.MatchServiceClient `name:"MatchClient"`
}

func NewMatchClient(host string, sSetting sfx.SecuritySettingsParams) (pb.MatchServiceClient, error) {
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
			return pb.NewMatchServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewMatchServiceClient(conn), nil
		}
	}
}

var MatchClientModule = fx.Provide(
	func(
		setting MatchSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out MatchClientResult, err error) {
		if cli, e := NewMatchClient(setting.MatchUrl, sSetting); e != nil {
			err = e
		} else {
			out.MatchClient = cli
		}
		return
	},
)
