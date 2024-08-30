package riface

import (
	"github.com/gstones/zinx/ziface"
)

type RoomCreator func() (IRoom, error)

type IMessage interface {
	GetMsgId() uint32 // Gets the ID of the message(获取消息ID)
	GetMsgData() []byte
}

type IRoom interface {
	Init(playId int32) error
	Run() error
	RoomId() string
	Receive(uid string, message ziface.IRequest)
	Exit(uid string)
}
