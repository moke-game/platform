package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const statusExpire = time.Second * 10

func SetProfileStatus(redisCli *redis.Client, profileID string, status int) error {
	if key, err := NewProfileStatusKey(profileID); err != nil {
		return err
	} else if err := redisCli.Set(context.Background(), key.String(), status, statusExpire).Err(); err != nil {
		return err
	} else {
		return nil
	}
}

func GetProfileStatus(redisCli *redis.Client, profileID ...string) map[string]int32 {
	result := make(map[string]int32)
	keys := make([]string, len(profileID))
	for i, id := range profileID {
		if key, err := NewProfileStatusKey(id); err != nil {
			continue
		} else {
			keys[i] = key.String()
		}
	}
	if res, err := redisCli.MGet(context.Background(), keys...).Result(); err != nil {
		return result
	} else {
		for i, id := range profileID {
			if res[i] == nil {
				continue
			}

			if status, err := strconv.Atoi(res[i].(string)); err != nil {
				continue
			} else {
				result[id] = int32(status)
			}
		}
	}
	return result
}
