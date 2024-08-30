package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/mail/api"
	"github.com/moke-game/platform/services/mail/internal/service/common"
)

type Database struct {
	*redis.Client
	logger *zap.Logger
}

func OpenDatabase(l *zap.Logger, client *redis.Client) *Database {
	return &Database{
		client,
		l,
	}
}

func (db *Database) AddMultiplyMails(targets []string, mails ...*pb.Mail) error {
	for _, v := range targets {
		if err := db.AddMails(v, mails...); err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) AddMails(target string, mails ...*pb.Mail) error {
	if len(mails) == 0 {
		return nil
	}
	key, err := makeMailKey(target)
	if err != nil {
		return err
	}
	uMails := make(map[string]interface{})
	for _, v := range mails {
		fieldKey, err := makeFieldMailKey(v.Id)
		if err != nil {
			return err
		}
		if msg, err := json.Marshal(v); err != nil {
			return err
		} else {
			uMails[fieldKey.String()] = msg
		}
	}
	if res := db.HMSet(context.Background(), key.String(), uMails); res.Err() != nil {
		return res.Err()
	}
	return nil
}
func (db *Database) getPublicIndex(id string, channel string) (int32, error) {
	key, err := makeMailPublicIndexKey(id, channel)
	if err != nil {
		return 0, err
	}
	if res := db.Get(context.Background(), key.String()); res.Err() != nil {
		if errors.Is(res.Err(), redis.Nil) {
			return 0, nil
		}
		return 0, err
	} else if v, err := strconv.ParseInt(res.Val(), 10, 32); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

func (db *Database) savePublicIndex(id, channel string, index int32) error {
	key, err := makeMailPublicIndexKey(id, channel)
	if err != nil {
		return err
	}
	if res := db.Set(context.Background(), key.String(), index, 0); res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (db *Database) GetSelfMails(id string) map[int64]*pb.Mail {
	key, err := makeMailKey(id)
	if err != nil {
		db.logger.Error("make mail key error", zap.Error(err))
		return nil
	}
	if data := db.HGetAll(context.Background(), key.String()); data.Err() != nil {
		if errors.Is(data.Err(), redis.Nil) {
			return make(map[int64]*pb.Mail)
		}
		db.logger.Error("get mail error", zap.Error(data.Err()))
		return nil
	} else {
		res := make(map[int64]*pb.Mail)
		for _, v := range data.Val() {
			m := &pb.Mail{}
			if err := json.Unmarshal([]byte(v), m); err != nil {
				db.logger.Error("unmarshal error", zap.Error(err))
				continue
			}
			res[m.Id] = m
		}
		db.checkAndDeleteMails(id, res)
		return res
	}
}

func (db *Database) checkAndDeleteMails(profileId string, mails map[int64]*pb.Mail) {
	needDel := make([]int64, 0)
	for _, v := range mails {
		db.checkAndMarkExpired(v)
		if v.Status == pb.MailStatus_DELETED {
			needDel = append(needDel, v.Id)
			delete(mails, v.Id)
		}
	}
	if err := db.DelMails(profileId, needDel...); err != nil {
		db.logger.Error("del mail error", zap.Error(err))
	}
}

func (db *Database) checkAndMarkExpired(mail *pb.Mail) {
	if mail.Status == pb.MailStatus_DELETED {
		return
	}
	if mail.ExpireAt > 0 && mail.ExpireAt <= time.Now().Unix() {
		mail.Status = pb.MailStatus_DELETED
	}
}

func (db *Database) DelMails(profileId string, ids ...int64) error {
	if len(ids) == 0 {
		return nil
	}
	key, err := makeMailKey(profileId)
	if err != nil {
		return err
	}
	fields := make([]string, 0)
	for _, v := range ids {
		fieldKey, err := makeFieldMailKey(v)
		if err != nil {
			return err
		}
		fields = append(fields, fieldKey.String())
	}
	if res := db.HDel(context.Background(), key.String(), fields...); res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (db *Database) MergePublicAndPrivateMail(uid, channel, language string, registerTime int64) error {
	index, err := db.getPublicIndex(uid, channel)
	if err != nil {
		return err
	}
	pMails, err := db.getPublicLeftMails(channel, index)
	if err != nil {
		return err
	}
	if len(pMails) == 0 {
		return nil
	}
	pMails, err = common.FilterMailsWithLanguage(pMails, language)
	if err != nil {
		return err
	}
	pMails = common.FilterMailsWithRegisterTime(pMails, registerTime)
	if err := db.AddMails(uid, pMails...); err != nil {
		return err
	}
	rIndex := index + int32(len(pMails))
	if err := db.savePublicIndex(uid, channel, rIndex); err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateAllStatus(profileId string, status pb.MailStatus) ([]*pb.Mail, error) {
	mAll := db.GetSelfMails(profileId)
	updates := make([]*pb.Mail, 0)
	for _, v := range mAll {
		if db.checkAndUpdateStatus(v, status) {
			updates = append(updates, v)
		}
	}
	if err := db.AddMails(profileId, updates...); err != nil {
		return nil, err
	}
	return updates, nil
}

func (db *Database) UpdateOneStatus(profileId string, id int64, status pb.MailStatus) (*pb.Mail, error) {
	mail := &pb.Mail{}
	if id == 0 {
		return nil, fmt.Errorf("id is 0")
	} else if key, err := makeMailKey(profileId); err != nil {
		return nil, err
	} else if filedKey, err := makeFieldMailKey(id); err != nil {
		return nil, err
	} else if data := db.HGet(context.Background(), key.String(), filedKey.String()); data.Err() != nil {
		return nil, data.Err()
	} else if err := json.Unmarshal([]byte(data.Val()), mail); err != nil {
		return nil, err
	} else if !db.checkAndUpdateStatus(mail, status) {
		return nil, fmt.Errorf("status now:%v to:%v is not valid", mail.Status, status)
	} else if err := db.AddMails(profileId, mail); err != nil {
		return nil, err
	} else {
		return mail, nil
	}
}

func (db *Database) PushMailToPublic(channel string, mail *pb.Mail) error {
	key, err := makeMailPublicListKey(channel)
	if err != nil {
		return err
	}
	m, err := json.Marshal(mail)
	if err != nil {
		return err
	}
	if res := db.RPush(context.Background(), key.String(), m); res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (db *Database) getPublicLeftMails(channel string, start int32) ([]*pb.Mail, error) {
	key, err := makeMailPublicListKey(channel)
	if err != nil {
		return nil, err
	}
	if data := db.LRange(context.Background(), key.String(), int64(start), -1); data.Err() != nil {
		return nil, data.Err()
	} else {
		res := make([]*pb.Mail, 0)
		for _, v := range data.Val() {
			m := &pb.Mail{}
			if err := json.Unmarshal([]byte(v), m); err != nil {
				db.logger.Error("unmarshal error", zap.Error(err))
				continue
			}
			res = append(res, m)
		}
		return res, nil
	}
}

func (db *Database) checkAndUpdateStatus(mail *pb.Mail, status pb.MailStatus) bool {
	if mail.Status >= status {
		return false
	}

	// can not delete mail if it has rewards or not read
	if status == pb.MailStatus_DELETED {
		if mail.Status == pb.MailStatus_UNREAD {
			return false
		} else if mail.Status < pb.MailStatus_REWARDED && mail.Rewards != nil && len(mail.Rewards) > 0 {
			return false
		}
	}

	mail.Status = status
	return true
}
