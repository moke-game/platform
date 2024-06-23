package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform.git/services/knapsack/internal/service/private"
	"github.com/moke-game/platform.git/services/knapsack/internal/service/public"
	"github.com/moke-game/platform.git/services/knapsack/pkg/kfx"
)

// KnapsackModule Provides knapsack service
var KnapsackModule = fx.Module("knapsack",
	kfx.SettingsModule,
	public.Module,
	private.Module,
)

// KnapsackPrivateModule Provides knapsack private service
var KnapsackPrivateModule = fx.Module("knapsack_private",
	kfx.SettingsModule,
	private.Module,
)

// KnapsackClientModule Provides knapsack client for grpc
var KnapsackClientModule = fx.Module("knapsack_client",
	kfx.SettingsModule,
	kfx.KnapsackClientModule,
)

// KnapsackAllModule Provides client, service for knapsack
var KnapsackAllModule = fx.Module("knapsack_all",
	kfx.SettingsModule,
	public.Module,
	private.Module,
	kfx.KnapsackClientModule,
)
