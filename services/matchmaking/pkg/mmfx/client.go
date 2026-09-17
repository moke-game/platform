package mmfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	matchmaking "github.com/moke-game/platform/api/gen/matchmaking/api"
	"github.com/moke-game/platform/pkg/platformfx"
)

type ClientParams struct {
	fx.In

	Client matchmaking.MatchServiceClient `name:"MatchServiceClient"`
}

type ClientResult struct {
	fx.Out

	Client matchmaking.MatchServiceClient `name:"MatchServiceClient"`
}

func NewClient(host string, setting sfx.SecuritySettingsParams) (matchmaking.MatchServiceClient, error) {
	return platformfx.NewClient(host, setting, matchmaking.NewMatchServiceClient)
}

var ClientModule = fx.Provide(
	func(
		setting MatchmakingSettingParams,
		security sfx.SecuritySettingsParams,
	) (out ClientResult, err error) {
		out.Client, err = NewClient(setting.URL, security)
		return
	},
)
