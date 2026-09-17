package mailfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/mail/api"
	"github.com/moke-game/platform/pkg/platformfx"
)

type MailClientParams struct {
	fx.In
	MailClient pb.MailServiceClient `name:"MailClient"`
}

type MailClientResult struct {
	fx.Out
	MailClient pb.MailServiceClient `name:"MailClient"`
}

func NewMailClient(target string, sSetting sfx.SecuritySettingsParams) (pb.MailServiceClient, error) {
	return platformfx.NewClient(target, sSetting, pb.NewMailServiceClient)
}

var MailClientModule = fx.Provide(
	func(
		a MailSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out MailClientResult, err error) {
		out.MailClient, err = NewMailClient(a.MailUrl, sSetting)
		return
	},
)
