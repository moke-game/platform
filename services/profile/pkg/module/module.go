package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform.git/services/profile/internal/private"
	"github.com/moke-game/platform.git/services/profile/internal/public"
	"github.com/moke-game/platform.git/services/profile/pkg/pfx"
)

// ProfileModule Provides profile service
var ProfileModule = fx.Module("profile",
	pfx.SettingsModule,
	public.Module,
	private.Module,
)

// ProfilePrivateModule Provides profile private service
var ProfilePrivateModule = fx.Module("profile_private",
	pfx.SettingsModule,
	private.Module,
)

// ProfileClientModule Provides profile client for grpc
var ProfileClientModule = fx.Module("profile_client",
	pfx.SettingsModule,
	pfx.ProfileClientModule,
)

// ProfileAllModule Provides client, service for profile
var ProfileAllModule = fx.Module("profile_all",
	pfx.SettingsModule,
	public.Module,
	private.Module,
	pfx.ProfileClientModule,
)
