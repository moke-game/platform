package cfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/moke-game/platform/api/gen/chat/api"
)

type ChatClientParams struct {
	fx.In

	ChatClient pb.ChatServiceClient `name:"ChatClient"`
}

type ChatClientResult struct {
	fx.Out

	ChatClient pb.ChatServiceClient `name:"ChatClient"`
}

func NewChatClient(host string, sSetting sfx.SecuritySettingsParams) (pb.ChatServiceClient, error) {
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
			return pb.NewChatServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewChatServiceClient(conn), nil
		}
	}
}

var ChatClientModule = fx.Provide(
	func(
		setting ChatSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out ChatClientResult, err error) {
		if cli, e := NewChatClient(setting.ChatUrl, sSetting); e != nil {
			err = e
		} else {
			out.ChatClient = cli
		}
		return
	},
)
