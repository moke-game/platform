package public

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"

	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"

	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/profile/internal/db"
	"github.com/moke-game/platform/services/profile/pkg/pfx"
)

type Service struct {
	logger         *zap.Logger
	db             *db.Database
	redisCli       *redis.Client
	mongoCli       *mongo.Client
	mq             miface.MessageQueue
	authMiddleware siface.IAuthMiddleware
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterProfileServiceServer(server.GrpcServer(), s)
	return nil
}

func NewService(
	l *zap.Logger,
	coll diface.ICollection,
	client *redis.Client,
	redisCache diface.ICache,
	mongoCli *mongo.Client,
	mq miface.MessageQueue,
	authMiddleware siface.IAuthMiddleware,
) (result *Service, err error) {
	result = &Service{
		logger:         l,
		db:             db.OpenDatabase(l, coll, redisCache),
		redisCli:       client,
		mongoCli:       mongoCli,
		mq:             mq,
		authMiddleware: authMiddleware,
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
		dbParams ofx.MongoParams,
		mqParams mfx.MessageQueueParams,
		authMiddlewareParams sfx.AuthMiddlewareParams,
	) (out sfx.GrpcServiceResult, err error) {
		if coll, e := dbProvider.DriverProvider.OpenDbDriver(pSetting.ProfileStoreName); e != nil {
			err = e
		} else {
			if svc, e := NewService(
				l,
				coll,
				redisParams.Redis,
				rcParams.RedisCache,
				dbParams.MongoClient,
				mqParams.MessageQueue,
				authMiddlewareParams.AuthMiddleware,
			); e != nil {
				err = e
			} else {
				out.GrpcService = svc
			}
		}
		return
	},
)
