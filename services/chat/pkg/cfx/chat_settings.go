package cfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
)

type ChatSettingParams struct {
	fx.In
	Name         string `name:"ChatName"`
	ChatUrl      string `name:"ChatUrl"`
	ChatInterval int    `name:"ChatInterval"`
}

type ChatSettingResult struct {
	fx.Out
	Name         string `name:"ChatName" envconfig:"CHAT_NAME" default:"chat"`
	ChatUrl      string `name:"ChatUrl" envconfig:"CHAT_URL" default:"localhost:8081"`
	ChatInterval int    `name:"ChatInterval" envconfig:"WORLD_CHAT_INTERVAL" default:"2"`
}

var ChatSettingsModule = platformfx.ProvideFromEnv[ChatSettingResult]()
