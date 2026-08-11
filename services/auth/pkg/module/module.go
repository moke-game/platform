package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/auth/internal"
	"github.com/moke-game/platform/services/auth/pkg/afx"
)

// AuthModule provides the Auth gRPC service (Authenticate / ValidateToken / …).
// Use this for the dedicated auth microservice (cmd/auth). The auth service itself
// embeds utility.WithoutAuth so login flows remain reachable without a prior token.
var AuthModule = fx.Module("auth",
	afx.SettingsModule,
	internal.ServiceModule,
)

// AuthClientModule provides an AuthServiceClient for calling AuthService.
var AuthClientModule = fx.Module("auth_client",
	afx.SettingsModule,
	afx.AuthClientModule,
)

// AuthMiddlewareModule provides AuthServiceClient + AuthCheckModule middleware.
// Import this in every process that hosts public gRPC/HTTP services.
// Private *PrivateService handlers should embed utility.WithoutAuth and stay
// on internal networks only.
var AuthMiddlewareModule = fx.Module("auth_middleware",
	afx.SettingsModule,
	afx.AuthClientModule,
	afx.AuthCheckModule,
	afx.ProdAuthGuardModule,
)

// SupabaseMiddlewareModule provides supabase middleware for grpc.
// if import this module, every grpc unary/stream will auth by supabase auth
// https://supabase.com/docs/guides/auth
var SupabaseMiddlewareModule = fx.Module("supabase_middleware",
	afx.SupabaseSettingsModule,
	afx.SupabaseCheckModule,
)

// AuthAllModule provides auth service + client + middleware for aggregate/monolith
// processes (cmd/platform, local game aggregate). Equivalent to AuthModule plus
// AuthMiddlewareModule without duplicating settings providers.
var AuthAllModule = fx.Module("auth_all",
	afx.SettingsModule,
	internal.ServiceModule,
	afx.AuthClientModule,
	afx.AuthCheckModule,
	afx.ProdAuthGuardModule,
)
