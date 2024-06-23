package redis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"

	pb "github.com/moke-game/platform.git/api/gen/profile"
)

type ProfileBasic struct {
	Uid       string `json:"uid"  redis:"uid"`
	Nickname  string `json:"nickname" redis:"nickname"`
	Avatar    string `json:"avatar" redis:"avatar"`
	HeroId    int32  `json:"hero_id" redis:"hero_id" `
	HallUrl   string `json:"hall_url" redis:"hall_url"`
	BattleUrl string `json:"battle_url" redis:"battle_url"`
	RoomId    string `json:"room_id" redis:"room_id"`
}

func (p *ProfileBasic) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *ProfileBasic) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *ProfileBasic) toProto() *pb.ProfileBasic {
	return &pb.ProfileBasic{
		Uid:       p.Uid,
		Nickname:  p.Nickname,
		Avatar:    p.Avatar,
		HeroId:    p.HeroId,
		HallUrl:   p.HallUrl,
		BattleUrl: p.BattleUrl,
		RoomId:    p.RoomId,
	}
}

func GetBasicInfo(redisCli *redis.Client, uids ...string) (map[string]*pb.ProfileBasic, error) {
	pCli := redisCli.Pipeline()
	for _, uid := range uids {
		if key, err := NewProfileBasicKey(uid); err != nil {
			return nil, err
		} else {
			if err := pCli.HGetAll(context.Background(), key.String()).Err(); err != nil {
				return nil, err
			}
		}
	}
	cmds, err := pCli.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	res := make(map[string]*pb.ProfileBasic)
	for _, v := range cmds {
		info := &ProfileBasic{}
		cmd := v.(*redis.MapStringStringCmd)
		if err := cmd.Scan(info); err != nil {
			return nil, err
		}
		res[info.Uid] = info.toProto()
	}
	return res, nil
}

func UpdateBasicWithProfile(redisCli *redis.Client, uid string, profile *pb.Profile) error {
	basic := &pb.ProfileBasic{
		Uid:      uid,
		Nickname: profile.Nickname,
		Avatar:   profile.Avatar,
		HeroId:   profile.HeroId,
	}
	return SetBasicInfo(redisCli, uid, basic)
}

func SetBasicInfo(redisCli *redis.Client, uid string, basic *pb.ProfileBasic) error {
	if basic == nil {
		return nil
	}
	if key, err := NewProfileBasicKey(uid); err != nil {
		return err
	} else {
		dataMap := make(map[string]interface{})
		if basic.Uid != "" {
			dataMap["uid"] = basic.Uid
		}
		if basic.Nickname != "" {
			dataMap["nickname"] = basic.Nickname
		}

		if basic.Avatar != "" {
			dataMap["avatar"] = basic.Avatar
		}

		if basic.HeroId != 0 {
			dataMap["hero_id"] = basic.HeroId
		}

		if basic.HallUrl != "" {
			dataMap["hall_url"] = basic.HallUrl
		}

		if basic.BattleUrl != "" {
			dataMap["battle_url"] = basic.BattleUrl
		}

		if basic.RoomId != "" {
			dataMap["room_id"] = basic.RoomId
		}
		if basic.RoomHostname != "" {
			dataMap["room_hostname"] = basic.RoomHostname
		}
		return redisCli.HSet(context.Background(), key.String(), dataMap).Err()
	}

}
