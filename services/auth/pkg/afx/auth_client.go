package afx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/auth/api"
	"github.com/moke-game/platform/pkg/platformfx"
)

type AuthClientParams struct {
	fx.In

	AuthClient pb.AuthServiceClient `name:"AuthClient"`
}

type AuthClientResult struct {
	fx.Out

	AuthClient pb.AuthServiceClient `name:"AuthClient"`
}

func NewAuthClient(host string, sSetting sfx.SecuritySettingsParams) (pb.AuthServiceClient, error) {
	return platformfx.NewClient(host, sSetting, pb.NewAuthServiceClient)
}

var AuthClientModule = fx.Provide(
	func(
		setting AuthSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out AuthClientResult, err error) {
		out.AuthClient, err = NewAuthClient(setting.AuthUrl, sSetting)
		return
	},
)
