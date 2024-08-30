package rfx

import (
	"go.uber.org/fx"

	"github.com/gstones/moke-kit/utility"
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

func (g *RoomSettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out RoomSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
