package module

import (
	"go.uber.org/fx"

	"github.com/gstones/platform/services/chat/internal/service/private"
	"github.com/gstones/platform/services/chat/internal/service/public"
	"github.com/gstones/platform/services/chat/pkg/cfx"
)

// ChatModule Provides chat service
var ChatModule = fx.Module("chat",
	public.ChatService,
	private.ChatService,
	cfx.ChatSettingsModule,
)

// ChatClientModule Provides chat client for grpc
var ChatClientModule = fx.Module("chat_client",
	cfx.ChatClientModule,
	cfx.ChatSettingsModule,
)

// ChatPrivateClientModule Provides chat private client for grpc
var ChatPrivateClientModule = fx.Module("chat_private_client",
	cfx.ChatPrivateClientModule,
	cfx.ChatSettingsModule,
)

// ChatAllModule  Provides client, service for chat
var ChatAllModule = fx.Module("chat_all",
	public.ChatService,
	private.ChatService,
	cfx.ChatClientModule,
	cfx.ChatSettingsModule,
	cfx.ChatPrivateClientModule,
)
