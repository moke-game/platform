package public

import (
	"context"
	"fmt"
	"time"

	"github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	pb "github.com/moke-game/platform/api/gen/party/api"
	"github.com/moke-game/platform/services/party/errors"
	"github.com/moke-game/platform/services/party/internal/db"
	"github.com/moke-game/platform/services/party/internal/service"
)

func (s *Service) GetParty(_ context.Context, request *pb.GetPartyRequest) (*pb.GetPartyResponse, error) {
	partyId := ""
	if request.GetUid() != "" {
		if pid, err := s.db.GetUid2Pid(request.GetUid()); err != nil {
			return &pb.GetPartyResponse{}, nil
		} else {
			partyId = pid
		}
	} else if request.GetPid() != "" {
		partyId = request.GetPid()
	}
	if party, err := s.db.GetPartySetting(partyId); err != nil {
		s.logger.Error("get party failed", zap.Error(err))
		return nil, errors.ErrPartyNotFound
	} else if party == nil {
		return &pb.GetPartyResponse{}, nil
	} else if members, err := s.db.GetPartyMembers(partyId); err != nil {
		s.logger.Error("get party members failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else {
		return &pb.GetPartyResponse{
			Party: &pb.PartyInfo{
				Party:   party,
				Members: members,
			},
		}, nil
	}
}

func (s *Service) JoinParty(request *pb.JoinPartyRequest, server pb.PartyService_JoinPartyServer) error {
	uid, ok := utility.FromContext(server.Context(), utility.UIDContextKey)
	if !ok {
		return errors.ErrNoMetaData
	}
	pid, err := s.db.GetUid2Pid(uid)
	bOnlyWatch := false
	if err == nil {
		if party, err := s.db.GetPartySetting(pid); err == nil {
			if party != nil && pid != request.Id {
				return errors.ErrHasParty
			}
		}
		if pid == request.Id {
			bOnlyWatch = true
		}
	}
	if request.GetIsCreate() {
		if pid, err := s.db.GetPartyId(); err != nil {
			return errors.ErrGeneralFailure
		} else if err := s.db.CreateParty(pid, uid, request.GetParty()); err != nil {
			s.logger.Error("create party failed", zap.Error(err))
			return errors.ErrGeneralFailure
		} else {
			request.Id = pid
			request.Party.Id = pid
		}
	}
	if !bOnlyWatch {
		if maxNumber, err := s.db.GetPartyMaxNumber(request.GetId()); err != nil {
			s.logger.Error("get party maxNumber number failed", zap.Error(err))
			return errors.ErrGeneralFailure
		} else if current, err := s.db.GetPartyMemberNum(request.GetId()); err != nil {
			s.logger.Error("check is full failed", zap.Error(err))
			return errors.ErrGeneralFailure
		} else if current >= int32(maxNumber) {
			s.logger.Error("party is full", zap.String("party", request.GetId()))
			return errors.ErrPartyFull
		} else if err := s.db.AddPartyMember(request.GetId(), uid, request.Member); err != nil {
			s.logger.Error("add party member failed", zap.Error(err))
			return errors.ErrGeneralFailure
		} else if p, err := s.db.GetPartySetting(request.GetId()); err != nil {
			s.logger.Error("get party failed", zap.Error(err))
			return errors.ErrPartyNotFound
		} else if p == nil {
			s.logger.Error("get party nil", zap.String("pid", request.Id))
			return errors.ErrPartyNotFound
		} else if members, err := s.db.GetPartyMembers(request.GetId()); err != nil {
			s.logger.Error("get party members failed", zap.Error(err))
			return errors.ErrGeneralFailure
		} else if err := server.Send(&pb.JoinPartyResponse{
			Party: &pb.PartyInfo{
				Party:   p,
				Members: members,
			},
		}); err != nil {
			s.logger.Error("send message failed", zap.Error(err))
			return errors.ErrGeneralFailure
		}
	}
	if !request.GetIsCreate() {
		if err := s.pubMemberChanges(request.GetId(), request.Member); err != nil {
			s.logger.Error("publish message failed", zap.Error(err))
		}
	}
	topic := s.makePartyTopic(request.GetId())
	_, err = s.mq.Subscribe(server.Context(), topic, func(msg miface.Message, err error) common.ConsumptionCode {
		if err != nil {
			s.logger.Error("subscribe party topic failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		changes := &pb.PartyInfo{}
		if err := proto.Unmarshal(msg.Data(), changes); err != nil {
			s.logger.Error("unmarshal message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}

		if err := server.Send(&pb.JoinPartyResponse{
			Party: changes,
		}); err != nil {
			s.logger.Error("send message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		return common.ConsumeAck
	})
	if err != nil {
		s.logger.Error("subscribe party topic failed", zap.Error(err))
		return errors.ErrGeneralFailure
	}
	<-server.Context().Done()
	if err := s.pubMemberChanges(request.GetId(), &pb.Member{
		Uid:       uid,
		IsOffline: true,
	}); err != nil {
		s.logger.Error("publish message failed", zap.Error(err))
	}
	return nil
}

func (s *Service) JoinPartyReply(_ context.Context, request *pb.JoinPartyReplyRequest) (*pb.JoinPartyReplyResponse, error) {
	if p, err := s.db.GetPartySetting(request.PartyId); err != nil {
		s.logger.Error("JoinPartyReply GetPartySetting failed", zap.Error(err))
		return nil, errors.ErrPartyNotFound
	} else if p == nil {
		s.logger.Error("JoinPartyReply GetPartySetting nil", zap.String("pid", request.PartyId))
		return nil, errors.ErrPartyNotFound
	} else {
		refuse := p.Refuse
		if refuse == nil {
			refuse = make(map[string]int64)
		}
		refuse[request.PlayerId] = time.Now().UTC().Unix()
		if err := s.db.UpdatePartySetting(p.Id, &pb.PartySetting{
			Id:     p.Id,
			Refuse: refuse,
		}); err != nil {
			s.logger.Error("JoinPartyReply UpdatePartySetting failed", zap.Error(err))
			return nil, err
		}
	}
	return &pb.JoinPartyReplyResponse{}, nil
}

func (s *Service) KickOut(ctx context.Context, request *pb.KickOutRequest) (*pb.KickOutResponse, error) {
	uid, ok := utility.FromContext(ctx, utility.UIDContextKey)
	if !ok {
		return nil, errors.ErrNoMetaData
	}
	party, err := s.db.GetPartySetting(request.PartyId)
	if err != nil || party == nil {
		s.logger.Error("get party failed", zap.Error(err))
		return nil, errors.ErrPartyNotFound
	}
	if party.Owner == request.Uid {
		s.logger.Error("can not kick out owner", zap.String("uid", uid), zap.String("owner", party.Owner))
		return nil, errors.ErrIllegal
	} else if party.Owner != uid {
		s.logger.Error("not owner", zap.String("uid", uid), zap.String("owner", party.Owner))
		return nil, errors.ErrNotOwner
	}
	members, err := s.db.GetPartyMembers(request.PartyId)
	if err != nil {
		s.logger.Error("GetPartyMembers failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	oldMembers := make([]string, 0)
	for _, member := range members {
		oldMembers = append(oldMembers, member.Uid)
	}
	if err := s.db.RemovePartyMember(request.PartyId, request.Uid); err != nil {
		s.logger.Error("remove party member failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	if err := s.pubLeave(request.PartyId, request.Uid, 2); err != nil {
		s.logger.Error("publish message failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.KickOutResponse{Ids: oldMembers}, nil
}

func (s *Service) LeaveParty(ctx context.Context, _ *pb.LeavePartyRequest) (*pb.LeavePartyResponse, error) {
	uid, ok := utility.FromContext(ctx, utility.UIDContextKey)
	if !ok {
		return nil, errors.ErrNoMetaData
	}
	pid, err := s.db.GetUid2Pid(uid)
	if err != nil {
		s.logger.Error("get uid2pid failed", zap.Error(err))
		return nil, errors.ErrPartyNotFound
	}
	owner, err := s.db.GetPartyOwner(pid)
	if err != nil {
		s.logger.Error("get party owner failed", zap.Error(err))
		return nil, errors.ErrPartyNotFound
	}
	members, err := s.db.GetPartyMembers(pid)
	if err != nil {
		s.logger.Error("GetPartyMembers failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	oldMembers := make([]string, 0)
	for _, member := range members {
		oldMembers = append(oldMembers, member.Uid)
	}
	if err := s.db.RemovePartyMember(pid, uid); err != nil {
		s.logger.Error("remove party member failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else if err := s.db.RemoveUid2Pid(uid); err != nil {
		s.logger.Error("remove uid2pid failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else if members, err := s.db.GetPartyMembers(pid); err != nil {
		s.logger.Error("get party members failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else if len(members) == 0 {
		s.db.ClearParty(pid)
		return &pb.LeavePartyResponse{Ids: oldMembers}, nil
	} else if err := s.pubLeave(pid, uid, 1); err != nil {
		s.logger.Error("publish message failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else if owner != uid {
		return &pb.LeavePartyResponse{Ids: oldMembers}, nil
	} else if err := s.updatePartyOwner(pid, members); err != nil {
		s.logger.Error("update party owner failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.LeavePartyResponse{Ids: oldMembers}, nil
}

func (s *Service) ManageParty(ctx context.Context, request *pb.ManagePartyRequest) (*pb.ManagePartyResponse, error) {
	uid, ok := utility.FromContext(ctx, utility.UIDContextKey)
	if !ok {
		return nil, errors.ErrNoMetaData
	}
	if party, err := s.db.GetPartySetting(request.Party.GetId()); err != nil {
		s.logger.Error("get party failed", zap.Error(err))
		return nil, errors.ErrPartyNotFound
	} else if party == nil {
		s.logger.Error("get party nil", zap.String("pid", request.Party.GetId()))
		return nil, errors.ErrPartyNotFound
	} else if party.Owner != uid {
		s.logger.Error("not owner", zap.String("uid", uid), zap.String("owner", party.Owner))
		return nil, errors.ErrNotOwner
	} else if err := s.db.UpdatePartySetting(request.Party.GetId(), request.Party); err != nil {
		s.logger.Error("update party failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	changes := &pb.PartyInfo{
		Party: request.Party,
	}
	if err := s.pubChanges(request.Party.GetId(), changes); err != nil {
		s.logger.Error("publish message failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}

	return &pb.ManagePartyResponse{}, nil
}

func (s *Service) makePartyTopic(id string) string {
	topic := fmt.Sprintf("%s.%s.party.%s", s.appId, s.deployment, id)
	return common.NatsHeader.CreateTopic(topic)
}

func (s *Service) UpdateMember(ctx context.Context, request *pb.UpdateMemberRequest) (*pb.UpdateMemberResponse, error) {
	uid, ok := utility.FromContext(ctx, utility.UIDContextKey)
	if !ok {
		return nil, errors.ErrNoMetaData
	}
	request.Member.Uid = uid
	if err := s.db.UpdatePartyMember(request.PartyId, request.Member); err != nil {
		s.logger.Error("add party member failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	if err := s.pubMemberChanges(request.PartyId, request.Member); err != nil {
		s.logger.Error("publish message failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.UpdateMemberResponse{}, nil
}

func (s *Service) updatePartyOwner(pid string, members map[string]*pb.Member) error {
	newOwner := ""
	for k := range members {
		newOwner = k
		break
	}
	if err := s.db.UpdatePartySetting(pid, &pb.PartySetting{
		Id:    pid,
		Owner: newOwner,
	}); err != nil {
		return err
	}

	if err := s.pubChanges(pid, &pb.PartyInfo{
		Party: &pb.PartySetting{
			Id:    pid,
			Owner: newOwner,
		},
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) pubMemberChanges(partyID string, member *pb.Member) error {
	changes := &pb.PartyInfo{}
	changes.Members = make(map[string]*pb.Member)
	changes.Members[member.Uid] = member
	return s.pubChanges(partyID, changes)
}

func (s *Service) pubLeave(partyID string, uid string, reason int32) error {
	changes := &pb.PartyInfo{}
	changes.Members = make(map[string]*pb.Member)
	changes.Members[uid] = &pb.Member{
		Uid:         uid,
		IsLeave:     true,
		LeaveReason: reason,
	}
	return s.pubChanges(partyID, changes)
}

func (s *Service) pubChanges(partyID string, changes *pb.PartyInfo) error {
	if data, err := proto.Marshal(changes); err != nil {
		return err
	} else if err := s.mq.Publish(s.makePartyTopic(partyID), miface.WithBytes(data)); err != nil {
		return err
	}
	return nil
}

func (s *Service) InviteJoinParty(_ context.Context, request *pb.InviteJoinRequest) (*pb.InviteJoinResponse, error) {
	invite, _ := s.db.GetInvite(request.PlayerId)
	if invite == nil || invite.Id == "" {
		invite = &db.PartyInvite{
			Id:     request.PlayerId,
			Refuse: make(map[string]int64),
		}
	}
	nowTim := time.Now().UTC().Unix()
	// chttodo 暂时为15秒邀请过期
	if nowTim-invite.InviteTime < service.InviteEffectiveTime {
		return &pb.InviteJoinResponse{
			ReplayCode: 4,
		}, nil
	} else {
		tim, ok := invite.Refuse[request.InviterId]
		if ok && nowTim-tim < service.RefuseTime { //拒绝5分钟
			return &pb.InviteJoinResponse{
				ReplayCode: 3,
			}, nil
		}
		invite.Inviter = request.InviterId
		invite.InviteTime = nowTim
		err := s.db.SaveInvite(invite)
		if err != nil {
			s.logger.Error("save invite err", zap.Error(err))
			return nil, err
		}
	}
	return &pb.InviteJoinResponse{}, nil
}

func (s *Service) InviteJoinReplay(ctx context.Context, request *pb.InviteReplayRequest) (*pb.InviteReplayResponse, error) {
	uid, ok := utility.FromContext(ctx, utility.UIDContextKey)
	if !ok {
		return nil, errors.ErrNoMetaData
	}
	invite, err := s.db.GetInvite(uid)
	if err != nil || invite == nil {
		return nil, errors.ErrNoMetaData
	}
	if invite.Id == "" {
		invite.Id = uid
	}
	party, err := s.db.GetPartySetting(request.PartyId)
	if err != nil || party == nil {
		return nil, errors.ErrPartyNotFound
	}
	if invite.Inviter != party.Owner {
		return nil, errors.ErrPartyNotFound
	}
	invite.Inviter = ""
	invite.InviteTime = 0
	if request.ReplayCode == 3 { //一定时间内拒绝再次被邀请
		invite.Refuse[party.Owner] = time.Now().UTC().Unix()
		err := s.db.SaveInvite(invite)
		if err != nil {
			return nil, err
		}
	} else {
		err := s.db.SaveInvite(invite)
		if err != nil {
			return nil, err
		}
	}
	return &pb.InviteReplayResponse{}, nil
}
