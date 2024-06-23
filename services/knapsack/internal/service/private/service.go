package private

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"

	pb "github.com/moke-game/platform/api/gen/knapsack"
	"github.com/moke-game/platform/services/knapsack/internal/db"
	"github.com/moke-game/platform/services/knapsack/pkg/kfx"
)

type Service struct {
	utility.WithoutAuth

	logger *zap.Logger
	db     *db.Database
	mq     miface.MessageQueue
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterKnapsackPrivateServiceServer(server.GrpcServer(), s)
	return nil
}
func NewService(
	l *zap.Logger,
	coll diface.ICollection,
	mq miface.MessageQueue,
	redisCache diface.ICache,
) (result *Service, err error) {
	result = &Service{
		logger: l,
		db:     db.OpenDatabase(l, coll, redisCache),
		mq:     mq,
	}
	return
}

var Module = fx.Provide(
	func(
		logger *zap.Logger,
		dbProvider ofx.DocumentStoreParams,
		setting kfx.KnapsackSettingParams,
		mParams mfx.MessageQueueParams,
		rcParams ofx.RedisCacheParams,
	) (out sfx.GrpcServiceResult, err error) {
		if coll, e := dbProvider.DriverProvider.OpenDbDriver(setting.KnapsackStoreName); e != nil {
			err = e
		} else {
			if svc, e := NewService(
				logger,
				coll,
				mParams.MessageQueue,
				rcParams.RedisCache,
			); e != nil {
				err = e
			} else {
				out.GrpcService = svc
			}
		}
		return
	},
)
