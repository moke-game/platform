package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/party/internal/service/public"
	"github.com/moke-game/platform/services/party/pkg/ptfx"
)

var settings = ptfx.PartySettingsModule

// PartyModule Provides party service
var PartyModule = fx.Module("party",
	settings,
	public.PartyService,
)

// PartyClientModule Provides party client for grpc
var PartyClientModule = fx.Module("party_client",
	settings,
	ptfx.PartyClientModule,
)

// PartyAllModule Provides client, service for party
var PartyAllModule = fx.Module("party_all",
	settings,
	public.PartyService,
	ptfx.PartyClientModule,
)
