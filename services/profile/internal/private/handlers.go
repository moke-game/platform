package private

import (
	"context"

	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/profile/errors"
	"github.com/moke-game/platform/services/profile/internal/db/redis"
)

func (s *Service) SetProfileStatus(_ context.Context, request *pb.SetProfileStatusRequest) (*pb.SetProfileStatusResponse, error) {
	if err := redis.SetProfileStatus(s.redisCli, request.Uid, int(request.Status)); err != nil {
		s.logger.Error("set profile status err", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.SetProfileStatusResponse{}, nil
}

func (s *Service) GetProfileBasicsPrivate(_ context.Context, request *pb.GetProfileBasicsPrivateRequest) (*pb.GetProfileBasicsPrivateResponse, error) {
	uids := make([]string, 0)
	if request.Uid == nil || len(request.GetUid()) <= 0 {
		s.logger.Error("get uid not found")
		return nil, errors.ErrNoMetaData
	} else {
		uids = request.Uid
	}
	basics, err := redis.GetBasicInfo(s.redisCli, uids...)
	if err != nil {
		s.logger.Error("get basic info error", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	if status := redis.GetProfileStatus(s.redisCli, uids...); len(status) > 0 {
		for _, basic := range basics {
			if v, ok := status[basic.Uid]; ok {
				basic.OnlineStatus = int32(v)
			}
		}
	}
	return &pb.GetProfileBasicsPrivateResponse{
		Basics: basics,
	}, nil
}
