package model

import (
	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"

	"github.com/moke-game/platform.git/services/buddy/internal/db/model/data"
)

type Dao struct {
	nosql.DocumentBase `bson:"-"`
	Data               *data.BuddyQueue `bson:"data"`
}

func (b *Dao) Init(id string, ros diface.ICollection, cache diface.ICache) error {
	if ros == nil {
		return nerrors.ErrDocumentStoreIsNil
	}
	k, e := NewBuddyQueueKey(id)
	if e != nil {
		return e
	}
	b.Data = data.NewBuddyQueue(id)
	b.DocumentBase.InitWithCache(&b.Data, b.clear, ros, k, cache)
	return nil
}

func (b *Dao) clear() {
	b.Data.Clear()
}

func (b *Dao) InitDefault() error {
	return nil
}

func (b *Dao) GetBuddyDataByUID(uid string) {

}
