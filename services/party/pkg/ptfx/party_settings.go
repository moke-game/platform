package ptfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
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

var PartySettingsModule = platformfx.ProvideFromEnv[PartySettingResult]()
