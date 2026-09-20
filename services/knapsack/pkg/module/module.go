package module

import (
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"go.uber.org/fx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
	"github.com/moke-game/platform/services/knapsack/internal/service/private"
	"github.com/moke-game/platform/services/knapsack/internal/service/public"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
)

var settings = kfx.SettingsModule

// KnapsackModule Provides knapsack service
var KnapsackModule = fx.Module("knapsack",
	settings,
	public.Module,
	private.Module,
)

// KnapsackPrivateModule Provides knapsack private service
var KnapsackPrivateModule = fx.Module("knapsack_private",
	settings,
	private.Module,
)

// KnapsackClientModule Provides knapsack client for grpc
var KnapsackClientModule = fx.Module("knapsack_client",
	settings,
	kfx.KnapsackClientModule,
)

// KnapsackAllModule Provides client, service for knapsack
var KnapsackAllModule = fx.Module("knapsack_all",
	settings,
	public.Module,
	private.Module,
	kfx.KnapsackClientModule,
)

// App is the standalone knapsack process. Keep infra in sync with assembly recipe "knapsack".
var App = fx.Module("knapsack_app",
	mfx.NatsModule,
	ofx.RedisCacheModule,
	auth.AuthMiddlewareModule,
	KnapsackModule,
)
