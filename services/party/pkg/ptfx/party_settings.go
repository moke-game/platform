package ptfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type PartySettingParams struct {
	fx.In
	Name     string `name:"PartyName"`
	PartyUrl string `name:"PartyUrl"`
}

type PartySettingResult struct {
	fx.Out
	Name     string `name:"PartyName" envconfig:"PARTY_NAME" default:"party"`
	PartyUrl string `name:"PartyUrl" envconfig:"PARTY_URL" default:"localhost:8081"`
}

func (l *PartySettingResult) LoadFromEnv() (err error) {
	err = utility.Load(l)
	return
}

var PartySettingsModule = fx.Provide(
	func() (out PartySettingResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
