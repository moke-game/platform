package model

import ppb "github.com/gstones/platform/api/gen/profile"

func (d *Dao) initData() {
	d.Data = &ppb.Profile{}
}

func (d *Dao) InitDefault(uid string, profile *ppb.Profile) error {
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
