package client

import (
	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
)

func RunAuthCmd(url string) {
	sh := ishell.New()
	slogger.Info(sh, "interactive auth connect to "+url)

	if authCmd, err := NewAuthClient(url); err != nil {
		slogger.Die(sh, err)
		return
	} else {
		sh.AddCmd(authCmd.GetCmd())
		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				err := authCmd.Close()
				if err != nil {
					slogger.Die(c, err)
				}
				slogger.Done(c, "interrupted, press again to exit")
			}
		})
	}

	sh.Run()
}
