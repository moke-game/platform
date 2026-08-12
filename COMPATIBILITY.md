# Compatibility

| Consumer | Minimum moke-kit | Notes |
|----------|------------------|-------|
| platform (this repo) | `v1.0.5-0.20260812022140-acb9f313d7fd` ([#228](https://github.com/GStones/moke-kit/pull/228)) | #221 Subscribe/CAS/CORS/TLS/Auth + #224 binder fail-closed + #228 CAS/StopServing/CI |
| game | platform `main` at/after this pin bump | kit `#228` + jwt expire prod guard |

Auth startup fail-closed:

- Platform `ProdAuthGuardModule` (inside `AuthMiddlewareModule` / `AuthAllModule` / `PrivateServiceAuthModule`) covers public gRPC + gateway when those modules are imported.
- Kit binder (`ValidateSecurityConfig`, #224) fails closed when grpc/gateway lack middleware or prod CORS allowlist is empty.
- `cmd/auth` uses `AuthAllModule`; private-only `cmd/analytics` uses `PrivateServiceAuthModule`.
- `JwtTokenExpire<=0` omits JWT `exp` and Redis TTL only outside prod; prod requires `JWT_TOKEN_EXPIRE > 0`.

Messaging topic conventions: [docs/messaging.md](docs/messaging.md).

Merged:

- kit hardening: https://github.com/GStones/moke-kit/pull/224
- kit create-game sync: https://github.com/GStones/moke-kit/pull/226
- kit CAS/StopServing/CI: https://github.com/GStones/moke-kit/pull/228
- platform P0: https://github.com/moke-game/platform/pull/24
- platform P1: https://github.com/moke-game/platform/pull/25
- platform jwt/v5 + CAS/chat tests: https://github.com/moke-game/platform/pull/27
- game P0/P1: https://github.com/moke-game/game/pull/19
- game P2 bump: https://github.com/moke-game/game/pull/20
- game CI + platform pin: https://github.com/moke-game/game/pull/22
- game post-merge pin: https://github.com/moke-game/game/pull/23

Tracking plans:

- kit: https://github.com/GStones/moke-kit/issues/223
- platform: https://github.com/moke-game/platform/issues/23
- game: https://github.com/moke-game/game/issues/18
