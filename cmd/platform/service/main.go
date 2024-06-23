package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	analytics "github.com/moke-game/platform/services/analytics/pkg/module"
	auth "github.com/moke-game/platform/services/auth/pkg/module"
	buddy "github.com/moke-game/platform/services/buddy/pkg/module"
	chat "github.com/moke-game/platform/services/chat/pkg/module"
	knapsack "github.com/moke-game/platform/services/knapsack/pkg/module"
	"github.com/moke-game/platform/services/leaderboard/pkg/module"
	mail "github.com/moke-game/platform/services/mail/pkg/module"
	match "github.com/moke-game/platform/services/matchmaking/pkg/module"
	party "github.com/moke-game/platform/services/party/pkg/module"
	profile "github.com/moke-game/platform/services/profile/pkg/module"
)

func initEnvs() {
	err := os.Setenv("APP_NAME", "platform")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("TIMEOUT", "15")
	if err != nil {
		panic(err)
	}
}
func initPprof() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
}
func main() {
	initEnvs()
	initPprof()
	fxmain.Main(
		// platform services
		auth.AuthAllModule,
		profile.ProfileAllModule,
		knapsack.KnapsackAllModule,
		chat.ChatAllModule,
		party.PartyModule,
		match.MatchModule,
		buddy.BuddyModule,
		mail.MailAllModule,
		analytics.AnalyticsModule,
		module.LeaderboardService,

		// infrastructures
		mfx.NatsModule,
		ofx.RedisCacheModule,
	)
}
