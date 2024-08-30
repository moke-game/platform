package client

import (
	"github.com/abiosoft/ishell"
	"github.com/duke-git/lancet/v2/random"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/znet"
	"google.golang.org/protobuf/proto"

	room "github.com/moke-game/platform/api/gen/room/api"
)

type RoomWSClient struct {
	cmd *ishell.Cmd

	username string
	client   ziface.IClient
}

func (ac *RoomWSClient) initShell() {
	ac.cmd = &ishell.Cmd{
		Name:    "room",
		Help:    "room interactive shell",
		Aliases: []string{"R"},
	}
	ac.initSubCmd()
}

func (ac *RoomWSClient) initSubCmd() {
	ac.cmd.AddCmd(&ishell.Cmd{
		Name:    "join",
		Help:    "join room",
		Aliases: []string{"J"},
		Func:    ac.joinRoom,
	})

	ac.cmd.AddCmd(&ishell.Cmd{
		Name:    "sync",
		Help:    "sync message",
		Aliases: []string{"S"},
		Func:    ac.sync,
	})

	ac.cmd.AddCmd(&ishell.Cmd{
		Name:    "exit",
		Help:    "exit room",
		Aliases: []string{"E"},
		Func:    ac.exit,
	})
}

func (ac *RoomWSClient) joinRoom(c *ishell.Context) {
	roomId := slogger.ReadLine(c, "input room id:")

	req := &room.ReqJoin{
		Token:  ac.username,
		RoomId: roomId,
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return
	}
	if err := ac.client.Conn().SendMsg(uint32(room.MsgID_MSG_ID_ROOM_JOIN), data); err != nil {
		return
	}
}

func (ac *RoomWSClient) sync(c *ishell.Context) {
	req := &room.ReqSync{
		Cmd: &room.CmdData{
			Uid:  ac.username,
			Data: random.RandBytes(10),
		},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return
	}
	if err := ac.client.Conn().SendMsg(uint32(room.MsgID_MSG_ID_ROOM_SYNC), data); err != nil {
		return
	}
}

func (ac *RoomWSClient) exit(c *ishell.Context) {
	req := &room.ReqExit{}

	data, err := proto.Marshal(req)
	if err != nil {
		return
	}
	if err := ac.client.Conn().SendMsg(uint32(room.MsgID_MSG_ID_ROOM_EXIT), data); err != nil {
		return
	}
}

func CreateRoomWSClient(url string, port int) (*ishell.Cmd, error) {
	client := znet.NewWsClient(url, port)
	rc := &RoomWSClient{
		client:   client,
		username: random.RandString(8),
	}
	rc.initShell()
	client.Start()
	return rc.cmd, nil
}
