package model

import (
	"fmt"
	"time"

	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"google.golang.org/protobuf/proto"

	pb "github.com/moke-game/platform.git/api/gen/knapsack"
)

type Dao struct {
	nosql.DocumentBase `bson:"-"`
	Data               *pb.Knapsack `bson:"data"`
	changesCache       *pb.Knapsack `bson:"-"`
}

func (d *Dao) Init(id string, doc diface.ICollection, cache diface.ICache) error {
	key, e := NewKnapsackKey(id)
	if e != nil {
		return e
	}
	d.initData()
	d.DocumentBase.InitWithCache(&d.Data, d.clear, doc, key, cache)
	return nil
}

func (d *Dao) ToProto() *pb.Knapsack {
	return proto.Clone(d.Data).(*pb.Knapsack)
}

func (d *Dao) GetAndDeleteChanges(incrItems, decrItems map[int64]*pb.Item, source string) *pb.KnapsackModify {
	if d.changesCache == nil {
		return nil
	}
	changes := proto.Clone(d.changesCache).(*pb.Knapsack)
	d.changesCache.Reset()
	d.changesCache.Items = make(map[int64]*pb.Item)
	d.changesCache.Features = make(map[int32]bool)
	return &pb.KnapsackModify{
		IncrItems: incrItems,
		DecrItems: decrItems,
		Knapsack:  changes,
		Source:    source,
	}
}

func (d *Dao) AddFeatures(features map[int32]bool) {
	if d.Data.Features == nil {
		d.Data.Features = make(map[int32]bool)
	}
	if d.changesCache.Features == nil {
		d.changesCache.Features = make(map[int32]bool)
	}
	for k, v := range features {
		d.Data.Features[k] = v
	}
	d.changesCache.Features = d.Data.Features
}

func (d *Dao) AddItems(items map[int64]*pb.Item) {
	if d.Data.Items == nil {
		d.Data.Items = map[int64]*pb.Item{}
	}
	for k, v := range items {
		if _, ok := d.Data.Items[k]; ok {
			d.Data.Items[k].Num += v.Num
		} else {
			d.Data.Items[k] = v
		}
		d.changesCache.Items[k] = d.Data.Items[k]
	}
}

func (d *Dao) RemoveItems(items map[int64]*pb.Item) error {
	nowTime := time.Now().UTC().Unix()
	for k, v := range items {
		if _, ok := d.Data.Items[k]; ok {
			if d.Data.Items[k].Expire > 0 && d.Data.Items[k].Expire < nowTime {
				return fmt.Errorf("item %d expire is %d", k, d.Data.Items[k].Expire)
			}
			if d.Data.Items[k].Num < v.Num {
				return fmt.Errorf("item %d not enough, need:%d  has:%d", k, v.Num, d.Data.Items[k].Num)
			}
			d.Data.Items[k].Num -= v.Num
			d.changesCache.Items[k] = d.Data.Items[k]
		} else {
			return fmt.Errorf("item %d not found", k)
		}
	}
	return nil
}
