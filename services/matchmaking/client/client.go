package client

import (
	"context"
	"fmt"

	"github.com/abiosoft/ishell"
	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"google.golang.org/grpc/metadata"

	matchmaking "github.com/moke-game/platform/api/gen/matchmaking/api"
	"github.com/moke-game/platform/services/matchmaking/pkg/mmfx"
)

type Client struct {
	client matchmaking.MatchServiceClient
	cmd    *ishell.Cmd
}

func CreateClient(host string) (*ishell.Cmd, error) {
	if client, err := mmfx.NewClient(host, sfx.SecuritySettingsParams{}); err != nil {
		return nil, err
	} else {
		p := &Client{
			client: client,
		}
		p.initShells()
		return p.cmd, nil
	}
}

func (p *Client) initShells() {
	p.cmd = &ishell.Cmd{
		Name:    "matchmaking",
		Help:    "matchmaking interactive",
		Aliases: []string{"MM"},
	}
	p.initSubShells()
}

func (p *Client) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "match",
		Help:    "match interactive",
		Aliases: []string{"M"},
		Func:    p.match,
	})
}

func (p *Client) match(c *ishell.Context) {
	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())
	if stream, err := p.client.Match(ctx, &matchmaking.MatchRequest{}); err != nil {
		slogger.Warn(c, err)
	} else {
		for {
			if res, err := stream.Recv(); err != nil {
				slogger.Warn(c, err)
				break
			} else {
				slogger.Infof(c, "game id: %s", res.GameId)
			}
		}
	}
}
