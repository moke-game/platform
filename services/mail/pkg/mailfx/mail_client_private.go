package mailfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/mail/api"
	"github.com/moke-game/platform/pkg/platformfx"
)

type MailClientPrivateParams struct {
	fx.In
	MailClient pb.MailPrivateServiceClient `name:"MailPrivateClient"`
}

type MailClientPrivateResult struct {
	fx.Out
	MailClient pb.MailPrivateServiceClient `name:"MailPrivateClient"`
}

func NewMailPrivateClient(target string, sSetting sfx.SecuritySettingsParams) (pb.MailPrivateServiceClient, error) {
	return platformfx.NewClient(target, sSetting, pb.NewMailPrivateServiceClient)
}

var MailClientPrivateModule = fx.Provide(
	func(
		a MailSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out MailClientPrivateResult, err error) {
		out.MailClient, err = NewMailPrivateClient(a.MailUrl, sSetting)
		return
	},
)
