package model

import (
	"context"

	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"
)

type Dao struct {
	nosql.DocumentBase `bson:"-"`
	Data               *Data `bson:"data"`
}

func (d *Dao) Init(id string, doc diface.ICollection, cache diface.ICache) error {
	key, e := NewAuthKey(id)
	if e != nil {
		return e
	}
	d.initData()
	d.DocumentBase.InitWithCache(context.Background(), &d.Data, d.clear, doc, key, cache)
	return nil
}

func NewAuthModel(id string, doc diface.ICollection, cache diface.ICache) (*Dao, error) {
	dm := &Dao{}
	if err := dm.Init(id, doc, cache); err != nil {
		return nil, err
	}
	return dm, nil
}
