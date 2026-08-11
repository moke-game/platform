package public

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/utility"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	pb "github.com/moke-game/platform/api/gen/mail/api"
	"github.com/moke-game/platform/services/mail/internal/service/common"
	"github.com/moke-game/platform/services/mail/internal/service/db"
)

type recordingMQ struct {
	mu     sync.Mutex
	topics []string
}

func (r *recordingMQ) Subscribe(_ context.Context, topic string, _ miface.SubResponseHandler, _ ...miface.SubOption) (miface.Subscription, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.topics = append(r.topics, topic)
	return &nopSub{}, nil
}
func (r *recordingMQ) Publish(string, ...miface.PubOption) error { return nil }
func (r *recordingMQ) snapshot() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]string, len(r.topics))
	copy(out, r.topics)
	return out
}

type nopSub struct{}

func (n *nopSub) IsValid() bool         { return true }
func (n *nopSub) Unsubscribe() error    { return nil }

type fakeMailWatchStream struct {
	ctx context.Context
	mu  sync.Mutex
	n   int
}

func (f *fakeMailWatchStream) Send(*pb.WatchMailResponse) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.n++
	return nil
}
func (f *fakeMailWatchStream) SetHeader(metadata.MD) error  { return nil }
func (f *fakeMailWatchStream) SendHeader(metadata.MD) error { return nil }
func (f *fakeMailWatchStream) SetTrailer(metadata.MD)       {}
func (f *fakeMailWatchStream) Context() context.Context     { return f.ctx }
func (f *fakeMailWatchStream) SendMsg(any) error            { return nil }
func (f *fakeMailWatchStream) RecvMsg(any) error            { return nil }

func TestWatchSubscribesPrivatePublicAndChannel(t *testing.T) {
	t.Parallel()

	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(mr.Close)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	mq := &recordingMQ{}
	svc := &Service{
		logger: zap.NewNop(),
		db:     db.OpenDatabase(zap.NewNop(), rdb),
		mq:     mq,
		maxNum: 100,
	}

	ctx, cancel := context.WithCancel(context.Background())
	ctx = utility.NewContext(ctx, utility.UIDContextKey, "uid-mail")
	stream := &fakeMailWatchStream{ctx: ctx}

	done := make(chan error, 1)
	go func() {
		done <- svc.Watch(&pb.WatchMailRequest{Channel: "ios", Language: "en"}, stream)
	}()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if len(mq.snapshot()) >= 3 {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	topics := mq.snapshot()
	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Watch: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Watch did not return after cancel")
	}

	want := map[string]bool{
		common.MakePrivateTopic("uid-mail"): true,
		common.MakePublicTopic(""):          true,
		common.MakePublicTopic("ios"):       true,
	}
	if len(topics) != 3 {
		t.Fatalf("topics=%v want 3", topics)
	}
	for _, topic := range topics {
		if !want[topic] {
			t.Fatalf("unexpected topic %q in %v", topic, topics)
		}
	}
}
