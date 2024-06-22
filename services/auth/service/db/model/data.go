package model

type Data struct {
	Uid string `json:"uid" bson:"uid"`
	Pid string `json:"pid" bson:"pid,omitempty"` //游客auth中 保存用户的openID
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
