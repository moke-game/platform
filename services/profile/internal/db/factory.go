package db

import (
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/orm/nosql/diface"

	pb "github.com/moke-game/platform.git/api/gen/profile"
	"github.com/moke-game/platform.git/services/profile/internal/db/model"
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

func NewProfileModel(id string, doc diface.ICollection, cache diface.ICache) (*model.Dao, error) {
	dm := &model.Dao{}
	if err := dm.Init(id, doc, cache); err != nil {
		return nil, err
	}
	return dm, nil
}

func (db *Database) LoadProfile(uid string) (*model.Dao, error) {
	if dm, err := NewProfileModel(uid, db.coll, db.cache); err != nil {
		return nil, err
	} else if err = dm.Load(); err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}

func (db *Database) CreateProfile(uid string, profile *pb.Profile) (*model.Dao, error) {
	if dm, err := NewProfileModel(uid, db.coll, db.cache); err != nil {
		return nil, err
	} else if err = dm.InitDefault(uid, profile); err != nil {
		return nil, err
	} else if err = dm.Create(); err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}
