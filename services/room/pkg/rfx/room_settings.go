package rfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
)

type RoomSettingParams struct {
	fx.In

	RoomUrl       string `name:"RoomUrl"`
	RoomCountMax  int32  `name:"RoomCountMax"`
	RoomPlayerMax int32  `name:"RoomPlayerMax"`
}

type RoomSettingsResult struct {
	fx.Out

	RoomUrl       string `name:"RoomUrl" envconfig:"ROOM_URL" default:"localhost:8888"`
	RoomCountMax  int32  `name:"RoomCountMax" envconfig:"ROOM_COUNT_MAX" default:"100"`
	RoomPlayerMax int32  `name:"RoomPlayerMax" envconfig:"ROOM_PLAYER_MAX" default:"100"`
}

var SettingsModule = platformfx.ProvideFromEnv[RoomSettingsResult]()
