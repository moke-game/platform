package private

import (
	"context"

	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/profile"
	"github.com/moke-game/platform/services/profile/errors"
	"github.com/moke-game/platform/services/profile/internal/db/redis"
)

func (s *Service) GetProfilePrivate(_ context.Context, request *pb.GetProfilePrivateRequest) (*pb.GetProfilePrivateResponse, error) {
	profiles := make([]*pb.Profile, 0)
	if request.GetUids() != nil {
		if ps, err := s.privateDao.GetProfiles(
			request.GetPlatformId(),
			request.GetChannelId(),
			request.GetUids().GetUid()...,
		); err != nil {
			s.logger.Error("privateDao GetProfiles err", zap.Error(err), zap.Any(
				"uids",
				request.GetUids().GetUid(),
			))
			return nil, errors.ErrGeneralFailure
		} else {
			profiles = append(profiles, ps...)
		}
	} else if request.GetName() != nil {
		page, pageSize := int64(1), int64(10)
		if request.GetName().GetIsRegexp() {
			if request.GetName().GetPage() > 1 {
				page = int64(request.GetName().GetPage())
			}
			if request.GetName().GetPageSize() > 0 {
				pageSize = int64(request.GetName().GetPageSize())
			}
		}
		if ps, err := s.privateDao.GetProfileByNickname(
			request.GetPlatformId(),
			request.GetChannelId(),
			request.GetName().GetName(),
			request.GetName().GetIsRegexp(),
			page,
			pageSize,
		); err != nil {
			s.logger.Error(
				"privateDao GetProfileByNickname err",
				zap.Error(err),
				zap.String("name", request.GetName().GetName()),
			)
			return nil, errors.ErrGeneralFailure
		} else {
			profiles = append(profiles, ps...)
		}
	} else if request.GetAll() != nil {
		if request.GetAll().Page < 1 || request.GetAll().PageSize < 1 {
			s.logger.Error("invalid page or page size", zap.Any("request", request))
			return nil, errors.ErrInvalidArgument
		}
		if ps, err := s.privateDao.GetAllProfiles(
			request.GetPlatformId(),
			request.GetChannelId(),
			int64(request.GetAll().Page),
			int64(request.GetAll().PageSize),
		); err != nil {
			s.logger.Error("privateDao GetProfiles err", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		} else {
			profiles = append(profiles, ps...)
		}
	} else if request.GetAccount() != "" {
		if p, err := s.privateDao.GetProfilesByAccount(
			request.GetAccount(),
		); err != nil {
			s.logger.Error("privateDao GetProfileByAccount err", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		} else {
			profiles = append(profiles, p)
		}
	} else {
		s.logger.Error("invalid request", zap.Any("request", request))
		return nil, errors.ErrInvalidArgument
	}

	uids := make([]string, 0)
	for _, v := range profiles {
		uids = append(uids, v.Uid)
	}

	if status := redis.GetProfileStatus(s.redisCli, uids...); len(status) > 0 {
		for _, v := range profiles {
			if st, ok := status[v.Uid]; ok {
				v.OnlineStatus = int32(st)
			}
		}
	}

	resp := &pb.GetProfilePrivateResponse{
		Profiles: profiles,
	}

	return resp, nil
}

func (s *Service) SetProfileStatus(_ context.Context, request *pb.SetProfileStatusRequest) (*pb.SetProfileStatusResponse, error) {
	if err := redis.SetProfileStatus(s.redisCli, request.Uid, int(request.Status)); err != nil {
		s.logger.Error("set profile status err", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.SetProfileStatusResponse{}, nil
}

func (s *Service) GetProfileBasics(_ context.Context, request *pb.GetProfileBasicsRequest) (*pb.GetProfileBasicsResponse, error) {
	var uids []string
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
