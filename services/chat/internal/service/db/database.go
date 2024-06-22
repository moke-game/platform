package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

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

func (db *Database) IsBlocked(uid string) (bool, error) {
	if key, err := makeBlockedListKey(uid); err != nil {
		return false, err
	} else if val := db.Exists(context.Background(), key.String()).Val(); val > 0 {
		return true, nil
	}
	return false, nil
}

func (db *Database) AddBlockedList(blockedUid string, duration int64) error {
	if key, err := makeBlockedListKey(blockedUid); err != nil {
		return err
	} else if err := db.Set(
		context.Background(),
		key.String(),
		duration,
		time.Duration(duration)*time.Second,
	).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) RemoveBlockedList(blockedUid string) error {
	if key, err := makeBlockedListKey(blockedUid); err != nil {
		return err
	} else if err := db.Del(context.Background(), key.String()).Err(); err != nil {
		return err
	}
	return nil
}
