package model

import (
	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"google.golang.org/protobuf/proto"

	ppb "github.com/moke-game/platform/api/gen/profile"
)

type Dao struct {
	nosql.DocumentBase `bson:"-"`
	Data               *ppb.Profile `bson:"data"`
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

func (d *Dao) ToProto() *ppb.Profile {
	return proto.Clone(d.Data).(*ppb.Profile)
}

func (d *Dao) UpdateData(profile *ppb.Profile, updatePet bool) bool {
	if profile.Nickname != "" {
		d.Data.Nickname = profile.Nickname
	}
	if profile.HeroId != 0 {
		d.Data.HeroId = profile.HeroId
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
	if profile.Openid != "" {
		d.Data.Openid = profile.Openid
	}
	if profile.RechargeAmount > 0 {
		d.Data.RechargeAmount += profile.RechargeAmount
	}
	if profile.PetProfileId != 0 || updatePet {
		d.Data.PetProfileId = profile.PetProfileId
	}
	if profile.DeleteTime > 0 && d.Data.DeleteTime == 0 {
		d.Data.DeleteTime = profile.DeleteTime
	}
	if profile.GuideStep > 0 {
		d.Data.GuideStep = profile.GuideStep
	}
	return true
}
