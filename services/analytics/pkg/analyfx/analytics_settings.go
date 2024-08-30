package analyfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type AnalyticsSettingParams struct {
	fx.In
	AnalyticsUrl string `name:"AnalyticsUrl"`

	// LocalBiPath is the path of the local bi logs, default is ./logs/bi.
	LocalBiPath string `name:"LocalBiPath"`

	// ClickHouse settings.
	CKAddr   string `name:"CKAddr"`
	CKDB     string `name:"CKDB"`
	CKUser   string `name:"CKUser"`
	CKPasswd string `name:"CKPasswd"`

	// ThinkingData settings.
	TDPath string `name:"TDPath"`
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
	CKAddr   string `name:"CKAddr" envconfig:"CK_ADDR" default:""`
	CKDB     string `name:"CKDB" envconfig:"CK_DB" default:"default"`
	CKUser   string `name:"CKUser" envconfig:"CK_USER" default:""`
	CKPasswd string `name:"CKPasswd" envconfig:"CK_PASSWD" default:""`

	// ThinkingData settings.
	TDPath string `name:"TDPath" envconfig:"TD_PATH" default:"./logs/td"`
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
