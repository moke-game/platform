package riface

import (
	"github.com/gstones/zinx/ziface"
	"google.golang.org/protobuf/proto"

	room "github.com/moke-game/platform/api/gen/room/api"
)

type IHandler func(uid string, request ziface.IRequest) (proto.Message, room.RoomErrorCode)
