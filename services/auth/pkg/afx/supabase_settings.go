package afx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
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

// LoadFromEnv load from env
func (g *SupabaseSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

// SupabaseSettingsModule is the supabase settings module
// you can find them in https://app.supabase.io/project/setting/api
var SupabaseSettingsModule = fx.Provide(
	func() (out SupabaseSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
