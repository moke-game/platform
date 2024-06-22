package lbfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

type LeaderboardSettingParams struct {
	fx.In
	Url string `name:"leaderboardUrl"`
	// leaderboard expired time in days
	Expire int64 `name:"leaderboardExpire"`
	// leaderboard max number
	MaxNum int32 `name:"leaderboardMaxNum"`
	// can star the leaderboard rank
	StarRank int32 `name:"leaderboardStarRank"`
}

type LeaderboardSettingResult struct {
	fx.Out

	Url string `name:"leaderboardUrl" envconfig:"LEADERBOARD_URL" default:"localhost:8081"`
	// Expire is the expired time of the leaderboard in days
	Expire int64 `name:"leaderboardExpire" envconfig:"LEADERBOARD_EXPIRE" default:"30"`
	// MaxNum is the max number of the leaderboard
	MaxNum int32 `name:"leaderboardMaxNum" envconfig:"LEADERBOARD_MAX_NUM" default:"5000"`
	// StarRank is the rank that can be starred
	StarRank int32 `name:"leaderboardStarRank" envconfig:"LEADERBOARD_STAR_RANK" default:"3"`
}

func (g *LeaderboardSettingResult) LoadFromEnv() (err error) {
	err = utility.Load(g)
	return
}

var LeaderboardSettingModule = fx.Provide(
	func() (out LeaderboardSettingResult, err error) {
		err = out.LoadFromEnv()
		return
	})
