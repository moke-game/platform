package analyfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/gstones/platform/api/gen/analytics"
	"github.com/gstones/platform/services/analytics/pkg/global"
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
	if sSetting.MTLSEnable {
		if conn, err := tools.DialWithSecurity(
			host,
			sSetting.ClientCert,
			sSetting.ClientKey,
			sSetting.ServerName,
			sSetting.ServerCaCert,
		); err != nil {
			return nil, err
		} else {
			return pb.NewAnalyticsServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewAnalyticsServiceClient(conn), nil
		}
	}
}

var AnalyticsClientModule = fx.Invoke(
	func(
		setting AnalyticsSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out AnalyticsClientResult, err error) {
		if cli, e := NewAnalyticsClient(setting.AnalyticsUrl, sSetting); e != nil {
			err = e
		} else {
			out.AnalyticsClient = cli
			// set global analytics client
			global.SetAnalyticsClient(cli)
		}
		return
	},
)
