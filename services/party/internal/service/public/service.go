package public

import (
	mfx2 "github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb "github.com/gstones/platform/api/gen/party"
	"github.com/gstones/platform/services/party/internal/db"
	"github.com/gstones/platform/services/party/pkg/ptfx"
)

type Service struct {
	logger     *zap.Logger
	mq         miface.MessageQueue
	redis      *redis.Client
	appId      string
	deployment string
	db         *db.Database
}

func NewService(
	l *zap.Logger,
	rClient *redis.Client,
	mq miface.MessageQueue,
	deployment string,
	appId string,
) (result *Service, err error) {
	result = &Service{
		logger:     l,
		redis:      rClient,
		mq:         mq,
		appId:      appId,
		deployment: deployment,
		db:         db.OpenDatabase(l, rClient),
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterPartyServiceServer(server.GrpcServer(), s)
	return nil
}

var PartyService = fx.Provide(
	func(
		l *zap.Logger,
		setting ptfx.PartySettingParams,
		mqParams mfx.MessageQueueParams,
		redisParams ofx.RedisParams,
		aParams mfx2.AppParams,
	) (out sfx.GrpcServiceResult, err error) {
		if s, err := NewService(
			l,
			redisParams.Redis,
			mqParams.MessageQueue,
			aParams.Deployment,
			aParams.AppId,
		); err != nil {
			return out, err
		} else {
			out.GrpcService = s
		}
		return
	},
)
