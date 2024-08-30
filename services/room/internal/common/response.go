package common

import (
	"github.com/gstones/zinx/ziface"
	"google.golang.org/protobuf/proto"

	room "github.com/moke-game/platform/api/gen/room/api"
)

func Response(connect ziface.IConnection, msgId room.MsgID, code room.RoomErrorCode, msg proto.Message) error {
	if connect == nil {
		return nil
	} else if msg, err := proto.Marshal(msg); err != nil {
		return err
	} else if resp, err := proto.Marshal(&room.Response{
		ErrorCode: code,
		Message:   msg,
	}); err != nil {
		return err
	} else if err := connect.SendBuffMsg(uint32(msgId), resp); err != nil {
		return err
	}
	return nil
}
