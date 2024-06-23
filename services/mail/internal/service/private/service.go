package private

import (
	"time"

	mfx2 "github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/mail"
	"github.com/moke-game/platform/services/mail/internal/service/db"
	"github.com/moke-game/platform/services/mail/pkg/mailfx"
)

type Service struct {
	utility.WithoutAuth

	appId         string
	logger        *zap.Logger
	deployment    string
	db            *db.Database
	url           string
	mq            miface.MessageQueue
	defaultExpire time.Duration // default expire time of mail, unit is day
	aesKey        string
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterMailPrivateServiceServer(
		server.GrpcServer(),
		s,
	)
	return nil
}

func NewService(
	l *zap.Logger,
	deployment string,
	redis *redis.Client,
	url string,
	defaultExpire int32,
	aesKey string,
	mq miface.MessageQueue,
	aParams mfx2.AppParams,
) (result *Service, err error) {
	result = &Service{
		appId:         aParams.AppId,
		logger:        l,
		deployment:    deployment,
		db:            db.OpenDatabase(l, redis),
		url:           url,
		mq:            mq,
		defaultExpire: time.Hour * 24 * time.Duration(defaultExpire),
		aesKey:        aesKey,
	}
	return
}

var Module = fx.Provide(
	func(
		l *zap.Logger,
		s mfx2.AppParams,
		ms mailfx.MailSettingParams,
		mqParams mfx.MessageQueueParams,
		redisParams ofx.RedisParams,
		aParams mfx2.AppParams,
	) (gOut sfx.GrpcServiceResult, err error) {
		if svc, e := NewService(
			l,
			s.Deployment,
			redisParams.Redis,
			ms.MailUrl,
			ms.MailDefaultExpire,
			ms.MailEncryptionKey,
			mqParams.MessageQueue,
			aParams,
		); e != nil {
			err = e
			return
		} else {
			gOut.GrpcService = svc
			return
		}
	},
)
