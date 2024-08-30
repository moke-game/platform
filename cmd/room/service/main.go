package main

import (
	agones "github.com/gstones/moke-kit/3rd/agones/pkg/module"
	"github.com/gstones/moke-kit/fxmain"

	"github.com/moke-game/platform/services/room/pkg/module"
)

func main() {
	fxmain.Main(
		agones.AgonesSDKModule,
		module.RoomModule,
	)
}
