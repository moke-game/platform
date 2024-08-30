package redis

import (
	"github.com/gstones/moke-kit/orm/nosql/key"
)

func NewBlockListKey(uid string) (key.Key, error) {
	return key.NewKeyFromParts("auth", "blocked", uid)
}

func NewTokenKey(token string) (key.Key, error) {
	return key.NewKeyFromParts("auth", "token", token)
}
