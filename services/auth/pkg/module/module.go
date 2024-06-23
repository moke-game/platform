package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/auth/pkg/afx"
	public "github.com/moke-game/platform/services/auth/service"
)

// AuthModule Provides auth service
var AuthModule = fx.Module("auth",
	afx.SettingsModule,
	public.ServiceModule,
)

// AuthClientModule Provides auth client for grpc
var AuthClientModule = fx.Module("auth_client",
	afx.SettingsModule,
	afx.AuthClientModule,
)

// AuthMiddlewareModule Provides auth middleware for grpc
// if import this module, every grpc unary/stream will auth by {mfx.AuthCheckModule}
var AuthMiddlewareModule = fx.Module("auth_middleware",
	afx.SettingsModule,
	afx.AuthClientModule,
	afx.AuthCheckModule,
)

// AuthAllModule  Provides client, service and middleware for auth
var AuthAllModule = fx.Module("auth_all",
	public.ServiceModule,
	afx.AuthClientModule,
	afx.AuthCheckModule,
	afx.SettingsModule,
)
