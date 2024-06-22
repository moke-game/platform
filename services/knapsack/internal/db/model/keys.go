package model

import (
	"github.com/gstones/moke-kit/orm/nosql/key"
)

func NewKnapsackKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("knapsack", id)
}
