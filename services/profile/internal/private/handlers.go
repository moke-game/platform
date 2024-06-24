package private

import (
	"context"
	errors2 "errors"

	"github.com/gstones/moke-kit/orm/nerrors"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/profile/errors"
	"github.com/moke-game/platform/services/profile/internal/db/redis"
)

func (s *Service) GetProfilePrivate(ctx context.Context, request *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	if profile, err := s.profileDb.LoadProfile(request.Uid); err != nil {
		if errors2.Is(err, nerrors.ErrNotFound) {
			return nil, errors.ErrNotFound
		}
		s.logger.Warn("load profile err", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else {
		resp := &pb.GetProfileResponse{
			Profile: profile.ToProto(),
		}
		if status := redis.GetProfileStatus(s.redisCli, request.Uid); len(status) > 0 {
			resp.Profile.OnlineStatus = int32(status[request.Uid])
		}
		return resp, nil
	}
}
func (s *Service) SetProfileStatus(_ context.Context, request *pb.SetProfileStatusRequest) (*pb.SetProfileStatusResponse, error) {
	if err := redis.SetProfileStatus(s.redisCli, request.Uid, int(request.Status)); err != nil {
		s.logger.Error("set profile status err", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.SetProfileStatusResponse{}, nil
}

func (s *Service) GetProfileBasics(_ context.Context, request *pb.GetProfileBasicsRequest) (*pb.GetProfileBasicsResponse, error) {
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
	return &pb.GetProfileBasicsResponse{
		Basics: basics,
	}, nil
}
