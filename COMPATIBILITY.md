# Compatibility

| Consumer | Minimum moke-kit | Notes |
|----------|------------------|-------|
| platform (this repo) | `v1.0.5-0.20260811085843-60e104db6db7` (#221) | Subscribe cancel, DocumentBase CAS, CORS allowlist, TLS/mTLS split, request-path auth fail-closed |
| game | platform commit that depends on the kit version above | See [game#18](https://github.com/moke-game/game/issues/18) |

Auth startup fail-closed: platform `ProdAuthGuardModule` (inside `AuthMiddlewareModule` / `AuthAllModule`) covers public gRPC + gateway when those modules are imported. Binder-level enforcement for “forgot middleware entirely” is tracked in kit [#224](https://github.com/GStones/moke-kit/pull/224).

Tracking:

- kit plan: https://github.com/GStones/moke-kit/issues/223
- kit hardening PR (startup binder): https://github.com/GStones/moke-kit/pull/224
- platform: https://github.com/moke-game/platform/issues/23
- game: https://github.com/moke-game/game/issues/18
