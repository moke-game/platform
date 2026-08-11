# Bugbot review rules (platform)

## Auth assembly
- Flag missing `AuthCheckModule` / AuthMiddleware on public services.
- Flag pass-through / `PrivateServiceAuthModule` used on public (internet-facing) gRPC or gateway.
- Flag AuthMiddleware that does not set `UIDContextKey` on successful ValidateToken.
- Flag fail-open production guards (public ports without auth when prod fail-closed is required).

## Chat / MQ
- `Chatter.Destroy` is terminal: best-effort `Unsubscribe`, then always clear the map.
- Client `unSubscribe` may retain map entries when `Unsubscribe` fails (retryable).
- Flag mid-stream multi-topic maps that discard `Subscribe` return values.

## Contracts / deps
- Flag `go.mod` Go version drift vs Docker builder image (`golang:*`).
- Prefer contract tests for Auth middleware (missing/malformed bearer, validate fail, UID success).
