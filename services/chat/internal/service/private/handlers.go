package private

import (
	"context"

	"go.uber.org/zap"

	pb "github.com/moke-game/platform.git/api/gen/chat"
	"github.com/moke-game/platform.git/services/chat/internal/service/errors"
)

func (s *Service) AddBlocked(_ context.Context, request *pb.AddBlockedRequest) (*pb.AddBlockedResponse, error) {
	if request.IsBlocked {
		if err := s.db.AddBlockedList(request.ProfileId, request.Duration); err != nil {
			s.logger.Error("add blocked failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
		return &pb.AddBlockedResponse{}, nil
	} else {
		if err := s.db.RemoveBlockedList(request.ProfileId); err != nil {
			s.logger.Error("remove blocked failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
		return &pb.AddBlockedResponse{}, nil
	}
}
