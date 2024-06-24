package analyfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type AnalyticsSettingParams struct {
	fx.In
	AnalyticsUrl string `name:"AnalyticsUrl"`

	LocalBiPath string `name:"LocalBiPath"`

	CKaddr   string `name:"CKaddr"`
	CKdb     string `name:"CKdb"`
	CKuser   string `name:"CKuser"`
	CKpasswd string `name:"CKpasswd"`
}

type AnalyticsSettingsResult struct {
	fx.Out
	// AnalyticsUrl is the url of the analytics service, default is localhost:8081.
	// 当前统计服务器地址, 默认为localhost:8081
	AnalyticsUrl string `name:"AnalyticsUrl" envconfig:"ANALYTICS_URL" default:"localhost:8081"`
	// LocalBiPath is the path of the local bi logs, default is ./logs/bi.
	// 本地bi日志路径, 默认为./logs/bi
	LocalBiPath string `name:"LocalBiPath" envconfig:"LOCAL_BI_PATH" default:"./logs/bi"`
	// ClickHouse settings.
	CKaddr   string `name:"CKaddr" envconfig:"CK_ADDR" default:""`
	CKdb     string `name:"CKdb" envconfig:"CK_DB" default:"default"`
	CKuser   string `name:"CKuser" envconfig:"CK_USER" default:""`
	CKpasswd string `name:"CKpasswd" envconfig:"CK_PASSWD" default:""`
}

func (g *AnalyticsSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out AnalyticsSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
