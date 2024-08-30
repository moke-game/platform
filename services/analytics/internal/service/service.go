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

	pb "github.com/moke-game/platform/api/gen/analytics/api"
	"github.com/moke-game/platform/services/analytics/internal/service/bi"
	"github.com/moke-game/platform/services/analytics/internal/service/bi/clickhouse"
	"github.com/moke-game/platform/services/analytics/internal/service/bi/local"
	"github.com/moke-game/platform/services/analytics/internal/service/bi/thinkingdata"
	"github.com/moke-game/platform/services/analytics/pkg/analyfx"
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

func NewService(
	l *zap.Logger,
	mq miface.MessageQueue,
	settings analyfx.AnalyticsSettingParams,
) (result *Service, err error) {
	processors := make(map[pb.DeliveryType]bi.DataProcessor)
	hostname := os.Getenv("HOST_NAME")
	if proc, e := local.NewDataProcessor(l, hostname, settings.LocalBiPath); e != nil {
		return nil, e
	} else {
		processors[pb.DeliveryType_Local] = proc
	}
	if settings.CKAddr != "" && settings.CKDB != "" && settings.CKUser != "" && settings.CKPasswd != "" {
		if proc, e := clickhouse.NewDataProcessor(
			l, settings.LocalBiPath, settings.CKAddr,
			settings.CKDB, settings.CKUser, settings.CKPasswd,
		); e != nil {
			return nil, e
		} else {
			processors[pb.DeliveryType_ClickHouse] = proc
		}
	}
	if proc, e := thinkingdata.NewDataProcessor(
		l,
		settings.TDPath,
	); e != nil {
		return nil, e
	} else {
		processors[pb.DeliveryType_ThinkingData] = proc
	}

	result = &Service{
		logger:    l,
		mq:        mq,
		hostName:  hostname,
		processes: processors,
		url:       settings.AnalyticsUrl,
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
		if svc, e := NewService(l, mq.MessageQueue, settings); e != nil {
			err = e
		} else {
			out.GrpcService = svc
			gw.GatewayService = svc
		}
		return
	},
)
