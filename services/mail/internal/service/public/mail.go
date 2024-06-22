package public

import (
	"context"
	"encoding/json"
	errors1 "errors"
	"sort"

	mcommon "github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/utility"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"

	pb "github.com/gstones/platform/api/gen/mail"
	"github.com/gstones/platform/services/mail/internal/service/common"
	"github.com/gstones/platform/services/mail/internal/service/errors"
)

func (s *Service) Watch(req *pb.WatchMailRequest, server pb.MailService_WatchServer) error {
	profileId, ok := utility.FromContext(server.Context(), utility.UIDContextKey)
	if !ok {
		return errors.ErrNoMetaData
	}

	// subscribe self private mail
	topic := common.MakePrivateTopic(profileId)
	if _, err := s.mq.Subscribe(server.Context(), topic,
		s.handlePrivateMails(profileId, server),
	); err != nil {
		s.logger.Error("subscribe err", zap.Error(err))
		return err
	}

	// subscribe public mail
	topic = common.MakePublicTopic("")
	if _, err := s.mq.Subscribe(server.Context(), topic,
		s.handlePublicMails(profileId, "", req.Language, req.RegisterTime, server),
	); err != nil {
		s.logger.Error("subscribe err", zap.Error(err))
		return err
	}
	// subscribe channel mails
	if req.Channel != "" {
		topic = common.MakePublicTopic(req.Channel)
		if _, err := s.mq.Subscribe(server.Context(), topic,
			s.handlePublicMails(profileId, req.Channel, req.Language, req.RegisterTime, server),
		); err != nil {
			s.logger.Error("subscribe err", zap.Error(err))
			return err
		}
	}

	if err := s.db.MergePublicAndPrivateMail(profileId, req.Channel, req.Language, req.RegisterTime); err != nil {
		s.logger.Error("merge public and private mail err", zap.Error(err))
		return err
	}
	mails, _ := s.checkAndDeleteMailLimit(profileId)
	if err := server.Send(&pb.WatchMailResponse{Mails: mails}); err != nil {
		s.logger.Error("send init mails err", zap.Error(err))
		return err
	}
	<-server.Context().Done()
	return nil
}

// return left mails and deleted mails
func (s *Service) checkAndDeleteMailLimit(profileId string) (map[int64]*pb.Mail, []int64) {
	mails := s.db.GetSelfMails(profileId)
	if len(mails) <= s.maxNum {
		return mails, nil
	}

	//1. sort mails by date
	list := lo.MapToSlice(mails, func(key int64, value *pb.Mail) *pb.Mail {
		return value
	})
	sort.Slice(list, func(i, j int) bool {
		return list[i].Date < list[j].Date
	})

	needDeleteNum := len(list) - s.maxNum
	removedNum := 0
	// 2. delete already rewarded mails and mails without rewards
	for i := 0; i < needDeleteNum; i++ {
		if list[i].Status == pb.MailStatus_REWARDED { // delete already rewarded mails
			list[i].Status = pb.MailStatus_DELETED
			removedNum++
		} else if list[i].Rewards == nil || len(list[i].Rewards) <= 0 { // delete mails without rewards
			list[i].Status = pb.MailStatus_DELETED
			removedNum++
		}
	}
	needDeleteNum -= removedNum

	//3. delete mails by date sort
	for i := 0; i < needDeleteNum; i++ {
		if list[i].Status != pb.MailStatus_DELETED {
			list[i].Status = pb.MailStatus_DELETED
		}
	}

	deleted := make([]int64, 0)
	for _, v := range list {
		if v.Status == pb.MailStatus_DELETED {
			deleted = append(deleted, v.Id)
			delete(mails, v.Id)
		}
	}

	err := s.db.DelMails(profileId, deleted...)
	if err != nil {
		s.logger.Error("add mails err", zap.Error(err))
		return mails, nil
	}
	return mails, deleted
}

