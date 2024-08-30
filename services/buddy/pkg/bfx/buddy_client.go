package bfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/buddy/api"
)

type BuddyClientParams struct {
	fx.In

	BuddyClient pb.BuddyServiceClient `name:"BuddyClient"`
}

type BuddyClientResult struct {
	fx.Out

	BuddyClient pb.BuddyServiceClient `name:"BuddyClient"`
}

func NewBuddyClient(host string, sSetting sfx.SecuritySettingsParams) (pb.BuddyServiceClient, error) {
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
			return pb.NewBuddyServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewBuddyServiceClient(conn), nil
		}
	}
}

var BuddyClientModule = fx.Provide(
	func(
		setting BuddySettingsParams,
		sSetting sfx.SecuritySettingsParams,
	) (out BuddyClientResult, err error) {
		if cli, e := NewBuddyClient(setting.BuddyUrl, sSetting); e != nil {
			err = e
		} else {
			out.BuddyClient = cli
		}
		return
	},
)
