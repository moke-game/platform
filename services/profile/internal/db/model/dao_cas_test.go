package model_test

import (
	"sync"
	"testing"

	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/nosql/mock"
	"go.uber.org/zap/zaptest"

	pb "github.com/moke-game/platform/api/gen/profile/api"
	"github.com/moke-game/platform/services/profile/internal/db"
)

func openProfileDB(t *testing.T) *db.Database {
	t.Helper()
	logger := zaptest.NewLogger(t)
	provider := mock.NewMockDriverProvider(logger)
	coll, err := provider.OpenDbDriver("profile-cas")
	if err != nil {
		t.Fatal(err)
	}
	return db.OpenDatabase(logger, coll, diface.DefaultDocumentCache())
}

func TestProfileCreateLoadAndUpdate(t *testing.T) {
	t.Parallel()
	database := openProfileDB(t)

	dao, err := database.CreateProfile("u1", &pb.Profile{Nickname: "alice"})
	if err != nil {
		t.Fatalf("CreateProfile: %v", err)
	}
	if err := dao.Update(func() bool {
		return dao.UpdateData(&pb.Profile{Avatar: "a.png", RechargeAmount: 10})
	}); err != nil {
		t.Fatalf("Update: %v", err)
	}

	loaded, err := database.LoadProfile("u1")
	if err != nil {
		t.Fatalf("LoadProfile: %v", err)
	}
	if loaded.Data.GetNickname() != "alice" {
		t.Fatalf("nickname=%q", loaded.Data.GetNickname())
	}
	if loaded.Data.GetAvatar() != "a.png" {
		t.Fatalf("avatar=%q", loaded.Data.GetAvatar())
	}
	if loaded.Data.GetRechargeAmount() != 10 {
		t.Fatalf("recharge=%d want 10", loaded.Data.GetRechargeAmount())
	}
}

func TestProfileConcurrentCASRecharge(t *testing.T) {
	t.Parallel()
	database := openProfileDB(t)
	if _, err := database.CreateProfile("u2", &pb.Profile{Nickname: "bob"}); err != nil {
		t.Fatalf("CreateProfile: %v", err)
	}

	const workers = 8
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make(chan error, workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			dao, err := database.LoadProfile("u2")
			if err != nil {
				errs <- err
				return
			}
			errs <- dao.Update(func() bool {
				return dao.UpdateData(&pb.Profile{RechargeAmount: 1})
			})
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatalf("concurrent Update: %v", err)
		}
	}

	loaded, err := database.LoadProfile("u2")
	if err != nil {
		t.Fatalf("LoadProfile: %v", err)
	}
	if got := loaded.Data.GetRechargeAmount(); got != workers {
		t.Fatalf("recharge=%d want %d", got, workers)
	}
}
