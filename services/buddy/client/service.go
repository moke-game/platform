package client

import (
	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
)

func RunBuddyCmd(url string) {
	sh := ishell.New()
	slogger.Info(sh, "interactive buddy connect to "+url)

	if buddyCmd, err := NewBuddyClient(url); err != nil {
		slogger.Die(sh, err)
	} else {
		sh.AddCmd(buddyCmd.GetCmd())
		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				err := buddyCmd.Close()
				if err != nil {
					slogger.Die(c, err)
				}
				slogger.Done(c, "interrupted, press again to exit")
			}
		})
	}

	sh.Run()
}
