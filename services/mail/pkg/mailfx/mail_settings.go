package mailfx

import (
	"time"

	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type MailSettingParams struct {
	fx.In
	MailStoreName     string `name:"MailStoreName"`
	MailUrl           string `name:"MailUrl"`
	MailNumMax        int    `name:"MailNumMax"`
	MailDefaultExpire int32  `name:"MailDefaultExpire"`
	MailEncryptionKey string `name:"MailEncryptionKey"`
}

type MailSettingsResult struct {
	fx.Out
	MailStoreName string `name:"MailStoreName" envconfig:"MAIL_STORE_NAME" default:"mail"`
	MailUrl       string `name:"MailUrl" envconfig:"MAIL_URL" default:"localhost:8081"`
	// MailNumMax is the max number of mail that can be stored in the mail store
	MailNumMax int `name:"MailNumMax" envconfig:"MAIL_NUM_MAX" default:"99"`
	// MailDefaultExpire is the default expire time of mail (day)
	MailDefaultExpire int32 `name:"MailDefaultExpire" envconfig:"MAIL_DEFAULT_EXPIRE" default:"90"`
	// MailEncryptionKey is the key used to encrypt mail data
	MailEncryptionKey string `name:"MailEncryptionKey" envconfig:"MAIL_ENCRYPTION_KEY" default:"CTeGahnbQWfAr5hW"`
}

func (msl *MailSettingsResult) LoadFromEnv() (err error) {
	time.Now().Unix()
	err = utility.Load(msl)
	return
}

var MailSettingsModule = fx.Provide(
	func() (out MailSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
