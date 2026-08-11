# Compatibility

| Consumer | Minimum moke-kit | Notes |
|----------|------------------|-------|
| platform (this repo) | `v1.0.5-0.20260811094419-bcdfe55515cd` (#224) | #221 Subscribe/CAS/CORS/TLS/Auth + #224 binder startup fail-closed (prod auth + CORS) |
| game | platform `main` at/after [#25](https://github.com/moke-game/platform/pull/25) (`fcdebd60…`) | [game#20](https://github.com/moke-game/game/pull/20) |

Auth startup fail-closed:

- Platform `ProdAuthGuardModule` (inside `AuthMiddlewareModule` / `AuthAllModule` / `PrivateServiceAuthModule`) covers public gRPC + gateway when those modules are imported.
- Kit binder (`ValidateSecurityConfig`, #224) fails closed when grpc/gateway lack middleware or prod CORS allowlist is empty.
- `cmd/auth` uses `AuthAllModule`; private-only `cmd/analytics` uses `PrivateServiceAuthModule`.

Messaging topic conventions: [docs/messaging.md](docs/messaging.md).

Merged:

- kit hardening: https://github.com/GStones/moke-kit/pull/224
- kit create-game sync: https://github.com/GStones/moke-kit/pull/226
- platform P0: https://github.com/moke-game/platform/pull/24
- platform P1: https://github.com/moke-game/platform/pull/25
- game P0/P1: https://github.com/moke-game/game/pull/19
- game P2 bump: https://github.com/moke-game/game/pull/20

Tracking plans:

- kit: https://github.com/GStones/moke-kit/issues/223
- platform: https://github.com/moke-game/platform/issues/23
- game: https://github.com/moke-game/game/issues/18
