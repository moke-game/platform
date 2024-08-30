package model

import "github.com/gstones/moke-kit/orm/nosql/key"

func NewAuthKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("auth", id)
}

func NewUidKey(appName string) (key.Key, error) {
	return key.NewKeyFromParts(appName, "uid")
}
