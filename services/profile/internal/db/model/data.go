package model

import pb "github.com/moke-game/platform/api/gen/profile/api"

func (d *Dao) initData() {
	d.Data = &pb.Profile{}
}

func (d *Dao) InitDefault(uid string, profile *pb.Profile) error {
	d.Data = profile
	d.Data.Uid = uid
	return nil
}

func (d *Dao) clear() {
	d.Data.Reset()
}

func (d *Dao) GetUid() string {
	return d.Data.GetUid()
}
