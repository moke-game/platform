package service

import (
	"context"
	"time"

	"github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	pb "github.com/moke-game/platform/api/gen/buddy/api"
	"github.com/moke-game/platform/services/buddy/internal/db/model/data"
	"github.com/moke-game/platform/services/buddy/internal/errors"
	"github.com/moke-game/platform/services/buddy/internal/utils"
)

func (s *Service) AddBuddy(ctx context.Context, request *pb.AddBuddyRequest) (*pb.AddBuddyResponse, error) {
	resp := &pb.AddBuddyResponse{}
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	//好友上限
	if int32(len(selfDao.Data.Buddies)) >= s.maxBuddies {
		return nil, errors.ErrSelfBuddiesTopLimit
	}
	nowTim := time.Now().Unix()
	for _, id := range request.Uid {
		//不允许添加自己为好友
		if id == uid {
			return nil, errors.ErrCanNotAddSelf
		}
		//对方在你的黑名单中
		if _, ok := selfDao.Data.BlockedProfiles[id]; ok {
			return nil, errors.ErrInSelfBlockedList
		}
		//已经是好友
		if _, ok := selfDao.Data.Buddies[id]; ok {
			return nil, errors.ErrBuddyAlreadyAdded
		}
		//对方已经申请过添加为好友
		if _, ok := selfDao.Data.Inviters[id]; ok {
			return nil, errors.ErrBuddyAlreadyRequested
		}
		if dao, err := s.db.LoadOrCreateBuddyQueue(id); err != nil {
			return nil, errors.ErrDBErr
		} else {
			//你在对方黑名单中
			if _, ok := dao.Data.BlockedProfiles[uid]; ok {
				return nil, errors.ErrInTargetBlockedList
			}
			//已经申请过添加为好友
			if _, ok := dao.Data.Inviters[uid]; ok {
				return nil, errors.ErrBuddyAlreadyRequested
			}
			//对方好友已达上限
			if int32(len(dao.Data.Buddies)) >= s.maxBuddies {
				return nil, errors.ErrTargetBuddiesTopLimit
			}
			//对方好友申请已经达到上限
			if int32(len(dao.Data.Inviters)) >= s.maxInviter {
				return nil, errors.ErrTargetInviterTopLimit
			}
			inviter := &data.Inviter{
				UID:     uid,
				ReqTime: nowTim,
				ReqInfo: request.ReqInfo,
			}
			dao.Data.AddInviter(inviter)
			er := dao.Save()
			if er != nil {
				s.logger.Error("AddBuddy save err", zap.Error(er))
				return nil, errors.ErrDBErr
			}
			//通知对方
			//监听topic
			topic := utils.MakeBuddyTopic(id)
			inviterInfo := &pb.Inviter{
				Uid:     uid,
				ReqInfo: request.ReqInfo,
				ReqTime: nowTim,
			}
			buddyChange := &pb.BuddyChanges{}
			buddyChange.InviterAdded = append(buddyChange.InviterAdded, inviterInfo)
			respByt, _ := proto.Marshal(buddyChange)
			option := miface.WithBytes(respByt)
			if mqErr := s.mq.Publish(topic, option); mqErr != nil {
				s.logger.Error("AddBuddy mq publish err", zap.Any("error", mqErr))
			}
		}
	}
	return resp, nil
}

func (s *Service) RemoveBuddy(ctx context.Context, request *pb.RemoveBuddyRequest) (*pb.Nothing, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	if dao, err := s.db.LoadOrCreateBuddyQueue(request.Uid); err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, err
	} else {
		dao.Data.Delete(uid)
		er := dao.Save()
		if er != nil {
			return nil, er
		}
		selfDao.Data.Delete(request.Uid)
		er = selfDao.Save()
		if er != nil {
			return nil, er
		}
	}
	return &pb.Nothing{}, nil
}

