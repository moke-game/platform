package module

import (
	"go.uber.org/fx"

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
