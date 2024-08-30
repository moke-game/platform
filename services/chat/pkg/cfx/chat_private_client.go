package cfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/moke-game/platform/api/gen/chat/api"
)

type ChatPrivateClientParams struct {
	fx.In

	ChatClient pb.ChatPrivateServiceClient `name:"ChatPrivateClient"`
}

type ChatPrivateClientResult struct {
	fx.Out

	ChatClient pb.ChatPrivateServiceClient `name:"ChatPrivateClient"`
}

func NewChatPrivateClient(host string, sSetting sfx.SecuritySettingsParams) (pb.ChatPrivateServiceClient, error) {
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
			return pb.NewChatPrivateServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewChatPrivateServiceClient(conn), nil
		}
	}
}

var ChatPrivateClientModule = fx.Provide(
	func(
		setting ChatSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out ChatPrivateClientResult, err error) {
		if cli, e := NewChatPrivateClient(setting.ChatUrl, sSetting); e != nil {
			err = e
		} else {
			out.ChatClient = cli
		}
		return
	},
)
