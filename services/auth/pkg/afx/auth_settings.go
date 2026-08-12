package afx

import (
	"fmt"

	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
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
	// JwtTokenExpire token lifetime in hours.
	// <=0 omits JWT exp and stores the redis auth token with no TTL (non-expiring).
	// Forbidden when DEPLOYMENT is prod.
	JwtTokenExpire int32 `name:"JwtTokenExpire" default:"12" envconfig:"JWT_TOKEN_EXPIRE"`
}

func (g *AuthSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

func validateJwtExpire(deployment string, expire int32) error {
	if expire > 0 {
		return nil
	}
	if utility.ParseDeployments(deployment).IsProd() {
		return fmt.Errorf("prod: JWT_TOKEN_EXPIRE must be > 0 (non-expiring tokens forbidden)")
	}
	return nil
}

var SettingsModule = fx.Provide(
	func(app mfx.AppParams) (out AuthSettingsResult, err error) {
		if err = out.LoadFromEnv(); err != nil {
			return
		}
		err = validateJwtExpire(app.Deployment, out.JwtTokenExpire)
		return
	},
)
