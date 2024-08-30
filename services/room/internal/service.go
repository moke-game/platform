package internal

import (
	"fmt"
	"runtime/debug"

	"github.com/gstones/moke-kit/3rd/agones/aiface"
	"github.com/gstones/moke-kit/3rd/agones/pkg/agonesfx"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/znet"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	roompb "github.com/moke-game/platform/api/gen/room/api"
	"github.com/moke-game/platform/services/room/internal/common"
	"github.com/moke-game/platform/services/room/internal/room"
	"github.com/moke-game/platform/services/room/internal/room/riface"
	"github.com/moke-game/platform/services/room/pkg/rfx"
)

type Service struct {
	znet.BaseRouter
	logger       *zap.Logger
	roomMgr      *RoomMgr
	agones       *Agones
	asyncHandler map[uint32]riface.IHandler

	setting rfx.RoomSettingParams
	mq      miface.MessageQueue
}

func (s *Service) RegisterWithServer(server siface.IZinxServer) {
	for k := range roompb.MsgID_name {
		if roompb.MsgID(k) == roompb.MsgID_MSG_ID_HEARTBEAT {
			continue
		}
		server.ZinxServer().AddRouter(uint32(k), s)
	}

	if err := s.agones.ready(); err != nil {
		s.logger.Error("agones ready failed", zap.Error(err))
	}
	if err := s.agones.watchGameServer(); err != nil {
		s.logger.Error("watch game server failed", zap.Error(err))
	}
}

func (s *Service) Handle(request ziface.IRequest) {
	defer s.recover()
	s.logRequest(request)
	if request.GetMsgID() == uint32(roompb.MsgID_MSG_ID_ROOM_JOIN) {
		if err := s.initJoin(request); err != nil {
			s.logger.Error("initJoin room failed", zap.Error(err))
			if err := common.Response(
				request.GetConnection(),
				roompb.MsgID_MSG_ID_ROOM_JOIN,
				roompb.RoomErrorCode_ROOM_ERROR_CODE_INVALID,
				nil,
			); err != nil {
				s.logger.Error("send response failed", zap.Error(err))
			}
		}
	}
	if uid, err := request.GetConnection().GetProperty(common.UID); err != nil {
		s.logger.Error("get uid failed", zap.Error(err))
	} else if prop, err := request.GetConnection().GetProperty(common.Room); err != nil {
		s.logger.Error("get roomId failed", zap.Error(err))
	} else if r, ok := prop.(riface.IRoom); !ok {
		s.logger.Error("room type error")
	} else {
		r.Receive(uid.(string), request)
	}
}

func (s *Service) initJoin(request ziface.IRequest) error {
	if _, err := request.GetConnection().GetProperty(common.UID); err == nil {
		return nil
	}

	req := &roompb.ReqJoin{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return err
	}
	uid, roomID, err := s.checkToken(req.Token)
	if err != nil {
		return err
	}

	if !s.agones.CheckAndDeleteReserve(uid) {
		return fmt.Errorf("player %s not reserved or reserved timeout", uid)
	}

	if r, err := s.roomMgr.LoadOrCreateRoom(
		roomID,
		func() (riface.IRoom, error) {
			return room.CreateRoom(
				roomID,
				s.logger,
				s.setting,
			)
		},
	); err != nil {
		return err
	} else {
		request.GetConnection().SetProperty(common.UID, uid)
		request.GetConnection().SetProperty(common.Room, r)
		go func(uid string) {
			<-request.GetConnection().Context().Done()
			if err := s.agones.DeletePlayer(uid); err != nil {
				s.logger.Warn("Agones delete player failed", zap.String("roomID", r.RoomId()), zap.Error(err))
			}
			r.Exit(uid)
		}(uid)
	}
	return nil
}
func NewService(
	l *zap.Logger,
	mq miface.MessageQueue,
	agones aiface.IAgones,
	setting rfx.RoomSettingParams,
) (result *Service, err error) {
	rm := NewRoomMgr(l, agones)
	ag, err := NewAgones(agones, l, rm)
	if err != nil {
		return nil, err
	}

	result = &Service{
		logger:  l,
		roomMgr: rm,
		agones:  ag,
		setting: setting,
		mq:      mq,
	}
	return
}

var Module = fx.Provide(
	func(
		l *zap.Logger,
		agParams agonesfx.SDKParams,
		setting rfx.RoomSettingParams,
		mqParams mfx.MessageQueueParams,
	) (out sfx.ZinxServiceResult, err error) {
		if svc, e := NewService(
			l,
			mqParams.MessageQueue,
			agParams.SDK,
			setting,
		); e != nil {
			err = e
		} else {
			out.ZinxService = svc
		}
		return
	},
)

func (s *Service) logRequest(request ziface.IRequest) {
	if name, ok := roompb.MsgID_name[int32(request.GetMsgID())]; ok {
		s.logger.Info(
			"room request",
			zap.String("msg", name),
			zap.Any("data", request.GetData()),
		)
	} else {
		s.logger.Info(
			"room request",
			zap.Uint32("msgID", request.GetMsgID()),
			zap.Any("data", request.GetData()),
		)
	}
}
func (s *Service) recover() {
	if !common.DeploymentGlobal.IsProd() {
		return
	}
	if r := recover(); r != nil {
		s.logger.Error("panic",
			zap.Any("recover", r),
			zap.String("stack", string(debug.Stack())),
		)
	}
}

func (s *Service) checkToken(token string) (string, string, error) {
	return token, "10000", nil
}
