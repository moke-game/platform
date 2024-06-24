package private

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"

	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/profile/internal/db"
	"github.com/moke-game/platform/services/profile/internal/db/model"
	"github.com/moke-game/platform/services/profile/pkg/pfx"
)

type Service struct {
	utility.WithoutAuth
	url string

	logger     *zap.Logger
	redisCli   *redis.Client
	mq         miface.MessageQueue
	privateDao *model.PrivateDao
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
	mongoDB *mongo.Database,
) (result *Service, err error) {
	pd, err := db.NewProfilePrivateDao(mongoDB)
	if err != nil {
		return nil, err
	}
	result = &Service{
		logger:     l,
		redisCli:   client,
		url:        url,
		mq:         mq,
		privateDao: pd,
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
			mongoParams.MongoClient.Database(pSetting.ProfileStoreName),
		); e != nil {
			err = e
		} else {
			out.GrpcService = svc
		}
		return
	},
)
