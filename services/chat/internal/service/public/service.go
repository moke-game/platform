package public

import (
	"time"

	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	mfx2 "github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"

	pb "github.com/gstones/platform/api/gen/chat"
	"github.com/gstones/platform/services/chat/internal/service/db"
	"github.com/gstones/platform/services/chat/internal/service/errors"
	"github.com/gstones/platform/services/chat/pkg/cfx"
)

type Service struct {
	logger       *zap.Logger
	mq           miface.MessageQueue
	chatInterval time.Duration

	appId      string
	deployment string
	db         *db.Database
}

func NewService(
	l *zap.Logger,
	mq miface.MessageQueue,
	chatInterval int,
	deployment string,
	appId string,
	db *db.Database,
) (result *Service, err error) {
	result = &Service{
		logger:       l,
		mq:           mq,
		appId:        appId,
		deployment:   deployment,
		chatInterval: time.Duration(chatInterval) * time.Second,
		db:           db,
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterChatServiceServer(server.GrpcServer(), s)
	return nil
}

var ChatService = fx.Provide(
	func(
		l *zap.Logger,
		setting cfx.ChatSettingParams,
		mqParams mfx.MessageQueueParams,
		aParams mfx2.AppParams,
		redisParams ofx.RedisParams,
	) (out sfx.GrpcServiceResult, err error) {
		if s, err := NewService(
			l,
			mqParams.MessageQueue,
			setting.ChatInterval,
			aParams.Deployment,
			aParams.AppId,
			db.OpenDatabase(l, redisParams.Redis),
		); err != nil {
			return out, err
		} else {
			out.GrpcService = s
		}
		return
	},
)

func (s *Service) Chat(server pb.ChatService_ChatServer) error {
	uid, ok := utility.FromContext(server.Context(), utility.UIDContextKey)
	if !ok {
		return errors.ErrNoMetaData
	}
	chatter := CreateChatter(
		uid,
		s.deployment,
		s.appId,
		server,
		s.logger,
		s.mq,
		s.chatInterval,
		s.db,
	)
	chatter.Init()
	go chatter.Update()
	<-server.Context().Done()
	return nil
}
