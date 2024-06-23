package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform.git/services/party/internal/service/public"
	"github.com/moke-game/platform.git/services/party/pkg/ptfx"
)

// PartyModule Provides party service
var PartyModule = fx.Module("party",
	public.PartyService,
	ptfx.PartySettingsModule,
)

// PartyClientModule Provides party client for grpc
var PartyClientModule = fx.Module("party_client",
	ptfx.PartyClientModule,
	ptfx.PartySettingsModule,
)

// PartyAllModule Provides client, service for party
var PartyAllModule = fx.Module("party_all",
	public.PartyService,
	ptfx.PartyClientModule,
	ptfx.PartySettingsModule,
)
