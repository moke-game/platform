package pfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/profile/api"
	"github.com/moke-game/platform/pkg/platformfx"
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
	conn, err := platformfx.Dial(host, sSetting)
	if err != nil {
		return nil, nil, err
	}
	return pb.NewProfileServiceClient(conn), pb.NewProfilePrivateServiceClient(conn), nil
}

var ProfileClientModule = fx.Provide(
	func(
		setting ProfileSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out ProfileClientResult, err error) {
		out.ProfileClient, out.ProfilePrivateClient, err = NewProfileClient(setting.ProfileUrl, sSetting)
		return
	},
)
