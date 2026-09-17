package kfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
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

var SettingsModule = platformfx.ProvideFromEnv[KnapsackSettingsResult]()
