package module

import (
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"go.uber.org/fx"

	auth "github.com/moke-game/platform/services/auth/pkg/module"
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

// App is the standalone chat process. Keep infra in sync with assembly recipe "chat".
var App = fx.Module("chat_app",
	mfx.NatsModule,
	auth.AuthMiddlewareModule,
	ChatModule,
)
