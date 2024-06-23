package private

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"

	pb "github.com/moke-game/platform.git/api/gen/profile"
	"github.com/moke-game/platform.git/services/profile/pkg/pfx"
)

type Service struct {
	utility.WithoutAuth
	url string

	//profileDb  *db.Database
	logger   *zap.Logger
	redisCli *redis.Client
	mq       miface.MessageQueue
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterProfilePrivateServiceServer(server.GrpcServer(), s)
	return nil
}

func NewService(
	l *zap.Logger,
	url string,
	client *redis.Client,
	mq miface.MessageQueue,
) (result *Service, err error) {
	result = &Service{
		logger:   l,
		redisCli: client,
		url:      url,
		mq:       mq,
	}
	return
}

var Module = fx.Provide(
	func(
		l *zap.Logger,
		pSetting pfx.ProfileSettingParams,
		dbProvider ofx.DocumentStoreParams,
		redisParams ofx.RedisParams,
		rcParams ofx.RedisCacheParams,
		mongoParams ofx.MongoParams,
		mqParams mfx.MessageQueueParams,
	) (out sfx.GrpcServiceResult, err error) {
		if svc, e := NewService(
			l,
			pSetting.ProfileUrl,
			redisParams.Redis,
			mqParams.MessageQueue,
		); e != nil {
			err = e
		} else {
			out.GrpcService = svc
		}
		return
	},
)
