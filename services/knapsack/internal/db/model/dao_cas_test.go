package model_test

import (
	"sync"
	"testing"

	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/nosql/mock"
	"go.uber.org/zap/zaptest"

	pb "github.com/moke-game/platform/api/gen/knapsack/api"
	"github.com/moke-game/platform/services/knapsack/internal/db"
)

func openTestDB(t *testing.T) *db.Database {
	t.Helper()
	logger := zaptest.NewLogger(t)
	provider := mock.NewMockDriverProvider(logger)
	coll, err := provider.OpenDbDriver("knapsack-cas")
	if err != nil {
		t.Fatal(err)
	}
	return db.OpenDatabase(logger, coll, diface.DefaultDocumentCache())
}

func TestKnapsackCreateLoadAndUpdate(t *testing.T) {
	t.Parallel()
	database := openTestDB(t)

	dao, err := database.CreateKnapsack("u1")
	if err != nil {
		t.Fatalf("CreateKnapsack: %v", err)
	}
	if err := dao.Update(func() bool {
		dao.AddItems(map[int64]*pb.Item{1001: {Id: 1001, Num: 3}})
		return true
	}); err != nil {
		t.Fatalf("Update: %v", err)
	}

	loaded, err := database.LoadKnapsack("u1")
	if err != nil {
		t.Fatalf("LoadKnapsack: %v", err)
	}
	if got := loaded.Data.Items[1001].GetNum(); got != 3 {
		t.Fatalf("item num=%d want 3", got)
	}
}

func TestKnapsackConcurrentCASAdds(t *testing.T) {
	t.Parallel()
	database := openTestDB(t)
	if _, err := database.CreateKnapsack("u2"); err != nil {
		t.Fatalf("CreateKnapsack: %v", err)
	}

	const workers = 8
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make(chan error, workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			dao, err := database.LoadKnapsack("u2")
			if err != nil {
				errs <- err
				return
			}
			errs <- dao.Update(func() bool {
				dao.AddItems(map[int64]*pb.Item{2002: {Id: 2002, Num: 1}})
				return true
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

	loaded, err := database.LoadKnapsack("u2")
	if err != nil {
		t.Fatalf("LoadKnapsack: %v", err)
	}
	if got := loaded.Data.Items[2002].GetNum(); got != workers {
		t.Fatalf("item num=%d want %d", got, workers)
	}
}

func TestKnapsackRemoveFailureDoesNotPersist(t *testing.T) {
	t.Parallel()
	database := openTestDB(t)
	dao, err := database.CreateKnapsack("u3")
	if err != nil {
		t.Fatalf("CreateKnapsack: %v", err)
	}
	if err := dao.Update(func() bool {
		dao.AddItems(map[int64]*pb.Item{3003: {Id: 3003, Num: 2}})
		return true
	}); err != nil {
		t.Fatalf("seed Update: %v", err)
	}

	dao, err = database.LoadKnapsack("u3")
	if err != nil {
		t.Fatalf("LoadKnapsack: %v", err)
	}
	err = dao.Update(func() bool {
		if err := dao.RemoveItems(map[int64]*pb.Item{3003: {Id: 3003, Num: 99}}); err != nil {
			return false
		}
		return true
	})
	if err == nil {
		t.Fatal("expected Update logic failure")
	}

	loaded, err := database.LoadKnapsack("u3")
	if err != nil {
		t.Fatalf("reload: %v", err)
	}
	if got := loaded.Data.Items[3003].GetNum(); got != 2 {
		t.Fatalf("item num=%d want 2 (unchanged)", got)
	}
}
