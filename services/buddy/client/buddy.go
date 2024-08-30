package client

import (
	"context"
	"fmt"

	"github.com/abiosoft/ishell"
	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"google.golang.org/grpc/metadata"

	buddy "github.com/moke-game/platform/api/gen/buddy/api"
	"github.com/moke-game/platform/services/buddy/pkg/bfx"
)

// BuddyClient is the client for buddy service
type BuddyClient struct {
	client buddy.BuddyServiceClient
	cmd    *ishell.Cmd
}

func CreateBuddyClient(host string) (*ishell.Cmd, error) {
	client, err := bfx.NewBuddyClient(host, sfx.SecuritySettingsParams{})
	if err != nil {
		return nil, err
	}
	p := &BuddyClient{
		client: client,
	}
	p.initShells()
	return p.cmd, nil

}

func (p *BuddyClient) initShells() {
	p.cmd = &ishell.Cmd{
		Name:    "buddy",
		Help:    "buddy service interactive",
		Aliases: []string{"B"},
	}
	p.initSubShells()
}

func (p *BuddyClient) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "add",
		Help:    "add buddy",
		Aliases: []string{"A"},
		Func:    p.add,
	})

}

func (p *BuddyClient) add(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	msg := slogger.ReadLine(c, "buddy uid: ")
	req := &buddy.AddBuddyRequest{
		Uid:     []string{msg},
		ReqInfo: "test",
	}

	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())

	if response, err := p.client.AddBuddy(ctx, req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: %s", response)
	}

}
