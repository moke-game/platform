package db

import (
	"errors"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/knapsack/internal/db/model"
)

type Database struct {
	logger *zap.Logger
	coll   diface.ICollection
	cache  diface.ICache
}

func OpenDatabase(l *zap.Logger, coll diface.ICollection, cache diface.ICache) *Database {
	return &Database{
		logger: l,
		coll:   coll,
		cache:  cache,
	}
}

func NewKnapsackModel(id string, doc diface.ICollection, cache diface.ICache) (*model.Dao, error) {
	dm := &model.Dao{}
	if err := dm.Init(id, doc, cache); err != nil {
		return nil, err
	}
	return dm, nil
}

func (db *Database) LoadKnapsack(uid string) (*model.Dao, error) {
	if dm, err := NewKnapsackModel(uid, db.coll, db.cache); err != nil {
		return nil, err
	} else if err = dm.Load(); err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}

func (db *Database) CreateKnapsack(uid string) (*model.Dao, error) {
	if dm, err := NewKnapsackModel(uid, db.coll, db.cache); err != nil {
		return nil, err
	} else if err = dm.InitDefault(uid); err != nil {
		return nil, err
	} else if err = dm.Create(); err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}

func (db *Database) LoadOrCreateKnapsack(uid string) (*model.Dao, error) {
	if dao, err := db.LoadKnapsack(uid); err != nil {
		if errors.Is(err, nerrors.ErrNotFound) {
			return db.CreateKnapsack(uid)
		}
		return nil, err
	} else {
		return dao, nil
	}
}
