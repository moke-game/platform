package db

import "github.com/gstones/moke-kit/orm/nosql/key"

func makeBlockedListKey(uid string) (key.Key, error) {
	return key.NewKeyFromParts("chat", "blocked", uid)
}
