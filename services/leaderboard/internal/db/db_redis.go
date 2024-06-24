package db

import (
	"context"
	errors2 "errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Database struct {
	redisCli *redis.Client
	maxNum   int32
	starRank int32
}

func OpenDatabase(cli *redis.Client, maxNum, starRank int32) *Database {
	return &Database{
		redisCli: cli,
		maxNum:   maxNum,
		starRank: starRank,
	}
}

func (db *Database) ClearLeaderboard(id string) error {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return err
	} else {
		return db.redisCli.Del(context.Background(), key.String()).Err()
	}
}

func (db *Database) ExpiredLeaderboard(id string, expire time.Duration) bool {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return false
	} else if res := db.redisCli.ExpireNX(context.Background(), key.String(), expire); res.Err() != nil {
		return false
	} else {
		return res.Val()
	}
}

func (db *Database) FilterWithMinScore(id string, scores map[string]float64) (map[string]float64, error) {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return make(map[string]float64), err
	} else if num := db.redisCli.ZCard(context.Background(), key.String()).Val(); num < int64(db.maxNum) {
		return scores, nil
	} else if res := db.redisCli.ZRangeByScoreWithScores(
		context.Background(),
		key.String(),
		&redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: 0, Count: 1},
	); res.Err() != nil {
		return make(map[string]float64), res.Err()
	} else {
		if len(res.Val()) == 0 {
			return scores, nil
		}
		score := res.Val()[0].Score
		filters := make(map[string]float64)
		for k, v := range scores {
			if v > score {
				filters[k] = v
			}
		}
		return filters, nil
	}
}

func (db *Database) MGetScores(id string, members []string) (map[string]float64, error) {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return make(map[string]float64), err
	} else if res := db.redisCli.ZMScore(context.Background(), key.String(), members...); res.Err() != nil {
		return make(map[string]float64), res.Err()
	} else {
		scores := make(map[string]float64)
		for k, v := range res.Val() {
			scores[members[k]] = v
		}
		return scores, nil
	}
}

func (db *Database) UpdateScore(id string, scores map[string]float64) error {
	key, err := MakeLeaderboardKey(id)
	if err != nil {
		return err
	}
	members := make([]redis.Z, 0, len(scores))
	for k, v := range scores {
		z := redis.Z{
			Member: k,
			Score:  v,
		}
		members = append(members, z)
	}
	if res := db.redisCli.ZAdd(context.Background(), key.String(), members...); res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (db *Database) GetTopWithScores(id string, start, stop int64) ([]redis.Z, error) {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return nil, err
	} else if res := db.redisCli.ZRevRangeWithScores(context.Background(), key.String(), start, stop); res.Err() != nil {
		if errors2.Is(res.Err(), redis.Nil) {
			return make([]redis.Z, 0), nil
		}
		return nil, res.Err()
	} else {
		return res.Val(), nil
	}
}

func (db *Database) GetRank(id, uid string) (redis.RankScore, error) {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return redis.RankScore{}, err
	} else if res, err := db.redisCli.ZRevRankWithScore(context.Background(), key.String(), uid).Result(); err != nil {
		return redis.RankScore{}, err
	} else {
		return res, nil
	}
}
func (db *Database) GetRankScore(id string, uid string) (redis.RankScore, error) {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return redis.RankScore{}, err
	} else if score, err := db.redisCli.ZScore(context.Background(), key.String(), uid).Result(); err != nil {
		return redis.RankScore{}, err
	} else if rank, err := db.redisCli.ZRevRank(context.Background(), key.String(), uid).Result(); err != nil {
		return redis.RankScore{}, err
	} else {
		return redis.RankScore{Rank: rank, Score: score}, nil
	}
}

func (db *Database) GetLeaderboardStars(ctx context.Context, id string, uids ...string) map[string]int64 {
	if len(uids) == 0 {
		return make(map[string]int64)
	}
	sk, err := MakeLeaderboardStarKey(id)
	if err != nil {
		return make(map[string]int64)
	}
	res, err := db.redisCli.HMGet(ctx, sk.String(), uids...).Result()
	if err != nil {
		return make(map[string]int64)
	}
	stars := make(map[string]int64)
	for i, v := range res {
		if v != nil {
			s := v.(string)
			if num, err := strconv.ParseInt(s, 10, 64); err != nil {
				stars[uids[i]] = 0
			} else {
				stars[uids[i]] = num
			}
		}
	}
	return stars
}

func (db *Database) StarLeaderboard(ctx context.Context, id string, uid string) (int64, error) {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return 0, err
	} else if rank, err := db.redisCli.ZRevRank(ctx, key.String(), uid).Result(); err != nil {
		return 0, err
	} else if rank >= int64(db.starRank) {
		return 0, fmt.Errorf("rank %d is over max star rank %d", rank, db.starRank)
	} else if sk, err := MakeLeaderboardStarKey(id); err != nil {
		return 0, err
	} else if num, err := db.redisCli.HIncrBy(ctx, sk.String(), uid, 1).Result(); err != nil {
		return 0, err
	} else {
		return num, nil
	}
}

func (db *Database) ExpireLeaderboardStar(id string, expire time.Duration) bool {
	if key, err := MakeLeaderboardStarKey(id); err != nil {
		return false
	} else if res := db.redisCli.ExpireNX(context.Background(), key.String(), expire); res.Err() != nil {
		return false
	} else {
		return res.Val()
	}
}

func (db *Database) GetLeaderboardAmount(id string) (int64, error) {
	if key, err := MakeLeaderboardKey(id); err != nil {
		return 0, err
	} else if amount, err := db.redisCli.ZCard(context.Background(), key.String()).Result(); err != nil {
		if errors2.Is(err, redis.Nil) {
			return 0, nil
		}
		return amount, err
	} else {
		return amount, nil
	}
	return 0, nil
}
