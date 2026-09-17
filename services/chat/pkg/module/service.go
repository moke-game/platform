package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/chat/internal/service/private"
	"github.com/moke-game/platform/services/chat/internal/service/public"
	"github.com/moke-game/platform/services/chat/pkg/cfx"
)

var settings = cfx.ChatSettingsModule

// ChatModule Provides chat service
var ChatModule = fx.Module("chat",
	settings,
	public.ChatService,
	private.ChatService,
)

// ChatClientModule Provides chat client for grpc
var ChatClientModule = fx.Module("chat_client",
	settings,
	cfx.ChatClientModule,
)

// ChatPrivateClientModule Provides chat private client for grpc
var ChatPrivateClientModule = fx.Module("chat_private_client",
	settings,
	cfx.ChatPrivateClientModule,
)

// ChatAllModule  Provides client, service for chat
var ChatAllModule = fx.Module("chat_all",
	settings,
	public.ChatService,
	private.ChatService,
	cfx.ChatClientModule,
	cfx.ChatPrivateClientModule,
)
