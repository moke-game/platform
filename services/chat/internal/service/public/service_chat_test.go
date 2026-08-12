package public

import (
	"context"
	"testing"
	"time"

	"github.com/gstones/moke-kit/utility"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	pb "github.com/moke-game/platform/api/gen/chat/api"
)

type scriptedChatStream struct {
	ctx  context.Context
	msgs []*pb.ChatRequest
	i    int
}

func (s *scriptedChatStream) Send(*pb.ChatResponse) error { return nil }
func (s *scriptedChatStream) Recv() (*pb.ChatRequest, error) {
	if s.i < len(s.msgs) {
		m := s.msgs[s.i]
		s.i++
		return m, nil
	}
	<-s.ctx.Done()
	return nil, s.ctx.Err()
}
func (s *scriptedChatStream) SetHeader(metadata.MD) error  { return nil }
func (s *scriptedChatStream) SendHeader(metadata.MD) error { return nil }
func (s *scriptedChatStream) SetTrailer(metadata.MD)       {}
func (s *scriptedChatStream) Context() context.Context     { return s.ctx }
func (s *scriptedChatStream) SendMsg(any) error            { return nil }
func (s *scriptedChatStream) RecvMsg(any) error            { return nil }

func TestChatWaitsForDestroyBeforeReturn(t *testing.T) {
	t.Parallel()

	unsubCalls := 0
	mq := &fakeMQ{sub: &fakeSub{calls: &unsubCalls}}
	svc := &Service{
		logger:       zap.NewNop(),
		mq:           mq,
		appId:        "app",
		deployment:   "local",
		chatInterval: 0,
	}

	ctx, cancel := context.WithCancel(context.Background())
	ctx = utility.NewContext(ctx, utility.UIDContextKey, "uid-chat")
	stream := &scriptedChatStream{
		ctx: ctx,
		msgs: []*pb.ChatRequest{{
			Kind: &pb.ChatRequest_Subscribe_{
				Subscribe: &pb.ChatRequest_Subscribe{
					Destination: &pb.Destination{Channel: 1, Id: "room"},
				},
			},
		}},
	}

	done := make(chan error, 1)
	go func() {
		done <- svc.Chat(stream)
	}()

	// Allow Subscribe to be processed, then cancel so Update returns and Destroy runs.
	time.Sleep(50 * time.Millisecond)
	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Chat: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Chat did not return")
	}

	// Chat returns only after Update's deferred Destroy; Unsubscribe must have run.
	if unsubCalls != 1 {
		t.Fatalf("Unsubscribe calls=%d want 1 (Destroy before Chat return)", unsubCalls)
	}
}
