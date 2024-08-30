package room

import (
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/room/internal/room/riface"
	"github.com/moke-game/platform/services/room/pkg/rfx"
)

func CreateRoom(
	roomId string,
	logger *zap.Logger,
	setting rfx.RoomSettingParams,
) (riface.IRoom, error) {
	if room, err := NewRoom(
		roomId,
		logger,
		setting,
	); err != nil {
		return nil, err
	} else if err := room.Init(1); err != nil {
		return nil, err
	} else {
		return room, nil
	}
}

func CreateMsgHub(logger *zap.Logger) (*MsgSender, error) {
	msg := &MsgSender{
		logger: logger,
	}
	if err := msg.Init(); err != nil {
		return nil, err
	}
	return msg, nil
}
