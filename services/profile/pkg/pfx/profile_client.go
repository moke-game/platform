package pfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/moke-game/platform.git/api/gen/profile"
)

type ProfileClientParams struct {
	fx.In

	ProfileClient        pb.ProfileServiceClient        `name:"ProfileClient"`
	ProfilePrivateClient pb.ProfilePrivateServiceClient `name:"ProfilePrivateClient"`
}

type ProfileClientResult struct {
	fx.Out

	ProfileClient        pb.ProfileServiceClient        `name:"ProfileClient"`
	ProfilePrivateClient pb.ProfilePrivateServiceClient `name:"ProfilePrivateClient"`
}

func NewProfileClient(
	host string,
	sSetting sfx.SecuritySettingsParams,
) (pb.ProfileServiceClient, pb.ProfilePrivateServiceClient, error) {
	if sSetting.MTLSEnable {
		if conn, err := tools.DialWithSecurity(
			host,
			sSetting.ClientCert,
			sSetting.ClientKey,
			sSetting.ServerName,
			sSetting.ServerCaCert,
		); err != nil {
			return nil, nil, err
		} else {
			return pb.NewProfileServiceClient(conn), pb.NewProfilePrivateServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, nil, err
		} else {
			return pb.NewProfileServiceClient(conn), pb.NewProfilePrivateServiceClient(conn), nil
		}
	}
}

var ProfileClientModule = fx.Provide(
	func(
		setting ProfileSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out ProfileClientResult, err error) {
		if cli, pCli, e := NewProfileClient(setting.ProfileUrl, sSetting); e != nil {
			err = e
		} else {
			out.ProfileClient = cli
			out.ProfilePrivateClient = pCli
		}
		return
	},
)
