package main

import (
	"github.com/gstones/moke-kit/fxmain"

	mm "github.com/moke-game/platform/services/matchmaking/pkg/module"
)

func main() {
	fxmain.Main(
		mm.MatchmakingModule,
	)
}
