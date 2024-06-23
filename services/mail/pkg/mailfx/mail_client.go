package mailfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/moke-game/platform.git/api/gen/mail"
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
			return pb.NewMailServiceClient(c), nil
		}
	} else {
		if c, e := tools.DialInsecure(target); e != nil {
			return nil, e
		} else {
			return pb.NewMailServiceClient(c), nil
		}
	}
}

func (g *MailClientResult) Execute(
	a MailSettingParams,
	sSetting sfx.SecuritySettingsParams,
) (err error) {
	g.MailClient, err = NewMailClient(a.MailUrl, sSetting)
	return
}

var MailClientModule = fx.Provide(
	func(
		a MailSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out MailClientResult, err error) {
		err = out.Execute(a, sSetting)
		return
	},
)
