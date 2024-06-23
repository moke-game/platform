package db

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/moke-game/platform/api/gen/matchmaking"
	"github.com/moke-game/platform/services/matchmaking/internal/db/model"
)

const messageExpire = time.Hour * 24 * 30

type Database struct {
	*redis.Client
	logger *zap.Logger
}

func OpenDatabase(l *zap.Logger, client *redis.Client) *Database {
	return &Database{
		client,
		l,
	}
}

func (db *Database) PushData(matchType, queueType, queueLevel, playId string, matchData *model.MatchData) error {
	if key, e := MakeMathKey(matchType, queueType, queueLevel, playId); e != nil {
		return e
	} else {
		if res := db.HSet(context.Background(), key.String(), matchData.Id, matchData); res.Err() != nil {
			return res.Err()
		} else {
			db.Expire(context.Background(), key.String(), messageExpire)
		}
	}
	return nil
}

func (db *Database) GetOneMatchData(id string, matchType, queueLevel, playId string) (matchData *model.MatchData, err error) {
	if val, err := db.getOneData(matchType, string(Queue_Type_Match), queueLevel, playId, id); err != nil {
		return matchData, err
	} else {
		matchData = &model.MatchData{}
		byt := []byte(val)
		if err = json.Unmarshal(byt, matchData); err != nil {
			return nil, err
		}
	}
	return matchData, err
}

func (db *Database) GetOneWaitData(id string, matchType, queueLevel, playId string) (waitData *model.MatchData, err error) {
	if val, err := db.getOneData(matchType, string(Queue_Type_Wait), queueLevel, playId, id); err != nil {
		return waitData, err
	} else {
		waitData = &model.MatchData{}
		byt := []byte(val)
		if err = json.Unmarshal(byt, waitData); err != nil {
			return nil, err
		}
	}
	return waitData, err
}

func (db *Database) GetAllMatchData(matchType, queueLevel, playId string) (matchData map[string]*model.MatchData, err error) {
	if val, err := db.getAllData(matchType, string(Queue_Type_Match), queueLevel, playId); err != nil {
		return matchData, err
	} else {
		return val, err
	}
}

func (db *Database) GetAllWaitData(matchType, queueLevel, playId string) (waitData map[string]*model.MatchData, err error) {
	if val, err := db.getAllData(matchType, string(Queue_Type_Wait), queueLevel, playId); err != nil {
		return waitData, err
	} else {
		return val, err
	}
}

func (db *Database) getOneData(matchType, queueType, queueLevel, playId string, id string) (val string, err error) {
	if key, e := MakeMathKey(matchType, queueType, queueLevel, playId); e != nil {
		return val, e
	} else {
		if val, err = db.HGet(context.Background(), key.String(), id).Result(); err != nil {
			return val, err
		}
	}
	return val, err
}

func (db *Database) getAllData(matchType, queueType, queueLevel, playId string) (allData map[string]*model.MatchData, err error) {
	if key, e := MakeMathKey(matchType, queueType, queueLevel, playId); e != nil {
		return allData, e
	} else {
		if val, err := db.HGetAll(context.Background(), key.String()).Result(); err != nil {
			return allData, err
		} else {
			allData = make(map[string]*model.MatchData)
			for _, v := range val {
				data := &model.MatchData{}
				byt := []byte(v)
				if er := json.Unmarshal(byt, data); er != nil {
					continue
				}
				allData[data.Id] = data
			}
		}
	}
	return allData, err
}

func (db *Database) HaseData(matchType, queueType, queueLevel, playId string, id string) (exist bool, err error) {
	if key, e := MakeMathKey(matchType, queueType, queueLevel, playId); e != nil {
		return false, e
	} else {
		if result := db.HExists(context.Background(), key.String(), id); result.Err() != nil {
			return false, result.Err()
		} else {
			exist = result.Val()
		}
	}
	return exist, nil
}

func (db *Database) DeleteData(matchType, queueType, queueLevel, playId string, id string) error {
	if key, e := MakeMathKey(matchType, queueType, queueLevel, playId); e != nil {
		return e
	} else {
		if _, e = db.HDel(context.Background(), key.String(), id).Result(); e != nil {
			return e
		}
		return nil
	}
}

func (db *Database) GetBattleRoom(uid string) (*matchmaking.BattleRoomData, error) {
	if key, e := MakeBattleRoomKey(uid); e != nil {
		return nil, e
	} else {
		if str, e := db.Get(context.Background(), key.String()).Result(); e != nil {
			return nil, e
		} else {
			if len(str) == 0 {
				return nil, nil
			}
			battleRoom := &matchmaking.BattleRoomData{}
			byt := []byte(str)
			if err := json.Unmarshal(byt, battleRoom); err != nil {
				return nil, err
			}
			return battleRoom, nil
		}
	}
}

func (db *Database) SaveBattleMap(uid string, playId int32, data *model.PlayerMapData) error {
	playIdStr := strconv.Itoa(int(playId))
	if key, e := MakeRoomMapKey(playIdStr); e != nil {
		return e
	} else {
		if res := db.HSet(context.Background(), key.String(), uid, data); res.Err() != nil {
			return res.Err()
		} else {
			db.Expire(context.Background(), key.String(), messageExpire)
		}
	}
	return nil
}

func (db *Database) GetBattleMap(uid string, playId int32) (*model.PlayerMapData, error) {
	playIdStr := strconv.Itoa(int(playId))
	if key, e := MakeRoomMapKey(playIdStr); e != nil {
		return nil, e
	} else {
		if str, e := db.HGet(context.Background(), key.String(), uid).Result(); e != nil {
			if errors.Is(e, redis.Nil) {
				return nil, nil
			}
			return nil, e
		} else {
			if len(str) == 0 {
				return nil, nil
			}
			battleMap := &model.PlayerMapData{}
			byt := []byte(str)
			if err := json.Unmarshal(byt, battleMap); err != nil {
				return nil, err
			}
			return battleMap, nil
		}
	}
}
