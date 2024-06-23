package public

import (
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/moke-game/platform.git/api/gen/leaderboard"
	"github.com/moke-game/platform.git/services/leaderboard/internal/db"
	"github.com/moke-game/platform.git/services/leaderboard/pkg/lbfx"
)

type Service struct {
	logger *zap.Logger
	db     *db.Database
	maxNum int32
}

func NewService(l *zap.Logger, cli *redis.Client, setting lbfx.LeaderboardSettingParams) (result *Service, err error) {
	return &Service{
		logger: l,
		db:     db.OpenDatabase(cli, setting.MaxNum, setting.StarRank),
		maxNum: setting.MaxNum,
	}, nil
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	leaderboard.RegisterLeaderboardServiceServer(
		server.GrpcServer(),
		s,
	)
	return nil
}

var Module = fx.Provide(
	func(
		logger *zap.Logger,
		cliParams ofx.RedisParams,
		setting lbfx.LeaderboardSettingParams,
	) (sfx.GrpcServiceResult, error) {
		if s, err := NewService(logger, cliParams.Redis, setting); err != nil {
			return sfx.GrpcServiceResult{}, err
		} else {
			return sfx.GrpcServiceResult{
				GrpcService: s,
			}, nil
		}
	})
