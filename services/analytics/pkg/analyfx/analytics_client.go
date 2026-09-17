package analyfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/analytics/api"
	"github.com/moke-game/platform/pkg/platformfx"
	"github.com/moke-game/platform/services/analytics/pkg/global"
)

type AnalyticsClientParams struct {
	fx.In

	AnalyticsClient pb.AnalyticsServiceClient `name:"AnalyticsClient"`
}

type AnalyticsClientResult struct {
	fx.Out

	AnalyticsClient pb.AnalyticsServiceClient `name:"AnalyticsClient"`
}

func NewAnalyticsClient(host string, sSetting sfx.SecuritySettingsParams) (pb.AnalyticsServiceClient, error) {
	return platformfx.NewClient(host, sSetting, pb.NewAnalyticsServiceClient)
}

var AnalyticsClientModule = fx.Provide(
	func(
		setting AnalyticsSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out AnalyticsClientResult, err error) {
		out.AnalyticsClient, err = NewAnalyticsClient(setting.AnalyticsUrl, sSetting)
		if err == nil {
			global.SetAnalyticsClient(out.AnalyticsClient)
		}
		return
	},
)
