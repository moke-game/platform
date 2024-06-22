package analyfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type AnalyticsSettingParams struct {
	fx.In

	IfunBiKinesisRegion     string `name:"IfunBiKinesisRegion"`
	IfunBiKinesisKey        string `name:"IfunBiKinesisKey"`
	IfunBiKinesisSecret     string `name:"IfunBiKinesisSecret"`
	IfunBiKinesisStreamName string `name:"IfunBiKinesisStreamName"`
	IfunBiAppId             string `name:"IfunBiAppId"`
	AnalyticsUrl            string `name:"AnalyticsUrl"`
	LocalBiPath             string `name:"LocalBiPath"`
	CKaddr                  string `name:"CKaddr"`
	CKdb                    string `name:"CKdb"`
	CKuser                  string `name:"CKuser"`
	CKpasswd                string `name:"CKpasswd"`
}

type AnalyticsSettingsResult struct {
	fx.Out

	IfunBiAppId             string `name:"IfunBiAppId" envconfig:"IFUN_BI_APP_ID" default:"5ab1c27fb7c04ce9abedeb23755a3aee"`
	IfunBiKinesisRegion     string `name:"IfunBiKinesisRegion" envconfig:"IFUN_BI_KINESIS_REGION" default:""`
	IfunBiKinesisKey        string `name:"IfunBiKinesisKey" envconfig:"IFUN_BI_KINESIS_KEY" default:""`
	IfunBiKinesisSecret     string `name:"IfunBiKinesisSecret" envconfig:"IFUN_BI_KINESIS_SECRET" default:""`
	IfunBiKinesisStreamName string `name:"IfunBiKinesisStreamName" envconfig:"IFUN_BI_KINESIS_STREAM_NAME" default:""`
	AnalyticsUrl            string `name:"AnalyticsUrl" envconfig:"ANALYTICS_URL" default:"localhost:8081"`
	LocalBiPath             string `name:"LocalBiPath" envconfig:"LOCAL_BI_PATH" default:"./logs/bi"`
	CKaddr                  string `name:"CKaddr" envconfig:"CK_ADDR" default:""`
	CKdb                    string `name:"CKdb" envconfig:"CK_DB" default:"fr"`
	CKuser                  string `name:"CKuser" envconfig:"CK_USER" default:""`
	CKpasswd                string `name:"CKpasswd" envconfig:"CK_PASSWD" default:""`
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
