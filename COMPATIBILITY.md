# Compatibility

| Consumer | Minimum moke-kit | Notes |
|----------|------------------|-------|
| platform (this repo) | `v1.0.5-0.20260811094419-bcdfe55515cd` (#224) | #221 Subscribe/CAS/CORS/TLS/Auth + #224 binder startup fail-closed (prod auth + CORS) |
| game | platform commit that depends on the kit version above | See [game#18](https://github.com/moke-game/game/issues/18) |

Auth startup fail-closed:

- Platform `ProdAuthGuardModule` (inside `AuthMiddlewareModule` / `AuthAllModule` / `PrivateServiceAuthModule`) covers public gRPC + gateway when those modules are imported.
- Kit binder (`ValidateSecurityConfig`, #224) fails closed when grpc/gateway lack middleware or prod CORS allowlist is empty.
- `cmd/auth` uses `AuthAllModule`; private-only `cmd/analytics` uses `PrivateServiceAuthModule`.

Messaging topic conventions: [docs/messaging.md](docs/messaging.md).

Tracking:

- kit plan: https://github.com/GStones/moke-kit/issues/223
- kit hardening: https://github.com/GStones/moke-kit/pull/224
- platform: https://github.com/moke-game/platform/issues/23
- game: https://github.com/moke-game/game/issues/18
