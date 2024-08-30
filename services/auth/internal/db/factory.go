package db

import (
	"errors"
	"strconv"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/nosql/key"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/auth/internal/db/model"
)

type Database struct {
	logger  *zap.Logger
	appName string
	coll    diface.ICollection
	cache   diface.ICache
}

func OpenDatabase(
	l *zap.Logger,
	appName string,
	coll diface.ICollection,
	cache diface.ICache,
) *Database {
	return &Database{
		logger:  l,
		coll:    coll,
		cache:   cache,
		appName: appName,
	}
}

const uidStart = 10000

func (db *Database) generateId() (string, error) {
	if k, err := model.NewUidKey(db.appName); err != nil {
		return "", err
	} else if uid, err := db.coll.Incr(k, "uid", 1); err != nil {
		if errors.Is(err, nerrors.ErrNotFound) {
			return strconv.FormatInt(uidStart, 10), nil
		}
		return "", err
	} else {
		return strconv.FormatInt(uidStart+uid, 10), nil
	}
}

func (db *Database) Delete(id string) (err error) {
	var index key.Key
	if index, err = model.NewAuthKey(id); err != nil {
		return
	}
	db.cache.DeleteCache(index)
	if err = db.coll.Delete(index); errors.Is(err, nerrors.ErrNotFound) {
		return nil
	}
	return
}

func (db *Database) LoadOrCreateUid(id string) (*model.Dao, error) {
	if dm, err := model.NewAuthModel(id, db.coll, db.cache); err != nil {
		return nil, err
	} else if err = dm.Load(); errors.Is(err, nerrors.ErrNotFound) {
		if uid, err := db.generateId(); err != nil {
			return nil, err
		} else if dm, err = model.NewAuthModel(id, db.coll, db.cache); err != nil {
			return nil, err
		} else if err := dm.InitDefault(uid); err != nil {
			return nil, err
		} else if err = dm.Create(); err != nil {
			if err = dm.Load(); err != nil {
				return nil, err
			} else {
				return dm, nil
			}
		} else {
			return dm, nil
		}
	} else if err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}
