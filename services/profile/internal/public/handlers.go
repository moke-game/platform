package public

import (
	"context"
	"errors"
	"time"

	"github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	pb "github.com/moke-game/platform/api/gen/profile/api"
	"github.com/moke-game/platform/services/profile/changes"
	errors2 "github.com/moke-game/platform/services/profile/errors"
	"github.com/moke-game/platform/services/profile/internal/db/redis"
)

func (s *Service) IsProfileExist(ctx context.Context, request *pb.IsProfileExistRequest) (*pb.IsProfileExistResponse, error) {
	if request.Uid == "" {
		if uid, ok := ctx.Value(utility.UIDContextKey).(string); !ok {
			s.logger.Error("get uid from context err")
			return nil, errors2.ErrNoMetaData
		} else {
			request.Uid = uid
		}
	}

	if _, err := s.db.LoadProfile(request.Uid); err != nil {
		if errors.Is(err, nerrors.ErrNotFound) {
			return &pb.IsProfileExistResponse{Exist: false}, nil
		}
		s.logger.Error("load profile err", zap.Error(err))
		return nil, errors2.ErrGeneralFailure
	}

	return &pb.IsProfileExistResponse{Exist: true}, nil
}

func (s *Service) GetProfile(ctx context.Context, request *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	if request.Uid == "" {
		if uid, ok := ctx.Value(utility.UIDContextKey).(string); !ok {
			s.logger.Warn("get uid from context err")
			return nil, errors2.ErrNoMetaData
		} else {
			request.Uid = uid
		}
	}
	if profile, err := s.db.LoadProfile(request.Uid); err != nil {
		if errors.Is(err, nerrors.ErrNotFound) {
			return nil, errors2.ErrNotFound
		}
		s.logger.Warn("load profile err", zap.Error(err))
		return nil, errors2.ErrGeneralFailure
	} else {
		resp := &pb.GetProfileResponse{
			Profile: profile.ToProto(),
		}
		if status := redis.GetProfileStatus(s.redisCli, request.Uid); len(status) > 0 {
			resp.Profile.OnlineStatus = status[request.Uid]
		}
		return resp, nil
	}
}

func (s *Service) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	if uid, ok := ctx.Value(utility.UIDContextKey).(string); !ok {
		s.logger.Error("get uid from context err")
		return nil, errors2.ErrNoMetaData
	} else if name, err := s.checkName(request.GetProfile().Nickname); err != nil {
		s.logger.Error("check name err", zap.Error(err), zap.String("nickname", request.GetProfile().Nickname))
		return nil, errors2.ErrGeneralFailure
	} else {
		request.Profile.Nickname = name
		request.Profile.RegisterTime = time.Now().Unix()
		if profile, err := s.db.CreateProfile(uid, request.Profile); err != nil {
			s.logger.Error("create profile err", zap.Error(err))
			return nil, errors2.ErrGeneralFailure
		} else {
			if err := redis.UpdateBasicWithProfile(s.redisCli, uid, request.Profile); err != nil {
				s.logger.Error("set basic info err", zap.Error(err))
			}
			return &pb.CreateProfileResponse{
				Profile: profile.ToProto(),
			}, nil
		}
	}
}

func (s *Service) checkName(name string) (string, error) {
	if isExist, err := redis.IsNameExist(s.redisCli, name); err != nil {
		return "", err
	} else if isExist {
		if name, err = redis.RandomName(s.redisCli, name); err != nil {
			return "", err
		}
	}
	return name, redis.SaveName(s.redisCli, name)
}

func (s *Service) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors2.ErrNoMetaData
	}

	if request.Profile != nil {
		if updated, err := s.updateProfileInfo(uid, request.Profile); err != nil {
			s.logger.Error("update profile info err", zap.Error(err))
			return nil, err
		} else {
			return &pb.UpdateProfileResponse{Profile: updated}, nil
		}
	}
	if request.Basic != nil {
		if err := redis.SetBasicInfo(s.redisCli, uid, request.Basic); err != nil {
			s.logger.Error("set basic info err", zap.Error(err))
		}
	}

	return &pb.UpdateProfileResponse{}, nil
}

func (s *Service) updateProfileInfo(uid string, update *pb.Profile) (*pb.Profile, error) {
	profile, err := s.db.LoadProfile(uid)
	if err != nil {
		s.logger.Error("load profile err", zap.Error(err))
		return nil, errors2.ErrLoadFailure
	}

	//修改了昵称
	if newName := update.Nickname; newName != "" && newName != profile.Data.Nickname {
		if err = redis.ChangeName(s.redisCli, profile.Data.Nickname, newName); err != nil {
			s.logger.Error("UpdateProfile SaveName fail！ err:%v", zap.Error(err))
			return nil, errors2.ErrUpdateFailure
		}
	}
	if err = profile.Update(func() bool {
		return profile.UpdateData(update)
	}); err != nil {
		s.logger.Error("update profile err", zap.Error(err))
		return nil, errors2.ErrUpdateFailure
	} else if err := s.OnProfileUpdate(uid, profile.ToProto()); err != nil {
		s.logger.Error("trigger watch event fail！ err:%v", zap.Error(err))
		return nil, errors2.ErrGeneralFailure
	}
	return profile.ToProto(), nil
}

func (s *Service) GetProfileStatus(_ context.Context, request *pb.GetProfileStatusRequest) (*pb.GetProfileStatusResponse, error) {
	status := redis.GetProfileStatus(s.redisCli, request.Uid...)
	resp := &pb.GetProfileStatusResponse{
		Status: make(map[string]int32),
	}
	for k, v := range status {
		resp.Status[k] = v
	}
	return resp, nil
}

func (s *Service) WatchProfile(request *pb.WatchProfileRequest, server pb.ProfileService_WatchProfileServer) error {
	topic := changes.MakeProfileTopic(request.Uid)
	_, err := s.mq.Subscribe(server.Context(), topic, func(msg miface.Message, err error) common.ConsumptionCode {
		if err != nil {
			s.logger.Error("subscribe party topic failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		change := &pb.Profile{}
		if err := proto.Unmarshal(msg.Data(), change); err != nil {
			s.logger.Error("unmarshal message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}

		if err := server.Send(&pb.WatchProfileResponse{
			Profile: change,
		}); err != nil {
			s.logger.Error("send message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		return common.ConsumeAck
	})
	if err != nil {
		s.logger.Error("subscribe party topic failed", zap.Error(err))
		return errors2.ErrGeneralFailure
	}
	<-server.Context().Done()
	return nil
}

func (s *Service) OnProfileUpdate(uid string, change *pb.Profile) error {
	if err := redis.UpdateBasicWithProfile(s.redisCli, uid, change); err != nil {
		return err
	}
	if data, err := proto.Marshal(change); err != nil {
		return err
	} else if err := s.mq.Publish(changes.MakeProfileTopic(uid), miface.WithBytes(data)); err != nil {
		return err
	}
	return nil
}
