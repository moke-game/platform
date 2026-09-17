package afx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
)

// SupabaseSettingParams module params for injecting SupabaseSettings
type SupabaseSettingParams struct {
	fx.In

	URL string `name:"supabaseUrl"`
	Key string `name:"supabaseKey"`
}

// SupabaseSettingsResult module result for exporting SupabaseSettings
type SupabaseSettingsResult struct {
	fx.Out

	URL string `name:"supabaseUrl" envconfig:"SUPABASE_URL" default:""`
	Key string `name:"supabaseKey" envconfig:"SUPABASE_KEY" default:""`
}

// SupabaseSettingsModule is the supabase settings module
// you can find them in https://app.supabase.io/project/setting/api
var SupabaseSettingsModule = platformfx.ProvideFromEnv[SupabaseSettingsResult]()
