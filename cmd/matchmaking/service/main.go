package main

import (
	"github.com/gstones/moke-kit/fxmain"

	"github.com/moke-game/platform/services/matchmaking/pkg/module"
)

func main() {
	fxmain.Main(module.App)
}
