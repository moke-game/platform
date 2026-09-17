package kfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/knapsack/api"
	"github.com/moke-game/platform/pkg/platformfx"
)

type KnapsackClientParams struct {
	fx.In

	KnapsackClient        pb.KnapsackServiceClient        `name:"KnapsackClient"`
	KnapsackPrivateClient pb.KnapsackPrivateServiceClient `name:"KnapsackPrivateClient"`
}

type KnapsackClientResult struct {
	fx.Out

	KnapsackClient        pb.KnapsackServiceClient        `name:"KnapsackClient"`
	KnapsackPrivateClient pb.KnapsackPrivateServiceClient `name:"KnapsackPrivateClient"`
}

func NewKnapsackClient(host string, sSetting sfx.SecuritySettingsParams) (pb.KnapsackServiceClient, error) {
	return platformfx.NewClient(host, sSetting, pb.NewKnapsackServiceClient)
}

func NewKnapsackPrivateClient(host string, sSetting sfx.SecuritySettingsParams) (pb.KnapsackPrivateServiceClient, error) {
	return platformfx.NewClient(host, sSetting, pb.NewKnapsackPrivateServiceClient)
}

var KnapsackClientModule = fx.Provide(
	func(
		setting KnapsackSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out KnapsackClientResult, err error) {
		conn, e := platformfx.Dial(setting.KnapsackUrl, sSetting)
		if e != nil {
			err = e
			return
		}
		out.KnapsackClient = pb.NewKnapsackServiceClient(conn)
		out.KnapsackPrivateClient = pb.NewKnapsackPrivateServiceClient(conn)
		return
	},
)
