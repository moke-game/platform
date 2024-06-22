package afx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type AuthSettingParams struct {
	fx.In

	AuthUrl        string `name:"AuthUrl"`
	AuthStoreName  string `name:"AuthStoreName"`
	JwtTokenSecret string `name:"JwtTokenSecret"`
	JwtTokenExpire int32  `name:"JwtTokenExpire"`
}

type AuthSettingsResult struct {
	fx.Out

	AuthStoreName  string `name:"AuthStoreName" envconfig:"AUTH_STORE_NAME" default:"auth"`
	AuthUrl        string `name:"AuthUrl" envconfig:"AUTH_URL" default:"localhost:8081"`
	JwtTokenSecret string `name:"JwtTokenSecret" default:"" envconfig:"JWT_TOKEN_SECRET"`
	// JwtTokenExpire (hours)
	JwtTokenExpire int32 `name:"JwtTokenExpire" default:"12" envconfig:"JWT_TOKEN_EXPIRE"`
}

func (g *AuthSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out AuthSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
