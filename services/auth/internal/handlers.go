package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/auth/api"
	"github.com/moke-game/platform/services/auth/internal/db/redis"
	"github.com/moke-game/platform/services/auth/internal/utils"
)

func (s *Service) Authenticate(_ context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	data, err := s.db.LoadOrCreateUid(request.Id)
	if err != nil {
		s.logger.Error("load or create uid failed", zap.Error(err))
		return nil, ErrGeneralFailure
	}
	id := data.GetUid()
	if isBlocked, err := redis.IsBlocked(s.redisCli, id); err != nil {
		s.logger.Error("check and unblock profile failed", zap.Error(err))
		return nil, ErrGeneralFailure
	} else if isBlocked {
		s.logger.Info("account is blocked", zap.String("uid", id))
		return &pb.AuthenticateResponse{
			Uid: id,
		}, ErrPermissionDenied
	}

	isOverride := false
	if access, err := utils.CreatJwt(id, utils.TokenTypeAccess, s.jwtSecret, nil, s.jwtExpire); err != nil {
		s.logger.Error("generate access jwt failed", zap.Error(err))
		return nil, ErrGenerateJwtFailure
	} else if isOverride, err = redis.IsAuthTokenExist(s.redisCli, id); err != nil {
		s.logger.Error("get auth token failed", zap.Error(err))
		return nil, ErrGeneralFailure
	} else if err := redis.SaveAuthToken(s.redisCli, id, access, s.jwtExpire); err != nil {
		s.logger.Error("save access token failed", zap.Error(err))
		return nil, ErrGeneralFailure
	} else if refresh, err := utils.CreatJwt(id, utils.TokenTypeRefresh, s.jwtSecret, nil, s.jwtExpire); err != nil {
		s.logger.Error("generate jwt refresh token failed", zap.Error(err))
		return nil, ErrGenerateJwtFailure
	} else {
		return &pb.AuthenticateResponse{
			AccessToken:  access,
			RefreshToken: refresh,
			Uid:          id,
			IsOverride:   isOverride,
		}, nil
	}
}

func (s *Service) ValidateToken(_ context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	if request.AccessToken == "" {
		s.logger.Error("access token is empty")
		return nil, ErrClientParamFailure
	} else if uid, data, err := utils.ParseToken(request.AccessToken, utils.TokenTypeAccess, s.jwtSecret); err != nil {
		s.logger.Error("parse jwt token failed", zap.Error(err))
		return nil, ErrParseJwtTokenFailure
	} else if isExist, err := redis.IsAuthTokenSame(s.redisCli, uid, request.AccessToken); err != nil {
		s.logger.Error("get auth token failed", zap.Error(err))
		return nil, ErrGeneralFailure
	} else if !isExist {
		return nil, ErrGeneralFailure
	} else if isBlocked, err := redis.IsBlocked(s.redisCli, uid); err != nil {
		s.logger.Error("check and unblock profile failed", zap.Error(err))
		return nil, ErrGeneralFailure
	} else if isBlocked {
		s.logger.Info("profile is blocked", zap.String("uid", uid))
		return nil, ErrPermissionDenied
	} else {
		return &pb.ValidateTokenResponse{
			Uid:        uid,
			CustomData: data,
		}, nil
	}
}

func (s *Service) RefreshToken(_ context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	if uid, _, err := utils.ParseToken(request.RefreshToken, utils.TokenTypeRefresh, s.jwtSecret); err != nil {
		return nil, errors.Wrap(ErrParseJwtTokenFailure, err.Error())
	} else if access, err := utils.CreatJwt(uid, utils.TokenTypeAccess, s.jwtSecret, nil, s.jwtExpire); err != nil {
		return nil, errors.Wrap(ErrGenerateJwtFailure, err.Error())
	} else if refresh, err := utils.CreatJwt(uid, utils.TokenTypeRefresh, s.jwtSecret, nil, s.jwtExpire); err != nil {
		return nil, errors.Wrap(ErrGenerateJwtFailure, err.Error())
	} else {
		return &pb.RefreshTokenResponse{
			AccessToken:  access,
			RefreshToken: refresh,
		}, nil
	}
}

func (s *Service) ClearToken(_ context.Context, request *pb.ClearTokenRequest) (*pb.ClearTokenResponse, error) {
	if request.Uid == "" {
		s.logger.Error("uid is empty", zap.String("uid", request.Uid))
		return nil, ErrClientParamFailure
	} else if isSame, err := redis.IsAuthTokenSame(s.redisCli, request.Uid, request.AccessToken); err != nil {
		s.logger.Error("get auth token failed", zap.Error(err))
		return nil, ErrGeneralFailure
	} else if !isSame {
		return &pb.ClearTokenResponse{}, nil
	} else if err := redis.ClearAuthToken(s.redisCli, request.Uid); err != nil {
		s.logger.Error("clear auth token failed", zap.Error(err))
		return nil, ErrGeneralFailure
	}
	return &pb.ClearTokenResponse{}, nil
}

func (s *Service) Delete(_ context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var err error
	if request.Id == "" {
		s.logger.Error("uid is empty", zap.String("uid", request.Id))
		return nil, ErrClientParamFailure
	}
	if err = s.db.Delete(request.Id); err != nil {
		if errors.Is(err, nerrors.ErrNotFound) {
			return &pb.DeleteResponse{}, nil
		}
		return &pb.DeleteResponse{}, fmt.Errorf("db delete fail: %w", err)
	}
	return &pb.DeleteResponse{}, nil
}

func (s *Service) AddBlocked(_ context.Context, request *pb.AddBlockedRequest) (*pb.AddBlockedResponse, error) {
	if request.Uid == "" {
		s.logger.Error("uid is empty", zap.String("uid", request.Uid))
		return nil, ErrClientParamFailure
	}

	if request.GetIsBlock() {
		if err := redis.BlockedProfile(
			s.redisCli,
			request.Uid,
			time.Duration(request.GetDuration())*time.Second,
		); err != nil {
			s.logger.Error("block profile failed", zap.Error(err))
			return nil, ErrGeneralFailure
		}
	} else if err := redis.UnBlockedProfile(s.redisCli, request.Uid); err != nil {
		s.logger.Error("unblock profile failed", zap.Error(err))
		return nil, ErrGeneralFailure
	}
	return &pb.AddBlockedResponse{}, nil
}

func (s *Service) PackToken(_ context.Context, request *pb.PackTokenRequest) (*pb.PackTokenResponse, error) {
	if request.Uid == "" {
		s.logger.Error("uid is empty", zap.String("uid", request.Uid))
		return nil, ErrClientParamFailure
	}
	expire := 60 * time.Minute
	if request.Duration > 0 {
		expire = time.Duration(request.Duration) * time.Minute
	}
	token, err := utils.CreatJwt(request.Uid, utils.TokenTypeAccess, s.jwtSecret, request.GetCustomData(), expire)
	if err != nil {
		s.logger.Error("generate jwt token failed", zap.Error(err))
		return nil, ErrGenerateJwtFailure
	}
	return &pb.PackTokenResponse{
		AccessToken: token,
	}, nil

}
