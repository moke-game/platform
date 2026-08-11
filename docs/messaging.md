# Messaging topic naming

Platform services publish/subscribe through moke-kit MQ helpers. Always wrap logical
names with `common.NatsHeader.CreateTopic(...)` (or `LocalHeader` for in-process).

## Conventions

| Service | Pattern | Helper |
|---------|---------|--------|
| Mail private | `mail.private.{uid}` | `services/mail/internal/service/common.MakePrivateTopic` |
| Mail public | `mail.public.{channel}` (`0` if empty) | `MakePublicTopic` |
| Profile changes | `profile.changes.{uid}` | `services/profile/changes.CreateTopic` |
| Knapsack changes | `knapsack.changes.{uid}` | `services/knapsack/changes.CreateTopic` |
| Chat | `{appId}.{deployment}.{channel}[.{id}]` | `Chatter.makeChatTopic` |

## Rules

1. **Never** publish raw logical strings to NATS — always `NatsHeader.CreateTopic`.
2. Prefer dedicated `topics.go` / `changes` helpers over inlined `fmt.Sprintf` in handlers.
3. Subscription lifecycle (kit #221+ cancelable `Subscription`):
   - **Context-scoped stream watchers** (mail/profile/… that only wait on `ctx.Done()`):
     parent cancel closes the sub; discarding the `Subscribe` return is OK when you never
     unsubscribe earlier.
   - **Mid-stream unsub / multi-topic maps** (chat): retain `miface.Subscription` and call
     `Unsubscribe()` when the client unsubscribes or the connection is destroyed.
4. Private topics are for internal/network-trusted consumers; public channel topics may fan out
   to many clients behind AuthMiddleware.

## Lifecycle

- Streaming watchers tied to request context: cancel → unsubscribe (via kit cancel helper).
- Chat `Chatter.Destroy`: best-effort `Unsubscribe` every retained subscription, then clear the map.
- Chat client `UnSubscribe`: only remove the map entry after `Unsubscribe` succeeds (retryable).
