package bfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
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

func (g *BuddySettingsResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var BuddySettingsModule = fx.Provide(
	func() (out BuddySettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
