package pfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
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

var SettingsModule = platformfx.ProvideFromEnv[ProfileSettingsResult]()
