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
3. Subscriptions must retain `miface.Subscription` and call `Unsubscribe()` on cancel/destroy
   (kit #221+ returns a cancelable subscription; do not discard it).
4. Private topics are for internal/network-trusted consumers; public channel topics may fan out
   to many clients behind AuthMiddleware.

## Lifecycle

- Streaming watchers (mail/chat): cancel request context → unsubscribe.
- Chat `Chatter.Destroy` must `Unsubscribe` every retained subscription before clearing the map.
