package service

import (
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"

	pb "github.com/gstones/platform/api/gen/buddy"
	"github.com/gstones/platform/services/buddy/internal/db"
	"github.com/gstones/platform/services/buddy/pkg/bfx"
)

type Service struct {
	logger     *zap.Logger
	db         *db.Database
	mq         miface.MessageQueue
	maxInviter int32
	maxBuddies int32
	maxBlocked int32
}

func NewService(
	l *zap.Logger,
	coll diface.ICollection,
	cache diface.ICache,
	mq miface.MessageQueue,
	setting bfx.BuddySettingsParams,
) (result *Service, err error) {
	result = &Service{
		logger:     l,
		db:         db.OpenDatabase(l, coll, cache),
		mq:         mq,
		maxBuddies: setting.BuddyMaxCount,
		maxBlocked: setting.BlockedMaxCount,
		maxInviter: setting.InviterMaxCount,
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterBuddyServiceServer(server.GrpcServer(), s)
	return nil
}

var Module = fx.Provide(
	func(
		l *zap.Logger,
		dProvider ofx.DocumentStoreParams,
		mqParams mfx.MessageQueueParams,
		setting bfx.BuddySettingsParams,
		rcParams ofx.RedisCacheParams,
	) (out sfx.GrpcServiceResult, err error) {
		if coll, e := dProvider.DriverProvider.OpenDbDriver(setting.Name); e != nil {
			err = e
		} else if s, e := NewService(l, coll, rcParams.RedisCache, mqParams.MessageQueue, setting); e != nil {
			err = e
		} else {
			out.GrpcService = s
		}
		return
	},
)
