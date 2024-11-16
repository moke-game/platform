package internal

import (
	"sync"
	"time"

	"agones.dev/agones/pkg/sdk"
	"github.com/gstones/moke-kit/3rd/agones/aiface"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/room/internal/common"
)

// Agones: list and count

const (
	WorldReserveTimeout = 30 // 秒
)

type Agones struct {
	agones  aiface.IAgones
	logger  *zap.Logger
	roomMgr *RoomMgr

	// 大世界 30s未进入房间的玩家将被禁止进入
	playerReserves *sync.Map // 当前预定的玩家
}

func NewAgones(
	agones aiface.IAgones,
	logger *zap.Logger,
	roomMgr *RoomMgr,
) (*Agones, error) {
	a := &Agones{
		agones:  agones,
		logger:  logger,
		roomMgr: roomMgr,
	}
	if err := a.init(); err != nil {
		return nil, err
	}
	go a.health()
	return a, nil
}

func (a *Agones) init() error {
	a.playerReserves = &sync.Map{}
	if err := a.agones.Init(); err != nil {
		return err
	}
	return nil
}

func (a *Agones) ready() error {
	if err := a.agones.Ready(); err != nil {
		return err
	}
	return nil
}

func (a *Agones) health() {
	duration := time.Second * 5
	for {
		if err := a.agones.Health(); err != nil {
			a.logger.Error("health check failed", zap.Error(err))
		}
		time.Sleep(duration)
	}
}

func (a *Agones) CheckAndDeleteReserve(uid string) bool {
	if !common.DeploymentGlobal.IsProd() {
		return true
	}
	_, ok := a.playerReserves.LoadAndDelete(uid)
	return ok
}

func (a *Agones) DeletePlayer(uid string) error {
	if err := a.agones.CounterList().DeleteListValue("players", uid); err != nil {
		return err
	}
	return nil
}

func (a *Agones) watchGameServer() error {
	if err := a.agones.WatchGameServer(a.watchCallBack); err != nil {
		return err
	}
	return nil
}

func (a *Agones) watchCallBack(gs *sdk.GameServer) {
	a.updateWorldPlayers(gs)
	a.updateBattleRooms(gs)
}

func (a *Agones) updateBattleRooms(gs *sdk.GameServer) {
	a.logger.Debug("watch game server", zap.Any("counters", gs.Status.Counters))
	if rooms, ok := gs.Status.Counters["rooms"]; !ok {
		return
	} else if rooms.Count >= rooms.GetCapacity() {
		a.logger.Info("rooms is full", zap.Any("rooms", rooms))
		a.roomMgr.SetCanShutdown()
	}
}

func (a *Agones) updateWorldPlayers(gs *sdk.GameServer) {
	if _, ok := gs.Status.Lists["players"]; !ok {
		return
	}

	a.playerReserves.Range(func(key, value interface{}) bool {
		if time.Now().Unix()-value.(int64) > int64(WorldReserveTimeout) {
			if err := a.agones.CounterList().DeleteListValue("players", key.(string)); err != nil {
				a.logger.Error("delete player failed", zap.Error(err))
			}
			a.logger.Info("delete player", zap.String("uid", key.(string)))
			a.playerReserves.Delete(key)
		}
		return true
	})
	if addPlayer, ok := gs.GetObjectMeta().GetLabels()["agones.dev/sdk-player-add"]; !ok || addPlayer == "" {
		a.logger.Debug("player add failed", zap.Any("addPlayer", addPlayer))
	} else {
		a.logger.Info("player add", zap.Any("addPlayer", addPlayer))
		a.playerReserves.Store(addPlayer, time.Now().Unix())
		if err := a.agones.SetLabel("player-add", ""); err != nil {
			a.logger.Error("agones sdk set label error", zap.Error(err))
		}
	}
}
