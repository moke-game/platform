package cfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/chat/api"
	"github.com/moke-game/platform/pkg/platformfx"
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
	return platformfx.NewClient(host, sSetting, pb.NewChatServiceClient)
}

var ChatClientModule = fx.Provide(
	func(
		setting ChatSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out ChatClientResult, err error) {
		out.ChatClient, err = NewChatClient(setting.ChatUrl, sSetting)
		return
	},
)
