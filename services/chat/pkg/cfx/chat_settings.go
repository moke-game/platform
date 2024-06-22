package cfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
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

func (l *ChatSettingResult) LoadFromEnv() (err error) {
	err = utility.Load(l)
	return
}

var ChatSettingsModule = fx.Provide(
	func() (out ChatSettingResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
