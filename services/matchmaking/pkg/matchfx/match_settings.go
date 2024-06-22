package matchfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/utility"
)

type MatchSettingParams struct {
	fx.In

	MatchUrl       string `name:"MatchUrl"`
	MatchStoreName string `name:"MatchStoreName"`

	AWSKey    string `name:"AWSKey"`
	AWSRegion string `name:"AWSRegion"`
	AWSSecret string `name:"AWSSecret"`
	SubNet    string `name:"SubNet"`
}

type MatchSettingsResult struct {
	fx.Out

	MatchUrl       string `name:"MatchUrl" envconfig:"MATCH_URL" default:"localhost:8081"`
	MatchStoreName string `name:"MatchStoreName" envconfig:"MATCH_STORE_NAME" default:"matchmaking"`

	AWSKey    string `name:"AWSKey" envconfig:"AWS_KEY" default:""`
	AWSRegion string `name:"AWSRegion" envconfig:"AWS_REGION" default:"us-west-2"`
	AWSSecret string `name:"AWSSecret" envconfig:"AWS_SECRET" default:""`
	SubNet    string `name:"SubNet" envconfig:"SUB_NET" default:""`
}

func (m *MatchSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(m)
	return
}

var SettingsModule = fx.Provide(
	func() (out MatchSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
