package room

import (
	"strconv"

	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/zpack"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type MsgSender struct {
	logger    *zap.Logger
	sessions  map[string]ziface.IConnection
	writeBuff map[string][]ziface.IMessage
}

func (ms *MsgSender) Init() error {
	ms.sessions = make(map[string]ziface.IConnection)
	ms.writeBuff = make(map[string][]ziface.IMessage)
	return nil
}

func (ms *MsgSender) AddSession(uid string, s ziface.IConnection) {
	ms.sessions[uid] = s
	ms.writeBuff[uid] = make([]ziface.IMessage, 0)
}

func (ms *MsgSender) RemoveSession(uid string) {
	delete(ms.sessions, uid)
	delete(ms.writeBuff, uid)
}

func (ms *MsgSender) HandleSendBuff() {
	for k, v := range ms.writeBuff {
		s, ok := ms.sessions[k]
		if !ok {
			continue
		}
		for _, v1 := range v {
			if err := s.SendBuffMsg(v1.GetMsgID(), v1.GetData()); err != nil {
				ms.logger.Error("send buff msg failed", zap.Error(err))
				continue
			}
		}
	}
	ms.writeBuff = make(map[string][]ziface.IMessage)
}

func (ms *MsgSender) CacheMsg(uid string, msg ziface.IMessage) {
	if ms.sessions[uid] == nil {
		return
	}
	ms.writeBuff[uid] = append(ms.writeBuff[uid], msg)
}

func (ms *MsgSender) SendResponse(uid string, msg ziface.IMessage) {
	if sess, ok := ms.sessions[uid]; ok {
		if err := sess.SendBuffMsg(msg.GetMsgID(), msg.GetData()); err != nil {
			ms.logger.Error("send buff msg failed", zap.Error(err))
		}
	}
}

func (ms *MsgSender) BroadcastInclude(msg ziface.IMessage, uids ...string) {
	for _, uid := range uids {
		ms.CacheMsg(uid, msg)
	}
}

func (ms *MsgSender) BroadcastExclude(msg ziface.IMessage, uids ...string) {
	for uid := range ms.sessions {
		if !ms.contains(uids, uid) {
			ms.CacheMsg(uid, msg)
		}
	}
}

func (ms *MsgSender) contains(uids []string, uid string) bool {
	for _, v := range uids {
		if v == uid {
			return true
		}
	}
	return false
}

func (ms *MsgSender) SendTo(uid int64, msgId uint32, msg interface{}) {
	if msgId == 0 {
		return
	}
	if _msg, ok := msg.(proto.Message); ok {
		if data, err := proto.Marshal(_msg); err != nil {
			ms.logger.Error("room ntf marshal failed", zap.Error(err))
		} else {
			pack := zpack.NewMsgPackage(msgId, data)
			ms.SendResponse(strconv.FormatInt(uid, 10), pack)
		}
	}
}
