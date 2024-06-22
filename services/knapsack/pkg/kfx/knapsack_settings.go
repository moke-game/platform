package kfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type KnapsackSettingParams struct {
	fx.In

	KnapsackUrl       string `name:"KnapsackUrl"`
	KnapsackStoreName string `name:"KnapsackStoreName"`
}

type KnapsackSettingsResult struct {
	fx.Out

	KnapsackStoreName string `name:"KnapsackStoreName" envconfig:"KNAPSACK_STORE_NAME" default:"knapsack"`
	KnapsackUrl       string `name:"KnapsackUrl" envconfig:"KNAPSACK_URL" default:"localhost:8081"`
}

func (g *KnapsackSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out KnapsackSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
