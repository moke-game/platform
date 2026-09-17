package module

import (
	"go.uber.org/fx"

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
