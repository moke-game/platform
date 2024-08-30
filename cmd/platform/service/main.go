package main

import (
	agones "github.com/gstones/moke-kit/3rd/agones/pkg/module"
	awsConfig "github.com/gstones/moke-kit/3rd/cloud/pkg/module"
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	analytics "github.com/moke-game/platform/services/analytics/pkg/module"
	auth "github.com/moke-game/platform/services/auth/pkg/module"
	buddy "github.com/moke-game/platform/services/buddy/pkg/module"
	chat "github.com/moke-game/platform/services/chat/pkg/module"
	knapsack "github.com/moke-game/platform/services/knapsack/pkg/module"
	leaderboard "github.com/moke-game/platform/services/leaderboard/pkg/module"
	mail "github.com/moke-game/platform/services/mail/pkg/module"
	matchmaking "github.com/moke-game/platform/services/matchmaking/pkg/module"
	party "github.com/moke-game/platform/services/party/pkg/module"
	profile "github.com/moke-game/platform/services/profile/pkg/module"
)

func main() {
	fxmain.Main(
		// infrastructure
		mfx.NatsModule,
		ofx.RedisCacheModule,
		agones.AgonesAllocateClientModule,
		awsConfig.AWSConfigModule,

		mail.MailModule,
		analytics.AnalyticsModule,
		auth.AuthAllModule,
		profile.ProfileModule,
		knapsack.KnapsackModule,
		party.PartyModule,
		buddy.BuddyModule,
		leaderboard.LeaderboardModule,
		chat.ChatModule,
		matchmaking.MatchmakingModule,
	)
}
