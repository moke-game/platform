package pfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type ProfileSettingParams struct {
	fx.In

	ProfileUrl       string `name:"ProfileUrl"`
	ProfileStoreName string `name:"ProfileStoreName"`
}

type ProfileSettingsResult struct {
	fx.Out

	ProfileStoreName string `name:"ProfileStoreName" envconfig:"PROFILE_STORE_NAME" default:"profile"`
	ProfileUrl       string `name:"ProfileUrl" envconfig:"PROFILE_URL" default:"localhost:8081"`
}

func (g *ProfileSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out ProfileSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
