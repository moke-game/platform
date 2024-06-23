package service

import (
	"context"

	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	pb "github.com/moke-game/platform.git/api/gen/matchmaking"
	"github.com/moke-game/platform.git/services/matchmaking/internal/errors"
	"github.com/moke-game/platform.git/services/matchmaking/internal/manager"
	"github.com/moke-game/platform.git/services/matchmaking/internal/utils"
)

func (s *Service) PveMatch(ctx context.Context, request *pb.PveMatchRequest) (*pb.PveMatchResponse, error) {
	resp := &pb.PveMatchResponse{}
	//处理请求
	groupSize := request.GroupSize
	ticket := request.Ticket
	if len(ticket) == 0 {
		s.logger.Error("Match ticket is empty")
		return nil, errors.ErrParamInvalid
	}
	//创建匹配数据
	matchManager := manager.GetGlobalMatchManager()
	matchManager.JoinPVEMatch(groupSize, ticket, request.PlayId, request.MapId)
	return resp, nil
}

func (s *Service) MatchWithRival(ctx context.Context, request *pb.MatchWithRivalRequest) (*pb.MatchWitchRivalResponse, error) {
	resp := &pb.MatchWitchRivalResponse{}
	//处理请求
	var groupSize int32 = 1
	ticket := request.Ticket
	rival := request.RivalTicket
	//创建匹配数据
	matchManager := manager.GetGlobalMatchManager()
	matchManager.JoinMatchWithRival(groupSize, ticket, rival, request.PlayId)
	return resp, nil
}

func (s *Service) Match(_ context.Context, request *pb.MatchRequest) (*pb.MatchResponse, error) {
	resp := &pb.MatchResponse{}
	//处理请求
	groupSize := request.GroupSize
	ticket := request.Ticket
	if len(ticket) == 0 {
		s.logger.Error("Match ticket is empty")
		return nil, errors.ErrParamInvalid
	}
	//创建匹配数据
	matchManager := manager.GetGlobalMatchManager()
	matchManager.JoinMatch(groupSize, ticket, request.PlayId, request.MapId)
	return resp, nil
}
func (s *Service) MatchCancel(_ context.Context, request *pb.MatchCancelRequest) (*pb.MatchCancelResponse, error) {
	matchManager := manager.GetGlobalMatchManager()
	matchManager.CancelMatch(request.ProfileId)
	return &pb.MatchCancelResponse{}, nil
}

func (s *Service) SendNotifyMsg(msg proto.Message, uid []string) {
	bt, _ := proto.Marshal(msg)
	option := miface.WithBytes(bt)
	for _, id := range uid {
		notifyTopic := utils.MakeNotifyTopic(id)
		er := s.mq.Publish(notifyTopic, option)
		if er != nil {
			s.logger.Error("matchmaking publish err", zap.Any("error", er))
		}
	}
}

func (s *Service) MatchStatus(ctx context.Context, request *pb.MatchStatusRequest) (*pb.MatchStatusResponse, error) {
	matchManager := manager.GetGlobalMatchManager()
	tim := matchManager.CheckMatchStatus(request.ProfileId)
	return &pb.MatchStatusResponse{MatchTime: tim}, nil
}
