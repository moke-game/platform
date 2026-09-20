package assembly

import (
	agones "github.com/gstones/moke-kit/3rd/agones/pkg/module"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"go.uber.org/fx"

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
	room "github.com/moke-game/platform/services/room/pkg/module"
)

// Infra bricks. fxmain.Main already has settings, logging, gRPC/gateway, Mongo, Redis client, MQ router.
var (
	NATS        = mfx.NatsModule
	Cache       = ofx.RedisCacheModule
	Agones      = agones.AgonesSDKModule
	Auth        = auth.AuthAllModule
	AuthMW      = auth.AuthMiddlewareModule
	AuthPrivate = auth.PrivateServiceAuthModule
)

// Service bricks (business only). Prefer Main("profile") so infra/auth are filled in.
var (
	Analytics   = analytics.AnalyticsModule
	Buddy       = buddy.BuddyModule
	Chat        = chat.ChatModule
	Knapsack    = knapsack.KnapsackModule
	Leaderboard = leaderboard.LeaderboardModule
	Mail        = mail.MailModule
	Matchmaking = matchmaking.MatchmakingModule
	Party       = party.PartyModule
	Profile     = profile.ProfileModule
	Room        = room.RoomModule
)

// Client bricks for a game process that dials platform.
var (
	AnalyticsClient   = analytics.AnalyticsClientModule
	AuthClient        = auth.AuthClientModule
	BuddyClient       = buddy.BuddyClientModule
	ChatClient        = chat.ChatClientModule
	KnapsackClient    = knapsack.KnapsackClientModule
	LeaderboardClient = leaderboard.LeaderboardClientModule
	MailClient        = mail.MailClientModule
	MatchmakingClient = matchmaking.MatchmakingClientModule
	PartyClient       = party.PartyClientModule
	ProfileClient     = profile.ProfileClientModule
)

var bricks = map[string]fx.Option{
	"nats":   NATS,
	"cache":  Cache,
	"agones": Agones,

	"auth":         Auth,
	"auth.mw":      AuthMW,
	"auth.private": AuthPrivate,
	"auth.client":  AuthClient,

	"analytics":                  Analytics,
	"analytics.client":           AnalyticsClient,
	"analytics.all":              analytics.AnalyticsAllModule,
	"buddy":                      Buddy,
	"buddy.client":               BuddyClient,
	"buddy.all":                  buddy.BuddyAllModule,
	"chat":                       Chat,
	"chat.client":                ChatClient,
	"chat.client.private":        chat.ChatPrivateClientModule,
	"chat.all":                   chat.ChatAllModule,
	"knapsack":                   Knapsack,
	"knapsack.client":            KnapsackClient,
	"knapsack.all":               knapsack.KnapsackAllModule,
	"leaderboard":                Leaderboard,
	"leaderboard.client":         LeaderboardClient,
	"leaderboard.client.private": leaderboard.LeaderboardClientPrivate,
	"leaderboard.all":            leaderboard.LeaderboardAllModule,
	"mail":                       Mail,
	"mail.client":                MailClient,
	"mail.client.private":        mail.MailClientPrivateModule,
	"mail.all":                   mail.MailAllModule,
	"matchmaking":                Matchmaking,
	"matchmaking.client":         MatchmakingClient,
	"party":                      Party,
	"party.client":               PartyClient,
	"party.all":                  party.PartyAllModule,
	"profile":                    Profile,
	"profile.client":             ProfileClient,
	"profile.all":                profile.ProfileAllModule,
	"room":                       Room,
}

// recipes are standalone processes. Values are brick names (not other recipes).
var recipes = map[string][]string{
	"auth":        {"cache", "auth"},
	"analytics":   {"analytics", "auth.private"},
	"profile":     {"nats", "cache", "auth.mw", "profile"},
	"knapsack":    {"nats", "cache", "auth.mw", "knapsack"},
	"buddy":       {"nats", "cache", "auth.mw", "buddy"},
	"chat":        {"nats", "auth.mw", "chat"},
	"mail":        {"nats", "auth.mw", "mail"},
	"party":       {"nats", "auth.mw", "party"},
	"leaderboard": {"auth.mw", "leaderboard"},
	"matchmaking": {"auth.mw", "matchmaking"},
	"room":        {"agones", "room"},
	"platform": {
		"nats", "cache", "auth",
		"analytics", "profile", "knapsack", "mail",
		"party", "buddy", "leaderboard", "chat", "matchmaking",
	},
}

// aliases map AI / Chinese words onto catalog names.
var aliases = map[string]string{
	"jwt":             "auth.mw",
	"auth.middleware": "auth.mw",
	"auth.all":        "auth",
	"redis":           "cache",
	"rediscache":      "cache",
	"all":             "platform",
	"认证":              "auth",
	"玩家":              "profile",
	"角色":              "profile",
	"背包":              "knapsack",
	"聊天":              "chat",
	"邮件":              "mail",
	"组队":              "party",
	"好友":              "buddy",
	"排行":              "leaderboard",
	"排行榜":             "leaderboard",
	"匹配":              "matchmaking",
	"房间":              "room",
	"统计":              "analytics",
	"分析":              "analytics",
	"中台":              "platform",
}
