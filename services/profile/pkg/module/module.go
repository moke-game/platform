package module

import (
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"go.uber.org/fx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
	"github.com/moke-game/platform/services/profile/internal/private"
	"github.com/moke-game/platform/services/profile/internal/public"
	"github.com/moke-game/platform/services/profile/pkg/pfx"
)

var settings = pfx.SettingsModule

// ProfileModule Provides profile service
var ProfileModule = fx.Module("profile",
	settings,
	public.Module,
	private.Module,
)

// ProfilePrivateModule Provides profile private service
var ProfilePrivateModule = fx.Module("profile_private",
	settings,
	private.Module,
)

// ProfileClientModule Provides profile client for grpc
var ProfileClientModule = fx.Module("profile_client",
	settings,
	pfx.ProfileClientModule,
)

// ProfileAllModule Provides client, service for profile
var ProfileAllModule = fx.Module("profile_all",
	settings,
	public.Module,
	private.Module,
	pfx.ProfileClientModule,
)

// App is the standalone profile process. Keep infra in sync with assembly recipe "profile".
var App = fx.Module("profile_app",
	mfx.NatsModule,
	ofx.RedisCacheModule,
	auth.AuthMiddlewareModule,
	ProfileModule,
)
