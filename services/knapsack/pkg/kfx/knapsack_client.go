package kfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/knapsack/api"
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
			return pb.NewKnapsackServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewKnapsackServiceClient(conn), nil
		}
	}
}

func NewKnapsackPrivateClient(host string, sSetting sfx.SecuritySettingsParams) (pb.KnapsackPrivateServiceClient, error) {
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
			return pb.NewKnapsackPrivateServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewKnapsackPrivateServiceClient(conn), nil
		}
	}
}

var KnapsackClientModule = fx.Provide(
	func(
		setting KnapsackSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out KnapsackClientResult, err error) {
		if cli, e := NewKnapsackClient(setting.KnapsackUrl, sSetting); e != nil {
			err = e
		} else {
			out.KnapsackClient = cli
		}
		if cli, e := NewKnapsackPrivateClient(setting.KnapsackUrl, sSetting); e != nil {
			err = e
		} else {
			out.KnapsackPrivateClient = cli
		}
		return
	},
)
