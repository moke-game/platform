package model

type Data struct {
	Uid string `json:"uid" bson:"uid"`
}

func createData() *Data {
	return &Data{}
}

func (d *Dao) initData() {
	data := createData()
	d.Data = data
}

func (d *Dao) InitDefault(uid string) error {
	data := createData()
	data.Uid = uid
	d.Data = data
	return nil
}

func (d *Dao) clear() {
	d.Data = nil
}

func (d *Dao) GetUid() string {
	return d.Data.Uid
}
