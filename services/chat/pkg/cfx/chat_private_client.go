package cfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	pb "github.com/moke-game/platform/api/gen/chat/api"
	"github.com/moke-game/platform/pkg/platformfx"
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
	return platformfx.NewClient(host, sSetting, pb.NewChatPrivateServiceClient)
}

var ChatPrivateClientModule = fx.Provide(
	func(
		setting ChatSettingParams,
		sSetting sfx.SecuritySettingsParams,
	) (out ChatPrivateClientResult, err error) {
		out.ChatClient, err = NewChatPrivateClient(setting.ChatUrl, sSetting)
		return
	},
)
