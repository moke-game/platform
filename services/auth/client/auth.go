package client

import (
	"context"

	"github.com/abiosoft/ishell"
	"google.golang.org/grpc"

	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/tools"

	pb2 "github.com/moke-game/platform/api/gen/auth"
)

type AuthClient struct {
	client pb2.AuthServiceClient
	cmd    *ishell.Cmd
	conn   *grpc.ClientConn
}

func NewAuthClient(host string) (*AuthClient, error) {
	if conn, err := tools.DialInsecure(host); err != nil {
		return nil, err
	} else {
		cmd := &ishell.Cmd{
			Name:    "auth",
			Help:    "auth interactive",
			Aliases: []string{"A"},
		}
		p := &AuthClient{
			client: pb2.NewAuthServiceClient(conn),
			cmd:    cmd,
			conn:   conn,
		}
		p.initSubShells()
		return p, nil
	}
}

func (p *AuthClient) GetCmd() *ishell.Cmd {
	return p.cmd
}

func (p *AuthClient) Close() error {
	return p.conn.Close()
}

func (p *AuthClient) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "auth",
		Help:    "authorize interactive",
		Aliases: []string{"A"},
		Func:    p.auth,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "validate",
		Help:    "validate interactive",
		Aliases: []string{"V"},
		Func:    p.validate,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "refresh",
		Help:    "refresh interactive",
		Aliases: []string{"R"},
		Func:    p.refresh,
	})

}

func (p *AuthClient) auth(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	slogger.Info(c, "Enter username...")
	msg := slogger.ReadLine(c, "username: ")
	req := &pb2.AuthenticateRequest{
		Id:    msg,
		AppId: "test",
		Auth:  pb2.AuthenticateRequest_CREATE_UID,
	}

	if response, err := p.client.Authenticate(context.TODO(), req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: access %s", response.AccessToken)
		slogger.Infof(c, "Response: refresh %s", response.RefreshToken)
	}
}

func (p *AuthClient) validate(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	slogger.Info(c, "Enter token...")
	msg := slogger.ReadLine(c, "token: ")
	req := &pb2.ValidateTokenRequest{
		AccessToken: msg,
	}
	if response, err := p.client.ValidateToken(context.TODO(), req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: %s", response)
	}
}

func (p *AuthClient) refresh(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	slogger.Info(c, "Enter refresh token...")
	msg := slogger.ReadLine(c, "refresh token: ")

	req := &pb2.RefreshTokenRequest{
		RefreshToken: msg,
	}

	if response, err := p.client.RefreshToken(context.TODO(), req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: refresh %s", response.RefreshToken)
		slogger.Infof(c, "Response: access %s", response.AccessToken)
	}
}
