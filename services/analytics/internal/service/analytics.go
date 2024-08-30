package service

import (
	"context"

	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/analytics/api"
	"github.com/moke-game/platform/services/analytics/internal/service/bi"
)

func (s *Service) Analytics(
	ctx context.Context,
	events *pb.AnalyticsEvents,
) (*pb.Nothing, error) {
	for _, v := range events.Events {
		eventType := bi.EventType(v.Event)
		p, ok := s.processes[v.DeliverTo]
		if !ok {
			s.logger.Warn("no processor found", zap.String("deliverTo", v.DeliverTo.String()))
			continue
		}
		if err := p.Handle(eventType, v.UserId, v.DistinctId, v.Properties); err != nil {
			s.logger.Error("bi data handle error", zap.Error(err))
		}
	}
	return &pb.Nothing{}, nil
}
