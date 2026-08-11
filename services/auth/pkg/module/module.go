package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/auth/internal"
	"github.com/moke-game/platform/services/auth/pkg/afx"
)

// AuthModule provides the Auth gRPC service (Authenticate / ValidateToken / …).
// Prefer AuthAllModule for cmd/auth so kit binder prod checks (#224+) see middleware.
// The auth service embeds utility.WithoutAuth so login flows remain reachable.
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
// processes (cmd/platform, cmd/auth, local game). Equivalent to AuthModule plus
// AuthMiddlewareModule without duplicating settings providers.
// AuthService embeds utility.WithoutAuth so Authenticate/ValidateToken stay public.
var AuthAllModule = fx.Module("auth_all",
	afx.SettingsModule,
	internal.ServiceModule,
	afx.AuthClientModule,
	afx.AuthCheckModule,
	afx.ProdAuthGuardModule,
)

// PrivateServiceAuthModule provides a pass-through AuthMiddleware for processes
// that only host utility.WithoutAuth services (e.g. analytics). Satisfies kit
// binder prod checks (#224+) without requiring a real AuthService client.
var PrivateServiceAuthModule = fx.Module("private_service_auth",
	afx.PassThroughAuthModule,
	afx.ProdAuthGuardModule,
)
