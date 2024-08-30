package main

import (
	agones "github.com/gstones/moke-kit/3rd/agones/pkg/module"
	awsConfig "github.com/gstones/moke-kit/3rd/cloud/pkg/module"
	"github.com/gstones/moke-kit/fxmain"

	mm "github.com/moke-game/platform/services/matchmaking/pkg/module"
)

func main() {
	fxmain.Main(
		agones.AgonesAllocateClientModule,
		awsConfig.AWSConfigModule,
		mm.MatchmakingModule,
	)
}
