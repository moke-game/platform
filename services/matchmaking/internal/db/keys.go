package db

import (
	"github.com/gstones/moke-kit/orm/nosql/key"
)

type MatchType string
type QueueType string
type QueueLevelType string

// 匹配类型
const (
	Match_Type_Single MatchType = "single"
	Match_Type_Team   MatchType = "team"
	Match_Type_PVE    MatchType = "pve"
	Match_Type_Guid   MatchType = "guid"
)

// 队列类型
const (
	Queue_Type_Match QueueType = "match"
	Queue_Type_Wait  QueueType = "wait"
)

// 队列等级
const (
	Queue_Level_Zero QueueLevelType = "0" //等级1 10秒以下 暂定
	Queue_Level_One  QueueLevelType = "1" //等级2 20秒以下 暂定
	Queue_Level_Two  QueueLevelType = "2" //等级3 30秒以下 暂定
)

func MakeMathKey(matchType string, queueType string, queueLevel string, playId string) (key.Key, error) {
	return key.NewKeyFromParts("match", matchType, queueType, queueLevel, playId)
}

func MakeBattleRoomKey(uid string) (key.Key, error) {
	return key.NewKeyFromParts("player", "battle", "room", uid)
}

func MakeRoomKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("battle", "room", id)
}

func MakeRoomMapKey(playId string) (key.Key, error) {
	return key.NewKeyFromParts("battle", "map", playId)
}
