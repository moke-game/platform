package room

import (
	"runtime/debug"
	"time"

	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/zpack"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	room "github.com/moke-game/platform/api/gen/room/api"
	"github.com/moke-game/platform/services/room/internal/common"
	"github.com/moke-game/platform/services/room/internal/room/player"
	"github.com/moke-game/platform/services/room/internal/room/riface"
	"github.com/moke-game/platform/services/room/pkg/rfx"
)

const (
	FrameRate      = 30               // 帧率
	TickerInterval = 1000 / FrameRate //基于上面定义的帧率推算每帧的间隔
	ReadBuffCap    = 1024
)

type RecMsg struct {
	uid string
	msg ziface.IRequest
}

type Room struct {
	logger     *zap.Logger
	roomId     string
	mapId      int32
	playId     int32
	players    *player.Players
	frames     *Frames
	handlers   map[uint32]riface.IHandler
	msgSender  *MsgSender
	msgCh      chan *RecMsg
	duration   time.Duration
	randomSeed int64
}

func (r *Room) Receive(uid string, message ziface.IRequest) {
	r.msgCh <- &RecMsg{uid: uid, msg: message}
}

func NewRoom(
	roomId string,
	logger *zap.Logger,
	setting rfx.RoomSettingParams,
) (*Room, error) {
	return &Room{
		roomId:  roomId,
		logger:  logger,
		players: player.NewPlayers(setting.RoomPlayerMax),
		frames:  NewFrames(),
	}, nil
}

func (r *Room) Init(playId int32) error {
	r.playId = playId
	r.msgCh = make(chan *RecMsg, ReadBuffCap)
	if hub, err := CreateMsgHub(r.logger); err != nil {
		r.logger.Error("create msg hub error", zap.Error(err))
		return err
	} else {
		r.msgSender = hub
	}
	r.randomSeed = time.Now().Unix()
	r.registerHandlers()
	return nil
}

func (r *Room) registerHandlers() {
	r.handlers = make(map[uint32]riface.IHandler)
	r.handlers[uint32(room.MsgID_MSG_ID_ROOM_JOIN)] = r.joinRoom
	r.handlers[uint32(room.MsgID_MSG_ID_ROOM_EXIT)] = r.exitRoom
	r.handlers[uint32(room.MsgID_MSG_ID_ROOM_SYNC)] = r.sync
}

func (r *Room) Exit(uid string) {
	r.exitRoom(uid, nil)
}

func (r *Room) recover() {
	if !common.DeploymentGlobal.IsProd() {
		return
	}
	if e := recover(); e != nil {
		r.logger.Error("panic",
			zap.Any("recover", e),
			zap.String("stack", string(debug.Stack())),
		)
	}
}

func (r *Room) Handle(uid string, request ziface.IRequest) {
	defer r.recover()
	h, ok := r.handlers[request.GetMsgID()]
	if !ok {
		r.logger.Error(
			"room handle can not found msg id",
			zap.String("roomId", r.roomId),
			zap.Uint32("msgId", request.GetMsgID()),
		)
		return
	}

	requestId := request.GetMsgID()
	if resp, errCode := h(uid, request); resp == nil {
		return
	} else if msgName, ok := room.MsgID_name[int32(requestId)]; !ok {
		r.logger.Error("room handle can not found response msg id", zap.Uint32("msgId", requestId))
	} else if err := common.Response(request.GetConnection(), room.MsgID(requestId), errCode, resp); err != nil {
		if err.Error() == "Connection closed when send buff msg" {
			r.logger.Info("client connection already closed", zap.String("uid", uid))
		} else {
			r.logger.Error("room send response failed", zap.Error(err))
		}
	} else {
		r.logger.Info(
			"room send response",
			zap.String("roomId", r.roomId),
			zap.String("uid", uid),
			zap.String("msgName", msgName),
			zap.Any("data", resp),
			zap.Int32("errorCode", int32(errCode)),
		)
	}
}

func (r *Room) Run() error {
	ticker := time.NewTicker(TickerInterval)
	defer ticker.Stop()
	defer r.Destroy()
	lastTime := time.Now()
	for {
		now := <-ticker.C
		duration := now.Sub(lastTime)
		lastTime = now
		if r.Update(duration) {
			r.GameEnd(duration)
		}
	}
}
func (r *Room) GameEnd(duration time.Duration) {
	r.duration = 0
}

func (r *Room) Destroy() {

}

func (r *Room) RoomId() string {
	return r.roomId
}

func (r *Room) broadcastInclude(id room.NoticeID, msg proto.Message, uids ...string) {
	r.logger.Debug("room broadcast include", zap.Any("noticeId", id), zap.String("roomId", r.roomId), zap.Any("msg", msg))
	if d, err := proto.Marshal(msg); err != nil {
		r.logger.Error("room broadcast include marshal failed", zap.Error(err))
	} else {
		pack := zpack.NewMsgPackage(uint32(id), d)
		r.msgSender.BroadcastInclude(pack, uids...)
	}
}

func (r *Room) broadcastExclude(id room.NoticeID, msg proto.Message, excludes ...string) {
	r.logger.Debug("room broadcast", zap.Any("noticeId", id), zap.String("roomId", r.roomId), zap.Any("msg", msg))
	if d, err := proto.Marshal(msg); err != nil {
		r.logger.Error("room broadcast marshal failed", zap.Error(err))
	} else {
		pack := zpack.NewMsgPackage(uint32(id), d)
		r.msgSender.BroadcastExclude(pack, excludes...)
	}
}

func (r *Room) Update(dt time.Duration) bool {
	r.frames.tick()
	r.consumeMsg()

	if frame := r.frames.getCurrentFrame(); frame != nil {
		r.broadcastExclude(room.NoticeID_NOTICE_ID_ROOM_SYNC, &room.NtfFrame{
			Frames: []*room.FrameData{
				frame,
			},
		})
	}
	r.msgSender.HandleSendBuff()
	return false
}

func (r *Room) consumeMsg() {
	if len(r.msgCh) <= 0 {
		return
	}
	msgRec, _, _, _ := lo.Buffer(r.msgCh, len(r.msgCh))
	for _, v := range msgRec {
		r.Handle(v.uid, v.msg)
	}
}
