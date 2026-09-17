package ptfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/party/api"
	"github.com/moke-game/platform/pkg/platformfx"
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
	return platformfx.NewClient(host, sSetting, pb.NewPartyServiceClient)
}

var PartyClientModule = fx.Provide(
	func(
		setting PartySettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out PartyClientResult, err error) {
		out.PartyClient, err = NewPartyClient(setting.PartyUrl, sSetting)
		return
	},
)
