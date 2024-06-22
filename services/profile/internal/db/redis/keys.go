package redis

import (
	"github.com/gstones/moke-kit/orm/nosql/key"
)

func NewNameKey() (key.Key, error) {
	return key.NewKeyFromParts("profile", "nickname")
}

func NewProfileStatusKey(uid string) (key.Key, error) {
	return key.NewKeyFromParts("profile", "status", uid)
}

func NewProfileBasicKey(uid string) (key.Key, error) {
	return key.NewKeyFromParts("profile", "basic", uid)
}
