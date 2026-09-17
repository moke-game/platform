package bfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/buddy/api"
	"github.com/moke-game/platform/pkg/platformfx"
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
	return platformfx.NewClient(host, sSetting, pb.NewBuddyServiceClient)
}

var BuddyClientModule = fx.Provide(
	func(
		setting BuddySettingsParams,
		sSetting sfx.SecuritySettingsParams,
	) (out BuddyClientResult, err error) {
		out.BuddyClient, err = NewBuddyClient(setting.BuddyUrl, sSetting)
		return
	},
)
