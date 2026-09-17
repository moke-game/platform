package mmfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
)

type MatchmakingSettingParams struct {
	fx.In

	// Matchmaking service URL
	URL string `name:"matchmakingUrl"`

	// AWS subnets
	AWSVPCSubnets string `name:"awsVPCSubnets"`

	// Open Match Frontend URL
	OMFrontendUrl string `name:"frontendUrl"`
	// Open Match Backend URL
	OMBackendUrl string `name:"backendUrl"`
	// Open Match Function URL
	OMFuncUrl string `name:"funcUrl"`
	// Open Match Function Port
	OMFuncPort int32 `name:"funcPort"`
}

type MatchmakingSettingResult struct {
	fx.Out

	// AWS subnets
	AWSVPCSubnets string `name:"awsVPCSubnets" envconfig:"AWS_VPC_SUBNETS" default:"subnet-0b1b2c3d4e5f6g7h8"`

	// Matchmaking service URL
	URL string `name:"matchmakingUrl" envconfig:"MATCHMAKING_URL" default:"localhost:8081"`
	// Open Match Frontend URL
	OMFrontendUrl string `name:"frontendUrl" envconfig:"OM_FRONTEND_URL" default:"localhost:50504"`
	// Open Match Backend URL
	OMBackendUrl string `name:"backendUrl" envconfig:"OM_BACKEND_URL" default:"localhost:50505"`

	// Open Match Function URL
	OMFuncUrl string `name:"funcUrl" envconfig:"OM_FUNC_URL" default:"192.168.50.11"`
	// Open Match Function Port
	OMFuncPort int32 `name:"funcPort" envconfig:"OM_FUNC_PORT" default:"8081"`
}

var MatchmakingSettingsModule = platformfx.ProvideFromEnv[MatchmakingSettingResult]()

// MatchmakingSettingModule is a compatibility alias.
var MatchmakingSettingModule = MatchmakingSettingsModule
