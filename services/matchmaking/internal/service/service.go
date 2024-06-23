package service

import (
	"context"
	"fmt"
	"time"

	allocation "agones.dev/agones/pkg/allocation/go"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/gstones/moke-kit/3rd/agones/pkg/agonesfx"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb2 "github.com/moke-game/platform/api/gen/auth"
	pb "github.com/moke-game/platform/api/gen/matchmaking"
	"github.com/moke-game/platform/services/auth/pkg/afx"
	"github.com/moke-game/platform/services/matchmaking/internal/agones"
	"github.com/moke-game/platform/services/matchmaking/internal/db"
	"github.com/moke-game/platform/services/matchmaking/internal/manager"
	"github.com/moke-game/platform/services/matchmaking/pkg/matchfx"
)

const (
	TickTime = time.Second // 1s1帧
)

type Service struct {
	utility.WithoutAuth
	logger    *zap.Logger
	db        *db.Database
	mq        miface.MessageQueue
	aClient   pb2.AuthServiceClient
	allocator *agones.Allocator // agones allocateCli client
}

func NewService(
	l *zap.Logger,
	rClient *redis.Client,
	messageQueue miface.MessageQueue,
	aClient pb2.AuthServiceClient,
	agonesClient allocation.AllocationServiceClient,
	accCli *globalaccelerator.Client,
	subNet string,
) (result *Service, err error) {

	result = &Service{
		logger:    l,
		db:        db.OpenDatabase(l, rClient),
		mq:        messageQueue,
		aClient:   aClient,
		allocator: agones.CreateAgonesAllocator(l, agonesClient, accCli, subNet),
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterMatchServiceServer(server.GrpcServer(), s)
	return nil
}

var MatchService = fx.Provide(
	func(
		l *zap.Logger,
		setting matchfx.MatchSettingParams,
		redisParams ofx.RedisParams,
		mqParams mfx.MessageQueueParams,
		aParams afx.AuthClientParams,
		agonesClient agonesfx.AllocateParams,
	) (out sfx.GrpcServiceResult, err error) {
		cfg, err := config.LoadDefaultConfig(context.Background(),
			config.WithRegion(setting.AWSRegion),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(setting.AWSKey, setting.AWSSecret, "")),
		)
		if err != nil {
			return out, fmt.Errorf("failed to load aws config: %w", err)
		}
		accCli := globalaccelerator.NewFromConfig(cfg)
		if s, err := NewService(
			l,
			redisParams.Redis,
			mqParams.MessageQueue,
			aParams.AuthClient,
			agonesClient.AllocateClient,
			accCli,
			setting.SubNet,
		); err != nil {
			return out, err
		} else {
			out.GrpcService = s
			s.initManager()
			go func() {
				s.run()
			}()
		}
		return
	},
)

func (s *Service) initManager() {
	manager.NewMatchManager(s.db, s.logger, s.mq, s.aClient, s.allocator)
}

func (s *Service) run() {
	ticker := time.NewTicker(TickTime)
	last := time.Now()
	matchManager := manager.GetGlobalMatchManager()
	for {
		select {
		case now := <-ticker.C:
			delta := now.Sub(last)
			if matchManager != nil {
				ms := int32(delta.Milliseconds())
				matchManager.Update(ms)
				matchManager.UpdateRetry(ms)
			}
		}
	}
}
