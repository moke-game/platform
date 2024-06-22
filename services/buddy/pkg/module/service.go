package module

import (
	"go.uber.org/fx"

	"github.com/gstones/platform/services/buddy/internal/service"
	"github.com/gstones/platform/services/buddy/pkg/bfx"
)

// BuddyModule Provides buddy service
var BuddyModule = fx.Module("buddy",
	service.Module,
	bfx.BuddySettingsModule,
)

// BuddyClientModule Provides buddy client for grpc
var BuddyClientModule = fx.Module("buddy_client",
	bfx.BuddyClientModule,
	bfx.BuddySettingsModule,
)

// BuddyAllModule  Provides client, service for buddy
var BuddyAllModule = fx.Module("buddy_all",
	service.Module,
	bfx.BuddyClientModule,
	bfx.BuddySettingsModule,
)
