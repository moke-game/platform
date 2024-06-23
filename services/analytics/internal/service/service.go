package service

import (
	"context"
	"os"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform.git/api/gen/analytics"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi/clickhouse"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi/local"
	"github.com/moke-game/platform.git/services/analytics/pkg/analyfx"
)

type Service struct {
	utility.WithoutAuth
	logger    *zap.Logger
	mq        miface.MessageQueue
	processes map[pb.DeliveryType]bi.DataProcessor
	hostName  string
	url       string
}

func (s *Service) RegisterWithGatewayServer(server siface.IGatewayServer) error {
	return pb.RegisterAnalyticsServiceHandlerFromEndpoint(
		context.Background(), server.GatewayRuntimeMux(), s.url, server.GatewayOption(),
	)
}

func NewService(l *zap.Logger, mq miface.MessageQueue, url, hostName, rootPath,
	kisAppId, kisRegion, kisKey, kisSecret, kisSname string,
	addr, dbname, uname, passwd string,
) (result *Service, err error) {
	processes := make(map[pb.DeliveryType]bi.DataProcessor)
	if process, err := local.NewDataProcessor(l, hostName, rootPath); err != nil {
		return nil, err
	} else {
		processes[pb.DeliveryType_Local] = process
	}
	if process, err := clickhouse.NewDataProcessor(l, rootPath, addr, dbname, uname, passwd); err != nil {
		return nil, err
	} else {
		processes[pb.DeliveryType_ClickHouse] = process
	}
	result = &Service{
		logger:    l,
		mq:        mq,
		hostName:  hostName,
		processes: processes,
		url:       url,
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterAnalyticsServiceServer(server.GrpcServer(), s)
	return nil
}

var ServiceModule = fx.Provide(
	func(
		l *zap.Logger,
		mq mfx.MessageQueueParams,
		settings analyfx.AnalyticsSettingParams,
	) (out sfx.GrpcServiceResult, gw sfx.GatewayServiceResult, err error) {
		hostname := os.Getenv("HOST_NAME")
		if svc, e := NewService(l, mq.MessageQueue, settings.AnalyticsUrl, hostname, settings.LocalBiPath,
			settings.IfunBiAppId, settings.IfunBiKinesisRegion, settings.IfunBiKinesisKey, settings.IfunBiKinesisSecret, settings.IfunBiKinesisStreamName,
			settings.CKaddr, settings.CKdb, settings.CKuser, settings.CKpasswd,
		); e != nil {
			err = e
		} else {
			out.GrpcService = svc
			gw.GatewayService = svc
		}
		return
	},
)
