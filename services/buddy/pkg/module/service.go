package module

import (
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"go.uber.org/fx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
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

// App is the standalone buddy process. Keep infra in sync with assembly recipe "buddy".
var App = fx.Module("buddy_app",
	mfx.NatsModule,
	ofx.RedisCacheModule,
	auth.AuthMiddlewareModule,
	BuddyModule,
)
