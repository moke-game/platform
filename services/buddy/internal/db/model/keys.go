package model

import "github.com/gstones/moke-kit/orm/nosql/key"

func NewBuddyQueueKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("buddy", id)
}