func (s *Service) GetBuddies(ctx context.Context, request *pb.GetBuddyRequest) (*pb.GetBuddyResponse, error) {
	pbBuddies := &pb.Buddies{
		Buddies:      make(map[string]*pb.Buddy),
		Inviters:     make(map[string]*pb.Inviter),
		InviterSends: make(map[string]*pb.Inviter),
		Blocked:      make(map[string]*pb.Blocked),
	}
	uid := request.Uid
	if len(uid) == 0 {
		iUid, ok := ctx.Value(utility.UIDContextKey).(string)
		if !ok {
			s.logger.Error("get uid from context err")
			return nil, errors.ErrNoMetaData
		}
		uid = iUid
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	buddies := make(map[string]*pb.Buddy)
	for _, buddy := range selfDao.Data.Buddies {
		buddies[buddy.UID] = buddy.ToProto()
	}

	inviters := make(map[string]*pb.Inviter)
	for _, inviter := range selfDao.Data.Inviters {
		inviters[inviter.UID] = inviter.ToProto()
	}
	inviterSends := make(map[string]*pb.Inviter)
	for k, inviter := range selfDao.Data.InviterSends {
		inviterSends[k] = inviter.ToProto()
	}
	blocked := make(map[string]*pb.Blocked)
	for _, profile := range selfDao.Data.BlockedProfiles {
		blocked[profile.ID] = profile.ToProto()
	}
	pbBuddies.Buddies = buddies
	pbBuddies.Inviters = inviters
	pbBuddies.InviterSends = inviterSends
	pbBuddies.Blocked = blocked
	return &pb.GetBuddyResponse{Buddies: pbBuddies}, nil
}

func (s *Service) ReplyAddBuddy(ctx context.Context, request *pb.ReplyAddBuddyRequest) (*pb.ReplyAddBuddyResponse, error) {
	resp := &pb.ReplyAddBuddyResponse{}
	errs := make([]string, 0)
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	nowTim := time.Now().Unix()
	f := &data.Buddy{
		UID:     uid,
		ActTime: nowTim,
	}
	for _, id := range request.Uid {
		//不允许添加自己为好友
		if id == uid {
			errs = append(errs, errors.ErrCanNotAddSelf.Error())
			continue
		}
		//没有申请过好友
		if _, ok := selfDao.Data.Inviters[id]; !ok {
			errs = append(errs, errors.ErrInviterNotFound.Error())
			continue
		}
		//自己好友上限
		if int32(len(selfDao.Data.Buddies)) >= s.maxBuddies {
			errs = append(errs, errors.ErrSelfBuddiesTopLimit.Error())
			continue
		}
		if dao, err := s.db.LoadOrCreateBuddyQueue(id); err != nil {
			s.logger.Error("load or create buddy queue err", zap.Error(err))
			errs = append(errs, err.Error())
			continue
		} else {
			//对方好友已达上限
			if int32(len(dao.Data.Buddies)) >= s.maxBuddies {
				errs = append(errs, errors.ErrTargetBuddiesTopLimit.Error())
				continue
			}
			//移除好友申请列表
			dao.Data.RemoveInviteSend(uid)
			dao.Data.AddBuddy(f)
			er := dao.Save()
			if er != nil {
				s.logger.Error("ReplyAddBuddy save err", zap.Error(er))
				errs = append(errs, errors.ErrDBErr.Error())
			}
			//移除好友申请
			selfDao.Data.RemoveInviter(id)
			buddy := &data.Buddy{
				UID:     id,
				ActTime: nowTim,
			}
			selfDao.Data.AddBuddy(buddy)
			//errs = append(errs, ocsp.Success.String())
		}
	}
	if err := selfDao.Save(); err != nil {
		s.logger.Error("ReplyAddBuddy Save err", zap.Error(err))
		return nil, errors.ErrDBErr
	}
	resp.Failed = errs
	return resp, nil
}

func (s *Service) WatchBuddies(_ *pb.Nothing, server pb.BuddyService_WatchBuddiesServer) error {
	uid, ok := utility.FromContext(server.Context(), utility.UIDContextKey)
	if !ok {
		return errors.ErrNoMetaData
	}
	//监听topic
	topic := utils.MakeBuddyTopic(uid)
	_, err := s.mq.Subscribe(server.Context(), topic, func(msg miface.Message, err error) common.ConsumptionCode {
		if err != nil {
			s.logger.Error("subscribe buddy topic failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		resp := &pb.BuddyChanges{}
		if err := proto.Unmarshal(msg.Data(), resp); err != nil {
			s.logger.Error("unmarshal message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		if err := server.Send(resp); err != nil {
			s.logger.Error("send message failed", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		return common.ConsumeAck
	})
	if err != nil {
		s.logger.Error("subscribe buddy topic failed", zap.Error(err))
		return errors.ErrGeneralFailure
	}
	<-server.Context().Done()
	return nil
}

func (s *Service) Remark(ctx context.Context, request *pb.RemarkRequest) (*pb.Nothing, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	buddy := selfDao.Data.GetBuddy(request.Uid)
	if buddy == nil {
		return nil, errors.ErrBuddiesNotFound
	}
	buddy.Remark = request.Remark
	err = selfDao.Save()
	if err != nil {
		return nil, err
	}
	return &pb.Nothing{}, nil
}

func (s *Service) GetBlockedProfiles(ctx context.Context, _ *pb.Nothing) (*pb.ProfileIds, error) {
	resp := &pb.ProfileIds{}
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	for _, profile := range selfDao.Data.BlockedProfiles {
		pf := &pb.ProfileId{
			ProfileId: profile.ID,
		}
		resp.ProfileIds = append(resp.ProfileIds, pf)
	}
	return resp, nil
}

func (s *Service) AddBlockedProfiles(ctx context.Context, ids *pb.ProfileIds) (*pb.Nothing, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	for _, pro := range ids.ProfileIds {
		//不允许操作自己
		if pro.ProfileId == uid {
			continue
		}
		selfDao.Data.AddBlocked(pro.ProfileId)
		buddy := selfDao.Data.GetBuddy(pro.ProfileId)
		if buddy != nil {
			//解除双方好友关系
			selfDao.Data.Delete(pro.ProfileId)
			if dao, err := s.db.LoadOrCreateBuddyQueue(pro.ProfileId); err != nil {
				s.logger.Error("load or create buddy queue err", zap.Error(err))
				continue
			} else {
				//移除好友申请列表
				dao.Data.Delete(uid)
				err := dao.Save()
				if err != nil {
					return nil, err
				}
			}
		}
	}
	err = selfDao.Save()
	if err != nil {
		return nil, err
	}
	return &pb.Nothing{}, nil
}

func (s *Service) RemoveBlockedProfiles(ctx context.Context, ids *pb.ProfileIds) (*pb.Nothing, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	for _, pro := range ids.ProfileIds {
		selfDao.Data.DeleteBlocked(pro.ProfileId)
	}
	err = selfDao.Save()
	if err != nil {
		return nil, err
	}

	return &pb.Nothing{}, nil
}

func (s *Service) RefuseBuddy(ctx context.Context, request *pb.RefuseBuddyRequest) (*pb.Nothing, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	for _, rUid := range request.Uid {
		if dao, err := s.db.LoadOrCreateBuddyQueue(rUid); err != nil {
			s.logger.Error("load or create buddy queue err", zap.Error(err))
			return nil, err
		} else {
			delete(dao.Data.InviterSends, uid)
			er := dao.Save()
			if er != nil {
				return nil, er
			}
			delete(selfDao.Data.Inviters, rUid)
		}
	}
	if er := selfDao.Save(); er != nil {
		return nil, er
	}
	return &pb.Nothing{}, nil
}

func (s *Service) IsBlocked(ctx context.Context, request *pb.IsBlockedRequest) (*pb.IsBlockedResponse, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}

	resp := &pb.IsBlockedResponse{}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return resp, errors.ErrNoMetaData
	}
	if selfDao.Data.IsBlocked(request.GetUid()) {
		resp.IsBlocked = true
		return resp, nil
	}
	otherDao, err := s.db.LoadOrCreateBuddyQueue(request.GetUid())
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return resp, errors.ErrNoMetaData
	}
	if otherDao.Data.IsBlocked(uid) {
		resp.IsBlocked = true
		return resp, nil
	}
	return resp, nil
}

func (s *Service) DeleteAccount(ctx context.Context, request *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	if !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	}
	selfDao, err := s.db.LoadOrCreateBuddyQueue(uid)
	if err != nil {
		s.logger.Error("load or create buddy queue err", zap.Error(err))
		return nil, errors.ErrNoMetaData
	}
	for _, buddy := range selfDao.Data.Buddies {
		if dao, err := s.db.LoadOrCreateBuddyQueue(buddy.UID); err != nil {
			s.logger.Error("load or create buddy queue err", zap.Error(err))
			return nil, err
		} else {
			dao.Data.Delete(uid)
			er := dao.Save()
			if er != nil {
				return nil, er
			}
			selfDao.Data.Delete(request.Uid)
		}
	}
	er := selfDao.Save()
	if er != nil {
		return nil, er
	}
	return &pb.DeleteAccountResponse{}, nil
}
