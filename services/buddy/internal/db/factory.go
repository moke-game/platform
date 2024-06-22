package db

import (
	"errors"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"go.uber.org/zap"

	"github.com/gstones/platform/services/buddy/internal/db/model"
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

func (d *Database) NewBuddyQueue(id string) (*model.Dao, error) {
	bq := new(model.Dao)
	err := bq.Init(id, d.coll, d.cache)
	if err != nil {
		return nil, err
	}
	return bq, nil
}

func (d *Database) CreateBuddyQueue(id string) error {
	if bq, err := d.NewBuddyQueue(id); err != nil {
		return err
	} else if err = bq.Create(); err != nil {
		return err
	}
	return nil
}

func (d *Database) LoadOrCreateBuddyQueue(id string) (*model.Dao, error) {
	if bq, err := d.NewBuddyQueue(id); err != nil {
		return nil, err
	} else if err := bq.Load(); errors.Is(err, nerrors.ErrKeyNotFound) {
		if bq, err := d.NewBuddyQueue(id); err != nil {
			return nil, err
		} else if err := bq.InitDefault(); err != nil {
			return nil, err
		} else if err := bq.Create(); err != nil {
			if err = bq.Load(); err != nil {
				return nil, err
			}
		} else {
			return bq, err
		}
	} else {
		return bq, nil
	}
	return nil, nil

}