func (s *Service) handlePrivateMails(
	profileId string,
	server pb.MailService_WatchServer,
) miface.SubResponseHandler {
	return func(msg miface.Message, err error) mcommon.ConsumptionCode {
		changes := make(map[int64]*pb.Mail)
		if err := json.Unmarshal(msg.Data(), &changes); err != nil {
			s.logger.Error("unmarshal err", zap.Error(err))
			return mcommon.ConsumeNackPersistentFailure
		}

		if _, deleted := s.checkAndDeleteMailLimit(profileId); deleted != nil {
			for _, v := range deleted {
				changes[v] = &pb.Mail{
					Id:     v,
					Status: pb.MailStatus_DELETED,
				}
			}
		}
		err = server.Send(&pb.WatchMailResponse{Mails: changes})
		if err != nil {
			s.logger.Error("send change err", zap.Error(err))
			return mcommon.ConsumeNackPersistentFailure
		}
		return mcommon.ConsumeAck
	}
}

func (s *Service) handlePublicMails(
	profileId, channel, language string, registerTime int64,
	server pb.MailService_WatchServer,
) miface.SubResponseHandler {
	return func(msg miface.Message, err error) mcommon.ConsumptionCode {
		changes := make(map[int64]*pb.Mail)
		if err := json.Unmarshal(msg.Data(), &changes); err != nil {
			s.logger.Error("unmarshal err", zap.Error(err))
			return mcommon.ConsumeNackPersistentFailure
		}
		if err := s.db.MergePublicAndPrivateMail(profileId, channel, language, registerTime); err != nil {
			s.logger.Error("merge public and private mail err", zap.Error(err))
			return mcommon.ConsumeNackPersistentFailure
		}

		changes, err = common.FilterMailsMapWithLanguage(changes, language)
		if err != nil {
			s.logger.Error("filter mails with language err", zap.Error(err))
			return mcommon.ConsumeAck
		}
		changes = common.FilterMailsMapWithRegisterTime(changes, registerTime)
		if _, deleted := s.checkAndDeleteMailLimit(profileId); deleted != nil {
			for _, v := range deleted {
				changes[v] = &pb.Mail{
					Id:     v,
					Status: pb.MailStatus_DELETED,
				}
			}
		}

		err = server.Send(&pb.WatchMailResponse{Mails: changes})
		if err != nil {
			s.logger.Error("send change err", zap.Error(err))
			return mcommon.ConsumeNackPersistentFailure
		}
		return mcommon.ConsumeAck
	}
}

func (s *Service) UpdateMail(ctx context.Context, request *pb.UpdateMailRequest) (*pb.UpdateMailResponse, error) {
	profileId, ok := utility.FromContext(ctx, utility.UIDContextKey)
	if !ok {
		return nil, errors.ErrNoMetaData
	}

	changes := make([]*pb.Mail, 0)
	for k, v := range request.GetUpdates() {
		if k == 0 { // update all mails
			if mails, err := s.db.UpdateAllStatus(profileId, v); err != nil {
				s.logger.Error("update all status err", zap.Error(err))
				continue
			} else {
				changes = append(changes, mails...)
			}
		} else if mail, err := s.db.UpdateOneStatus(profileId, k, v); err != nil {
			if errors1.Is(err, redis.Nil) {
				s.logger.Error(
					"mail not found",
					zap.Error(err),
					zap.Int64("id", k),
					zap.Int32("status", int32(v)),
				)
			} else {
				s.logger.Error(
					"update one status err",
					zap.Error(err),
					zap.Int64("id", k),
					zap.Int32("status", int32(v)),
				)
			}
		} else {
			changes = append(changes, mail)
		}
	}
	notice, rewards := s.makeMsg(changes)
	topic := common.MakePrivateTopic(profileId)
	if err := s.mq.Publish(
		topic,
		miface.WithJSON(notice),
	); err != nil {
		return nil, err
	}
	return &pb.UpdateMailResponse{Rewards: rewards}, nil
}

func (s *Service) makeMsg(changes []*pb.Mail) (map[int64]*pb.Mail, map[int64]*pb.MailReward) {
	notice := make(map[int64]*pb.Mail)
	rewards := make(map[int64]*pb.MailReward)
	for _, v := range changes {
		notice[v.Id] = &pb.Mail{
			Id:     v.Id,
			Status: v.Status,
		}
		if v.Status == pb.MailStatus_REWARDED {
			for _, v1 := range v.Rewards {
				if _, ok := rewards[v1.Id]; !ok {
					rewards[v1.Id] = &pb.MailReward{
						Id:     v1.Id,
						Num:    v1.Num,
						Expire: v1.Expire,
						Type:   v1.Type,
					}
				} else {
					rewards[v1.Id].Num += v1.Num
				}
			}
		}
	}
	return notice, rewards
}
