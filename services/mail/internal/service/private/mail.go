package private

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	pb "github.com/gstones/platform/api/gen/mail"
	"github.com/gstones/platform/services/mail/internal/service/common"
	"github.com/gstones/platform/services/mail/internal/service/errors"
)

func (s *Service) SendMail(_ context.Context, request *pb.SendMailRequest) (*pb.SendMailResponse, error) {
	s.logger.Info("send mail request", zap.Any("request", request))
	if mail, e := s.makeMailFromRequest(request); e != nil {
		s.logger.Error("make mail from request failed", zap.Error(e))
		return nil, errors.ErrParamsInvalid
	} else if request.SendType == pb.SendMailRequest_ROLE {
		if len(request.RoleIds) <= 0 {
			s.logger.Error("role id is empty")
			return nil, errors.ErrParamsInvalid
		} else if e := s.savePrivateMail(request.RoleIds, mail); e != nil {
			s.logger.Error("save private mail failed", zap.Error(e))
			return nil, errors.ErrSaveMailFailed
		}
		for _, v := range request.RoleIds {
			topic := common.MakePrivateTopic(v)
			if e := s.mq.Publish(
				topic,
				miface.WithJSON(map[int64]*pb.Mail{mail.Id: mail}),
			); e != nil {
				s.logger.Error("publish private mail failed", zap.Error(e))
				return nil, errors.ErrPublishMailFailed
			}
		}
	} else if request.SendType == pb.SendMailRequest_ALL {
		if e := s.db.PushMailToPublic(request.ChannelId, mail); e != nil {
			s.logger.Error("push mail to public failed", zap.Error(e))
			return nil, errors.ErrSaveMailFailed
		} else if e := s.mq.Publish(
			common.MakePublicTopic(request.ChannelId),
			miface.WithJSON(map[int64]*pb.Mail{mail.Id: mail}),
		); e != nil {
			s.logger.Error("publish public mail failed", zap.Error(e))
			return nil, errors.ErrPublishMailFailed
		}
	}
	return &pb.SendMailResponse{}, nil
}

func (s *Service) makeMailFromRequest(req *pb.SendMailRequest) (*pb.Mail, error) {
	if req.Mail == nil {
		return nil, fmt.Errorf("mail data is empty %v", req)
	}
	res := proto.Clone(req.Mail).(*pb.Mail)
	if res.Id == 0 {
		res.Id = time.Now().UnixMilli()
	}
	if res.Date == 0 {
		res.Date = time.Now().Unix()
	}
	if res.Filters == nil {
		res.Filters = &pb.Mail_Filter{}
	}

	if res.Filters.RegisterTime == 0 {
		res.Filters.RegisterTime = time.Now().Unix()
	} else if res.Filters.RegisterTime < 0 {
		res.Filters.RegisterTime = math.MaxInt64
	}

	if res.ExpireAt <= 0 {
		res.ExpireAt = res.Date + int64(s.defaultExpire.Seconds())
	} else if res.ExpireAt <= res.Date {
		return nil, fmt.Errorf("expire time is invalid %v", res)
	}

	return res, nil
}

func (s *Service) makeMailRewards(items []string) ([]*pb.MailReward, error) {
	rewards := make([]*pb.MailReward, 0)
	for _, v := range items {
		strs := strings.Split(v, ":")
		if len(strs) < 2 {
			return nil, fmt.Errorf("invalid item %s", v)
		}
		if id, err := strconv.ParseInt(strs[0], 10, 32); err != nil {
			return nil, fmt.Errorf("invalid item id %s", v)
		} else if num, err := strconv.ParseInt(strs[1], 10, 32); err != nil {
			return nil, fmt.Errorf("invalid item num %s", v)
		} else {
			rewards = append(rewards, &pb.MailReward{
				Id:   id,
				Num:  int32(num),
				Type: int32(id),
			})
		}
	}
	return rewards, nil
}

func (s *Service) initDefault(mail *pb.Mail) *pb.Mail {
	mail.Id = time.Now().Unix()
	if mail.Date == 0 {
		mail.Date = time.Now().Unix()
	}
	if mail.ExpireAt <= 0 {
		mail.ExpireAt = mail.Date + int64(s.defaultExpire.Seconds())
	} else {
		duration := time.Hour * time.Duration(mail.ExpireAt)
		mail.ExpireAt = mail.Date + int64(duration.Seconds())
	}
	for _, v := range mail.Rewards {
		if v.Id == 0 && v.Type != 0 {
			v.Id = int64(v.Type)
		}
		if v.Type == 0 && v.Id != 0 {
			v.Type = int32(v.Id)
		}
	}
	return mail
}

func (s *Service) savePrivateMail(targets []string, mail *pb.Mail) error {
	if len(targets) <= 0 {
		return fmt.Errorf("target is empty")
	}
	if err := s.db.AddMultiplyMails(targets, mail); err != nil {
		return err
	}
	return nil
}
