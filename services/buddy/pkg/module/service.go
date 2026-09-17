package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/buddy/internal/service"
	"github.com/moke-game/platform/services/buddy/pkg/bfx"
)

var settings = bfx.BuddySettingsModule

// BuddyModule Provides buddy service
var BuddyModule = fx.Module("buddy",
	settings,
	service.Module,
)

// BuddyClientModule Provides buddy client for grpc
var BuddyClientModule = fx.Module("buddy_client",
	settings,
	bfx.BuddyClientModule,
)

// BuddyAllModule  Provides client, service for buddy
var BuddyAllModule = fx.Module("buddy_all",
	settings,
	service.Module,
	bfx.BuddyClientModule,
)
