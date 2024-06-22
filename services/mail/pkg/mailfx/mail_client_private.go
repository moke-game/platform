package mailfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/gstones/platform/api/gen/mail"
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
	if sSetting.MTLSEnable {
		if c, e := tools.DialWithSecurity(
			target,
			sSetting.ClientCert,
			sSetting.ClientKey,
			sSetting.ServerName,
			sSetting.ServerCaCert,
		); e != nil {
			return nil, e
		} else {
			return pb.NewMailPrivateServiceClient(c), nil
		}
	} else {
		if c, e := tools.DialInsecure(target); e != nil {
			return nil, e
		} else {
			return pb.NewMailPrivateServiceClient(c), nil
		}
	}
}

func (g *MailClientPrivateResult) Execute(
	a MailSettingParams,
	sSetting sfx.SecuritySettingsParams,
) (err error) {
	g.MailClient, err = NewMailPrivateClient(a.MailUrl, sSetting)
	return
}

var MailClientPrivateModule = fx.Provide(
	func(
		a MailSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out MailClientPrivateResult, err error) {
		err = out.Execute(a, sSetting)
		return
	},
)
