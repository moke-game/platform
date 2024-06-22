package db

import "github.com/gstones/moke-kit/orm/nosql/key"

func MakeLeaderboardKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("leaderboard", id)
}

func MakeLeaderboardStarKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("leaderboard", "star", id)
}

func MakeLeaderboardStarSelfKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("leaderboard", "star", "self", id)
}
