package model

import (
	"github.com/gstones/moke-kit/orm/nosql/key"
)

func NewProfileKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("profile", id)
}

func NewProfileCollectionName() (key.Key, error) {
	return key.NewKeyFromParts("profile")
}
