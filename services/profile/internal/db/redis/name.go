package redis

import (
	"context"
	"fmt"

	"github.com/duke-git/lancet/v2/random"
	"github.com/redis/go-redis/v9"
)

func IsNameExist(redisCli *redis.Client, name string) (bool, error) {
	if key, err := NewNameKey(); err != nil {
		return false, err
	} else if res, err := redisCli.HExists(
		context.Background(),
		key.String(),
		name,
	).Result(); err != nil {
		return false, err
	} else {
		return res, nil
	}
}

// SaveName  TODO fix Redis String to Hash
func SaveName(redisCli *redis.Client, name string) error {
	if key, err := NewNameKey(); err != nil {
		return err
	} else if err := redisCli.HSet(context.Background(), key.String(), name, 0).Err(); err != nil {
		return err
	} else {
		return nil
	}
}

func ChangeName(redisCli *redis.Client, oldName, newName string) error {
	key, err := NewNameKey()
	if err != nil {
		return err
	}
	ctx := context.Background()
	cmd := redisCli.HSet(ctx, key.String(), newName, 0)
	ret, err := cmd.Result()
	if err != nil {
		return err
	}
	if ret != 1 {
		return fmt.Errorf("name already exists")
	}
	redisCli.HDel(ctx, key.String(), oldName)
	return nil
}

const (
	NameTryCount = 5
)

func RandomName(redisCli *redis.Client, name string) (string, error) {
	index := 0
	for index <= NameTryCount {
		rs := random.RandString(6)
		name := fmt.Sprintf("%s%s", name, rs)
		if ok, err := IsNameExist(redisCli, name); err != nil {
			return "", err
		} else if !ok {
			return name, nil
		}
		index++
	}
	return "", fmt.Errorf("failed to generate random name")
}
