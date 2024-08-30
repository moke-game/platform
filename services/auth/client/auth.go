package client

import (
	"context"

	"github.com/abiosoft/ishell"
	"github.com/duke-git/lancet/v2/random"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"

	pb "github.com/moke-game/platform/api/gen/auth/api"
	"github.com/moke-game/platform/services/auth/pkg/afx"
)

type AuthClient struct {
	client pb.AuthServiceClient
	cmd    *ishell.Cmd
}

func CreateAuthClient(host string) (*ishell.Cmd, error) {
	if client, err := afx.NewAuthClient(host, sfx.SecuritySettingsParams{}); err != nil {
		return nil, err
	} else {
		p := &AuthClient{
			client: client,
		}
		p.initShells()
		return p.cmd, nil
	}
}

func (p *AuthClient) initShells() {
	p.cmd = &ishell.Cmd{
		Name:    "auth",
		Help:    "auth interactive",
		Aliases: []string{"A"},
	}
	p.initSubShells()
}

func (p *AuthClient) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "signup",
		Help:    "signup interactive",
		Aliases: []string{"S"},
		Func:    p.signUp,
	})

	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "signin",
		Help:    "signin interactive",
		Aliases: []string{"I"},
		Func:    p.signIn,
	})

	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "verify",
		Help:    "verify interactive",
		Aliases: []string{"V"},
		Func:    p.verify,
	})

	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "token",
		Help:    "request a jwt token",
		Aliases: []string{"T"},
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

	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "pack",
		Help:    "pack a jwt token with custom data",
		Func:    p.pack,
		Aliases: []string{"P"},
	})

}

func (p *AuthClient) verify(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	slogger.Info(c, "Enter token...")
	msg := slogger.ReadLine(c, "token: ")
	client, err := supabase.NewClient("https://aslojfweajpclwrdcimz.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFzbG9qZndlYWpwY2x3cmRjaW16Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjMxMDcxMzgsImV4cCI6MjAzODY4MzEzOH0.Q3AoUiBFderVlwJZ9aR5H8RSQNaTOwjo2JD_3DFWfmI", nil)
	if err != nil {
		slogger.Warn(c, err)
		return
	}
	req := types.VerifyRequest{
		Type:  types.VerificationTypeSignup,
		Token: msg,
	}
	if response, err := client.Auth.Verify(req); err != nil {
		slogger.Warn(c, err)
		return
	} else {
		slogger.Infof(c, "Response: %v", response)
	}

}

func (p *AuthClient) signIn(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	slogger.Info(c, "Enter email...")
	email := slogger.ReadLine(c, "email: ")

	pwd := slogger.ReadLine(c, "password: ")

	client, err := supabase.NewClient("https://aslojfweajpclwrdcimz.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFzbG9qZndlYWpwY2x3cmRjaW16Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjMxMDcxMzgsImV4cCI6MjAzODY4MzEzOH0.Q3AoUiBFderVlwJZ9aR5H8RSQNaTOwjo2JD_3DFWfmI", nil)
	if err != nil {
		slogger.Warn(c, err)
		return
	}

	if resp, err := client.Auth.SignInWithEmailPassword(email, pwd); err != nil {
		slogger.Warn(c, err)
	} else {
		req := &pb.ValidateTokenRequest{
			AccessToken: resp.AccessToken,
		}
		if response, err := p.client.ValidateToken(context.TODO(), req); err != nil {
			slogger.Warn(c, err)
		} else {
			slogger.Infof(c, "Response: %s", response)
		}
	}

}

func (p *AuthClient) signUp(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	slogger.Info(c, "Enter email...")
	email := slogger.ReadLine(c, "email: ")
	client, err := supabase.NewClient("https://aslojfweajpclwrdcimz.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFzbG9qZndlYWpwY2x3cmRjaW16Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjMxMDcxMzgsImV4cCI6MjAzODY4MzEzOH0.Q3AoUiBFderVlwJZ9aR5H8RSQNaTOwjo2JD_3DFWfmI", nil)
	if err != nil {
		slogger.Warn(c, err)
		return
	}

	pwd := slogger.ReadLine(c, "password: ")
	req := types.SignupRequest{
		Email:    email,
		Password: pwd,
	}
	if resp, err := client.Auth.Signup(req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: %v", resp)
	}
}

func (p *AuthClient) auth(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	slogger.Info(c, "Enter username...")
	msg := slogger.ReadLine(c, "username: ")
	req := &pb.AuthenticateRequest{
		Id:    msg,
		AppId: "test",
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
	req := &pb.ValidateTokenRequest{
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

	req := &pb.RefreshTokenRequest{
		RefreshToken: msg,
	}

	if response, err := p.client.RefreshToken(context.TODO(), req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: refresh %s", response.RefreshToken)
		slogger.Infof(c, "Response: access %s", response.AccessToken)
	}
}

func (p *AuthClient) pack(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	custom := slogger.ReadLine(c, "custom data: ")

	req := &pb.PackTokenRequest{
		Uid:        random.RandString(8),
		CustomData: []byte(custom),
	}
	if response, err := p.client.PackToken(context.TODO(), req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "Response: %s", response)
	}
}
