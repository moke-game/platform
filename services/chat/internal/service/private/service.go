package private

import (
	mfx2 "github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/chat/api"
	"github.com/moke-game/platform/services/chat/internal/service/db"
	"github.com/moke-game/platform/services/chat/pkg/cfx"
)

type Service struct {
	utility.WithoutAuth
	logger     *zap.Logger
	appId      string
	deployment string
	db         *db.Database
}

func NewService(
	l *zap.Logger,
	db *db.Database,
	deployment string,
	appId string,
) (result *Service, err error) {
	result = &Service{
		logger:     l,
		appId:      appId,
		deployment: deployment,
		db:         db,
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterChatPrivateServiceServer(server.GrpcServer(), s)
	return nil
}

var ChatService = fx.Provide(
	func(
		l *zap.Logger,
		setting cfx.ChatSettingParams,
		aParams mfx2.AppParams,
		redisParams ofx.RedisParams,
	) (out sfx.GrpcServiceResult, err error) {
		if s, err := NewService(
			l,
			db.OpenDatabase(l, redisParams.Redis),
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
