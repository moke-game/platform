package room

import (
	"github.com/duke-git/lancet/v2/random"
	"github.com/gstones/zinx/ziface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	roompb "github.com/moke-game/platform/api/gen/room/api"
)

func (r *Room) joinRoom(uid string, request ziface.IRequest) (proto.Message, roompb.RoomErrorCode) {
	req := &roompb.ReqJoin{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, roompb.RoomErrorCode_ROOM_ERROR_CODE_INVALID
	}
	p := &roompb.Player{
		Uid:      uid,
		Nickname: random.RandString(8),
	}

	r.msgSender.AddSession(uid, request.GetConnection())
	if err := r.players.AddPlayer(p); err != nil {
		return nil, roompb.RoomErrorCode_ROOM_ERROR_CODE_FULL
	}
	// notice current player to add other players
	r.broadcastInclude(roompb.NoticeID_NOTICE_ID_ROOM_JOINED, &roompb.NtfRoomJoined{
		Players: r.players.GetAllPlayers(),
	}, uid)
	// notice other players to add current player
	r.broadcastExclude(roompb.NoticeID_NOTICE_ID_ROOM_JOINED, &roompb.NtfRoomJoined{
		Players: []*roompb.Player{p},
	}, uid)

	fs := r.frames.getRangeFrames(req.LastFrameIndex, r.frames.getFrameIndex())
	if len(fs) > 0 {
		r.broadcastInclude(roompb.NoticeID_NOTICE_ID_ROOM_SYNC, &roompb.NtfFrame{
			Frames: fs,
		}, uid)
	}
	return &roompb.RspJoin{
		RandomSeed: r.randomSeed,
	}, 0
}
func (r *Room) exitRoom(uid string, _ ziface.IRequest) (proto.Message, roompb.RoomErrorCode) {
	r.players.RemovePlayer(uid)
	r.broadcastExclude(roompb.NoticeID_NOTICE_ID_ROOM_EXIT, &roompb.NtfRoomExit{
		Uids: []string{uid},
	})
	r.msgSender.RemoveSession(uid)
	return &roompb.RspExit{}, 0
}

func (r *Room) sync(uid string, request ziface.IRequest) (proto.Message, roompb.RoomErrorCode) {
	req := &roompb.ReqSync{}
	if err := proto.Unmarshal(request.GetData(), req); err != nil {
		return nil, roompb.RoomErrorCode_ROOM_ERROR_CODE_INVALID
	}
	if ok := r.frames.PushCmd(req.GetCmd()); !ok {
		r.logger.Warn("sync cmd is exist", zap.String("uid", uid))
	}

	return &roompb.RspSync{}, 0
}
