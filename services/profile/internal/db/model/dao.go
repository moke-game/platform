package model

import (
	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"google.golang.org/protobuf/proto"

	pb "github.com/moke-game/platform/api/gen/profile/api"
)

type Dao struct {
	nosql.DocumentBase `bson:"-"`
	Data               *pb.Profile `bson:"data"`
}

func (d *Dao) Init(id string, doc diface.ICollection, cache diface.ICache) error {
	key, e := NewProfileKey(id)
	if e != nil {
		return e
	}
	d.initData()
	d.DocumentBase.InitWithCache(&d.Data, d.clear, doc, key, cache)
	return nil
}

func (d *Dao) ToProto() *pb.Profile {
	return proto.Clone(d.Data).(*pb.Profile)
}

func (d *Dao) UpdateData(profile *pb.Profile) bool {
	if profile.Nickname != "" {
		d.Data.Nickname = profile.Nickname
	}

	if profile.Avatar != "" {
		d.Data.Avatar = profile.Avatar
	}
	if profile.Phone != "" {
		d.Data.Phone = profile.Phone
	}
	if profile.Email != "" {
		d.Data.Email = profile.Email
	}

	if profile.RechargeAmount > 0 {
		d.Data.RechargeAmount += profile.RechargeAmount
	}
	return true
}
