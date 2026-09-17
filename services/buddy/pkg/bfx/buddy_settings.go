package bfx

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/pkg/platformfx"
)

type BuddySettingsParams struct {
	fx.In

	BuddyUrl        string `name:"BuddyUrl"`
	InviterMaxCount int32  `name:"InviterMaxCount"`
	BuddyMaxCount   int32  `name:"BuddyMaxCount"`
	BlockedMaxCount int32  `name:"BlockedMaxCount"`
	Name            string `name:"Name"`
}

type BuddySettingsResult struct {
	fx.Out

	BuddyUrl        string `name:"BuddyUrl" envconfig:"BUDDY_URL" default:"localhost:8081"`
	BuddyMaxCount   int32  `name:"BuddyMaxCount" envconfig:"BUDDY_MAX_COUNT" default:"1000"`
	BlockedMaxCount int32  `name:"BlockedMaxCount" envconfig:"BLOCKED_MAX_COUNT" default:"100"`
	InviterMaxCount int32  `name:"InviterMaxCount" envconfig:"INVITER_MAX_COUNT" default:"100"`
	Name            string `name:"Name" envconfig:"NAME" default:"buddy"`
}

var BuddySettingsModule = platformfx.ProvideFromEnv[BuddySettingsResult]()
