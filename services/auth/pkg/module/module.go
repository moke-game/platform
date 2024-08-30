package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/auth/internal"
	"github.com/moke-game/platform/services/auth/pkg/afx"
)

// AuthModule Provides auth service
var AuthModule = fx.Module("auth",
	afx.SettingsModule,
	internal.ServiceModule,
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

// SupabaseMiddlewareModule Provides supabase middleware for grpc
// if import this module, every grpc unary/stream will auth by supabase auth
// https://supabase.com/docs/guides/auth
var SupabaseMiddlewareModule = fx.Module("supabase_middleware",
	afx.SupabaseSettingsModule,
	afx.SupabaseCheckModule,
)

// AuthAllModule  Provides client, service and middleware for auth
var AuthAllModule = fx.Module("auth_all",
	internal.ServiceModule,
	afx.AuthClientModule,
	//afx.AuthCheckModule,
	afx.SettingsModule,
)
