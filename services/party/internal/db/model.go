package db

import (
	"encoding/json"

	pb "github.com/moke-game/platform/api/gen/party"
)

type Party struct {
	Id        string           `redis:"id"`
	Type      int32            `redis:"type"`
	Owner     string           `redis:"owner"`
	Name      string           `redis:"name"`
	MaxMember int32            `redis:"max_member"`
	Refuse    map[string]int64 `redis:"refuse"` //拒绝列表 key:申请的玩家ID value:拒绝时间戳 秒
}

type PartyInvite struct {
	Id         string           `redis:"id"`
	Inviter    string           `redis:"Inviter"`     //邀请玩家ID
	InviteTime int64            `redis:"invite_time"` //邀请时间戳 秒
	Refuse     map[string]int64 `redis:"refuse"`      //拒绝列表 key:邀请玩家ID value:拒绝时间戳 秒
}

func (p *Party) ToProto() *pb.PartySetting {
	return &pb.PartySetting{
		Id:        p.Id,
		Owner:     p.Owner,
		Name:      p.Name,
		Type:      p.Type,
		MaxMember: p.MaxMember,
		Refuse:    p.Refuse,
	}
}

func (p *Party) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (pi *PartyInvite) MarshalBinary() ([]byte, error) {
	return json.Marshal(pi)
}
