package model

import pb "github.com/moke-game/platform/api/gen/knapsack"

func (d *Dao) initData() {
	d.Data = &pb.Knapsack{}
	d.Data.Items = make(map[int64]*pb.Item)
	d.Data.Features = make(map[int32]bool)
	d.changesCache = &pb.Knapsack{}
	d.changesCache.Items = make(map[int64]*pb.Item)
	d.changesCache.Features = make(map[int32]bool)
}

func (d *Dao) InitDefault(uid string) error {
	d.Data = &pb.Knapsack{}
	d.Data.Uid = uid
	d.Data.Items = make(map[int64]*pb.Item)
	d.Data.Features = make(map[int32]bool)
	d.changesCache = &pb.Knapsack{}
	d.changesCache.Items = make(map[int64]*pb.Item)
	d.changesCache.Features = make(map[int32]bool)
	return nil
}

func (d *Dao) clear() {
	d.Data.Reset()
}

func (d *Dao) GetUid() string {
	return d.Data.GetUid()
}
