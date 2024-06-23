package ptfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/moke-game/platform.git/api/gen/party"
)

type PartyClientParams struct {
	fx.In

	PartyClient pb.PartyServiceClient `name:"PartyClient"`
}

type PartyClientResult struct {
	fx.Out

	PartyClient pb.PartyServiceClient `name:"PartyClient"`
}

func NewPartyClient(host string, sSetting sfx.SecuritySettingsParams) (pb.PartyServiceClient, error) {
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
			return pb.NewPartyServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewPartyServiceClient(conn), nil
		}
	}
}

var PartyClientModule = fx.Provide(
	func(
		setting PartySettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out PartyClientResult, err error) {
		if cli, e := NewPartyClient(setting.PartyUrl, sSetting); e != nil {
			err = e
		} else {
			out.PartyClient = cli
		}
		return
	},
)
