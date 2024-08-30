package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func SaveAuthToken(redisCli *redis.Client, uid, token string, duration time.Duration) error {
	if key, err := NewTokenKey(uid); err != nil {
		return err
	} else if err := redisCli.Set(context.Background(), key.String(), token, duration).Err(); err != nil {
		return err
	}
	return nil
}

func IsAuthTokenSame(redisCli *redis.Client, uid, token string) (bool, error) {
	if key, err := NewTokenKey(uid); err != nil {
		return false, err
	} else if res := redisCli.Get(context.Background(), key.String()); res.Err() != nil {
		return false, res.Err()
	} else {
		return res.Val() == token, nil
	}
}

func IsAuthTokenExist(redisCli *redis.Client, uid string) (bool, error) {
	if key, err := NewTokenKey(uid); err != nil {
		return false, err
	} else if res := redisCli.Exists(context.Background(), key.String()); res.Err() != nil {
		return false, res.Err()
	} else {
		return res.Val() > 0, nil
	}
}

func BlockedProfile(redisCli *redis.Client, profileID string, duration time.Duration) error {
	if key, err := NewBlockListKey(profileID); err != nil {
		return err
	} else if err := redisCli.Set(context.Background(), key.String(), profileID, duration).Err(); err != nil {
		return err
	}
	return nil
}

func UnBlockedProfile(redisCli *redis.Client, profileID string) error {
	if key, err := NewBlockListKey(profileID); err != nil {
		return err
	} else if err := redisCli.Del(context.Background(), key.String()).Err(); err != nil {
		return err
	}
	return nil
}

func IsBlocked(redisCli *redis.Client, profileID string) (bool, error) {
	if key, err := NewBlockListKey(profileID); err != nil {
		return false, err
	} else if res := redisCli.Exists(context.Background(), key.String()); res.Err() != nil {
		return false, res.Err()
	} else {
		return res.Val() > 0, nil
	}
}

func ClearAuthToken(redisCli *redis.Client, uid string) error {
	if key, err := NewTokenKey(uid); err != nil {
		return err
	} else if err := redisCli.Del(context.Background(), key.String()).Err(); err != nil {
		return err
	}
	return nil
}
