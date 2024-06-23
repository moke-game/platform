package db

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform.git/api/gen/party"
)

const PartyExpireTime = time.Hour * 4

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

func (db *Database) ClearParty(id string) {
	if key, e := makePartyKey(id); e != nil {
		db.logger.Error("make party msg key failed", zap.Error(e), zap.String("id", id))
	} else {
		db.Del(context.Background(), key.String())
	}

	if key, e := makePartyMemberKey(id); e != nil {
		db.logger.Error("make party member key failed", zap.Error(e))
	} else {
		db.Del(context.Background(), key.String())
	}
}

func (db *Database) RemovePartyMember(id string, uid string) error {
	if key, e := makePartyKey(id); e != nil {
		return e
	} else if res := db.HDel(context.Background(), key.String(), uid); res.Err() != nil {
		return res.Err()
	} else if key, e := makePartyMemberKey(id); e != nil {
		return e
	} else if res := db.HDel(context.Background(), key.String(), uid); res.Err() != nil {
		return res.Err()
	} else if err := db.RemoveUid2Pid(uid); err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdatePartySetting(id string, msg *pb.PartySetting) error {
	if key, e := makePartyKey(id); e != nil {
		db.logger.Error("make party msg key failed", zap.Error(e))
		return e
	} else {
		update := make([]any, 0)
		if msg.Name != "" {
			update = append(update, "name", msg.Name)
		}
		if msg.Owner != "" {
			update = append(update, "owner", msg.Owner)
		}
		if msg.Type != 0 {
			update = append(update, "type", msg.Type)
		}
		if msg.MaxMember != 0 {
			update = append(update, "max_member", msg.MaxMember)
		}
		if len(msg.Refuse) > 0 {
			update = append(update, "refuse", msg.Refuse)
		}
		db.HMSet(context.Background(), key.String(), update)
	}
	return nil
}

func (db *Database) GetPartySetting(id string) (*pb.PartySetting, error) {
	partyMsg := &Party{}
	if key, e := makePartyKey(id); e != nil {
		return nil, e
	} else if cmd := db.HGetAll(context.Background(), key.String()); cmd.Err() != nil {
		return nil, cmd.Err()
	} else if len(cmd.Val()) <= 0 {
		return nil, nil
	} else if err := cmd.Scan(partyMsg); err != nil {
		return nil, err
	}
	return partyMsg.ToProto(), nil
}

func (db *Database) CreateParty(id, uid string, party *pb.PartySetting) error {
	ctx := context.Background()
	party.Owner = uid
	if key, e := makePartyKey(id); e != nil {
		db.logger.Error("make party msg key failed", zap.Error(e))
		return e
	} else if !db.HSetNX(ctx, key.String(), "id", id).Val() {
		return nil
	} else if err := db.UpdatePartySetting(id, party); err != nil {
		return err
	} else if err := db.Expire(ctx, key.String(), PartyExpireTime).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetPartyMaxNumber(id string) (int, error) {
	if key, e := makePartyKey(id); e != nil {
		return 0, e
	} else if res := db.HGet(context.Background(), key.String(), "max_member"); res.Err() != nil {
		return 0, res.Err()
	} else {
		return res.Int()
	}
}

func (db *Database) GetPartyOwner(id string) (string, error) {
	if key, e := makePartyKey(id); e != nil {
		return "", e
	} else if res := db.HGet(context.Background(), key.String(), "owner"); res.Err() != nil {
		return "", res.Err()
	} else {
		return res.Val(), nil
	}
}

func (db *Database) GetPartyMemberNum(partyId string) (int32, error) {
	ctx := context.Background()
	if key, e := makePartyMemberKey(partyId); e != nil {
		return 0, e
	} else if res := db.HLen(ctx, key.String()); res.Err() != nil {
		return 0, res.Err()
	} else {
		return int32(res.Val()), nil
	}
}

func (db *Database) AddPartyMember(partyId string, uid string, member *pb.Member) error {
	ctx := context.Background()
	if key, e := makePartyMemberKey(partyId); e != nil {
		return e
	} else if data, err := json.Marshal(member); err != nil {
		return err
	} else if err := db.HSet(ctx, key.String(), uid, data).Err(); err != nil {
		return err
	} else if err := db.Expire(ctx, key.String(), PartyExpireTime).Err(); err != nil {
		return err
	} else if err := db.SaveUid2Pid(uid, partyId); err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdatePartyMember(partyId string, member *pb.Member) error {
	if key, e := makePartyMemberKey(partyId); e != nil {
		return e
	} else if data, err := db.HGet(context.Background(), key.String(), member.Uid).Bytes(); err != nil {
		return err
	} else {
		old := &pb.Member{}
		if err := json.Unmarshal(data, old); err != nil {
			return err
		}
		if member.HeroId != 0 {
			old.HeroId = member.HeroId
		}
		if member.Status != 0 {
			old.Status = member.Status
		}
		if sData, err := json.Marshal(old); err != nil {
			return err
		} else if res := db.HSet(context.Background(), key.String(), member.Uid, sData); res.Err() != nil {
			return res.Err()
		}
	}
	return nil
}

func (db *Database) GetPartyMembers(partyId string) (map[string]*pb.Member, error) {
	members := make(map[string]*pb.Member)
	if key, e := makePartyMemberKey(partyId); e != nil {
		return nil, e
	} else if res := db.HGetAll(context.Background(), key.String()); res.Err() != nil {
		return nil, res.Err()
	} else {
		for k, v := range res.Val() {
			m := &pb.Member{}
			if err := json.Unmarshal([]byte(v), m); err != nil {
				continue
			}
			members[k] = m
		}
	}
	return members, nil
}

func (db *Database) SaveUid2Pid(uid string, pid string) error {
	if key, err := makeUid2PidKey(); err != nil {
		return err
	} else if res := db.HSet(context.Background(), key.String(), uid, pid); res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (db *Database) GetUid2Pid(uid string) (string, error) {
	if key, err := makeUid2PidKey(); err != nil {
		return "", err
	} else if res := db.HGet(context.Background(), key.String(), uid); res.Err() != nil {
		return "", res.Err()
	} else {
		return res.Val(), nil
	}
}

func (db *Database) RemoveUid2Pid(uid string) error {
	if key, err := makeUid2PidKey(); err != nil {
		return err
	} else if res := db.HDel(context.Background(), key.String(), uid); res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (db *Database) GetInvite(uid string) (*PartyInvite, error) {
	ctx := context.Background()
	key, e := makeInviteKey(uid)
	if e != nil {
		db.logger.Error("get invite msg key failed", zap.Error(e))
		return nil, e
	}
	result, err := db.Get(ctx, key.String()).Result()
	if err != nil {
		if err == redis.Nil {
			invite := &PartyInvite{}
			return invite, nil
		}
		db.logger.Error("get invite failed", zap.Error(e))
		return nil, e
	}
	invite := &PartyInvite{}
	if err = json.Unmarshal([]byte(result), invite); err != nil {
		db.logger.Error("Unmarshal invite failed", zap.Error(e))
		return nil, e
	}
	return invite, nil
}

func (db *Database) SaveInvite(invite *PartyInvite) error {
	ctx := context.Background()
	key, e := makeInviteKey(invite.Id)
	if e != nil {
		db.logger.Error("save invite msg key failed", zap.Error(e))
		return e
	}
	status := db.Set(ctx, key.String(), invite, PartyExpireTime)
	if status.Err() != nil {
		db.logger.Error("save invite failed", zap.Error(e))
		return status.Err()
	}
	return nil
}

func (db *Database) GetPartyId() (string, error) {
	ctx := context.Background()
	key, e := makePartyIdKey()
	if e != nil {
		db.logger.Error("get invite msg key failed", zap.Error(e))
		return "", e
	}
	result, err := db.Get(ctx, key.String()).Result()
	val := int64(0)
	if err != nil || len(result) == 0 {
		val, err = db.IncrBy(ctx, key.String(), 1000000).Result()
		if err != nil {
			db.logger.Error("IncrBy  key failed", zap.Error(e))
			return "", e
		}
	} else {
		val, err = db.Incr(ctx, key.String()).Result()
		if err != nil {
			db.logger.Error("IncrBy  key failed", zap.Error(e))
			return "", e
		}
	}
	partyId := strconv.FormatInt(val, 16)
	return partyId, nil
}
