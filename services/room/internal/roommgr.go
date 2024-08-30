package internal

import (
	"fmt"
	"sync"

	"github.com/gstones/moke-kit/3rd/agones/aiface"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/room/internal/room/riface"
)

type RoomMgr struct {
	logger *zap.Logger
	rooms  *sync.Map
	count  *atomic.Int32

	// agones 相关
	canShutdown *atomic.Bool // 当前服务是否可以关闭(agones)，当没有房间时，关闭当前进程
	agones      aiface.IAgones
}

func NewRoomMgr(logger *zap.Logger, agonesSdk aiface.IAgones) *RoomMgr {
	return &RoomMgr{
		logger:      logger,
		count:       atomic.NewInt32(0),
		rooms:       &sync.Map{},
		canShutdown: atomic.NewBool(false),
		agones:      agonesSdk,
	}
}

func (rm *RoomMgr) SetCanShutdown() {
	rm.canShutdown.Store(true)
}

func (rm *RoomMgr) LoadOrCreateRoom(roomId string, creator riface.RoomCreator) (riface.IRoom, error) {
	if r, ok := rm.rooms.Load(roomId); !ok {
		cr, err := creator()
		if err != nil {
			return nil, err
		}
		r, ok = rm.rooms.LoadOrStore(roomId, cr)
		room := r.(riface.IRoom)
		if !ok {
			rm.runRoom(room)
		} else {
			rm.logger.Warn("room already exists", zap.String("roomId", roomId))
		}
		return room, nil
	} else {
		return r.(riface.IRoom), nil
	}
}

func (rm *RoomMgr) LoadRoom(roomId string) (riface.IRoom, error) {
	if r, ok := rm.rooms.Load(roomId); !ok {
		return nil, fmt.Errorf("room %s not found", roomId)
	} else {
		return r.(riface.IRoom), nil
	}
}

func (rm *RoomMgr) runRoom(room riface.IRoom) {
	go func() {
		defer func() {
			rm.rooms.Delete(room.RoomId())
			rm.count.Dec()
			if rm.canShutdown.Load() && rm.count.Load() == 0 {
				if err := rm.agones.Shutdown(); err != nil {
					rm.logger.Error("agonesSDK shutdown failed", zap.Error(err))
				}
			}
			rm.logger.Info(
				"room destroy",
				zap.String("roomId", room.RoomId()),
				zap.String("roomCount", rm.count.String()),
			)
		}()
		rm.count.Inc()
		rm.logger.Info(
			"room run",
			zap.String("roomId", room.RoomId()),
			zap.String("roomCount", rm.count.String()),
		)
		if err := room.Run(); err != nil {
			rm.logger.Error("room run error", zap.Error(err))
		}
	}()
}
