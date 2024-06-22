package model

import (
	"encoding/json"
)

// MatchData 匹配数据
type MatchData struct {
	Id        string                 `json:"id"`        //匹配数据ID
	Members   map[string]*PlayerData `json:"members"`   //队伍成员
	GroupSize int32                  `json:"groupSize"` //队伍所需人数
	Score     int32                  `json:"score"`     //队伍中英雄熟练度等级平均值
	PlayId    int32                  `json:"playId"`    //玩法ID
	MatchTime int64                  `json:"matchTime"` //开始匹配时间戳
	MapId     []int32                `json:"mapId"`     //玩法可选地图
}

// PlayerData 玩家简要信息
type PlayerData struct {
	Uid             string            `json:"uid"` //玩家唯一ID
	Nickname        string            `json:"nickname"`
	Avatar          string            `json:"avatar"`
	HeroId          int32             `json:"hero_id"` //英雄id 配置ID
	HeroLevel       int32             `json:"hero_level"`
	HeroCups        int32             `json:"hero_cups"` // 英雄杯数
	SkinId          int32             `json:"skin_id"`   //皮肤ID
	Attribute       map[int32]float64 `json:"attribute"`
	Score           int32             `json:"score"`          //英雄熟练度
	PetProfileId    int64             `json:"pet_profile_id"` //宠物外观ID
	PetAddAttribute map[int32]float64 `json:"pet_add_attribute"`
	PetSkill        map[int32]int32   `json:"pet_skill"` //宠物技能
	IsAgain         bool              `json:"is_again"`  //是否再次匹配
}

// PlayerMapData 玩家地图信息
type PlayerMapData struct {
	Uid      string `json:"uid"`
	MapId    int32  `json:"map_id"`    //上一次使用的地图ID
	MapCount int32  `json:"map_count"` //地图使用的次数
}

func (m *MatchData) AddMember(uid string, data *PlayerData) {
	m.Members[uid] = data
}

func (m *MatchData) AppendMember(other *MatchData) {
	for pid, data := range other.Members {
		m.AddMember(pid, data)
	}
}

func (m *MatchData) RemoveMember(uid string) {
	delete(m.Members, uid)
}

func (m *MatchData) UpdateScore() {
	var score int32 = 0
	for _, data := range m.Members {
		score += data.Score
	}
	m.Score = score / int32(len(m.Members))
}

func (m *MatchData) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (h *PlayerData) MarshalBinary() ([]byte, error) {
	return json.Marshal(h)
}

func (p *PlayerMapData) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}
