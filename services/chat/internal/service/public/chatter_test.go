package public

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	pb "github.com/moke-game/platform/api/gen/chat/api"
)

type fakeSub struct {
	unsubErr error
	calls    *int
}

func (f *fakeSub) IsValid() bool { return true }
func (f *fakeSub) Unsubscribe() error {
	if f.calls != nil {
		*f.calls++
	}
	return f.unsubErr
}

type fakeMQ struct {
	sub miface.Subscription
	err error
}

func (f *fakeMQ) Subscribe(context.Context, string, miface.SubResponseHandler, ...miface.SubOption) (miface.Subscription, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.sub, nil
}
func (f *fakeMQ) Publish(string, ...miface.PubOption) error { return nil }

type fakeChatStream struct {
	ctx context.Context
}

func (f *fakeChatStream) Send(*pb.ChatResponse) error { return nil }
func (f *fakeChatStream) Recv() (*pb.ChatRequest, error) {
	<-f.ctx.Done()
	return nil, f.ctx.Err()
}
func (f *fakeChatStream) SetHeader(metadata.MD) error  { return nil }
func (f *fakeChatStream) SendHeader(metadata.MD) error { return nil }
func (f *fakeChatStream) SetTrailer(metadata.MD)       {}
func (f *fakeChatStream) Context() context.Context     { return f.ctx }
func (f *fakeChatStream) SendMsg(any) error            { return nil }
func (f *fakeChatStream) RecvMsg(any) error            { return nil }

func newTestChatter(mq miface.MessageQueue) *Chatter {
	ctx := context.Background()
	c := CreateChatter(
		"uid-1",
		"local",
		"app",
		&fakeChatStream{ctx: ctx},
		zap.NewNop(),
		mq,
		0,
		nil,
	)
	c.Init()
	return c
}

func TestMakeChatTopic(t *testing.T) {
	t.Parallel()
	c := newTestChatter(&fakeMQ{})
	got := c.makeChatTopic(1, "room-9")
	wantPrefix := string(common.NatsHeader)
	if !strings.HasPrefix(got, wantPrefix) {
		t.Fatalf("topic=%q want prefix %q", got, wantPrefix)
	}
	if !strings.Contains(got, "app.local.1.room-9") {
		t.Fatalf("topic=%q missing logical parts", got)
	}
}

func TestSubscribeRetainsSubscription(t *testing.T) {
	t.Parallel()
	sub := &fakeSub{}
	c := newTestChatter(&fakeMQ{sub: sub})
	c.subscribe(&pb.ChatRequest_Subscribe{
		Destination: &pb.Destination{Channel: 1, Id: "r1"},
	})
	if len(c.subscripts) != 1 {
		t.Fatalf("subscripts=%d want 1", len(c.subscripts))
	}
}

func TestUnSubscribeRetainsOnFailure(t *testing.T) {
	t.Parallel()
	sub := &fakeSub{unsubErr: errors.New("boom")}
	c := newTestChatter(&fakeMQ{sub: sub})
	c.subscribe(&pb.ChatRequest_Subscribe{
		Destination: &pb.Destination{Channel: 2, Id: "r2"},
	})
	c.unSubscribe(&pb.ChatRequest_UnSubscribe{
		Destination: &pb.Destination{Channel: 2, Id: "r2"},
	})
	if len(c.subscripts) != 1 {
		t.Fatalf("subscripts=%d want 1 after failed unsub", len(c.subscripts))
	}
}

func TestDestroyClearsEvenWhenUnsubscribeFails(t *testing.T) {
	t.Parallel()
	calls := 0
	sub := &fakeSub{unsubErr: errors.New("boom"), calls: &calls}
	c := newTestChatter(&fakeMQ{sub: sub})
	c.subscribe(&pb.ChatRequest_Subscribe{
		Destination: &pb.Destination{Channel: 3, Id: "r3"},
	})
	c.Destroy()
	if c.subscripts != nil {
		t.Fatalf("subscripts=%v want nil after Destroy", c.subscripts)
	}
	if calls != 1 {
		t.Fatalf("Unsubscribe calls=%d want 1", calls)
	}
}
