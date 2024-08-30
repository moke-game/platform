package main

import (
	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"

	analytics "github.com/moke-game/platform/services/analytics/client"
	auth "github.com/moke-game/platform/services/auth/client"
	buddy "github.com/moke-game/platform/services/buddy/client"
	"github.com/moke-game/platform/services/matchmaking/client"
)

const (
	defaultURL      = "localhost:8081"
	defaultUsername = "test"
)

var options struct {
	url      string
	username string
}

func main() {
	rootCmd := &cobra.Command{
		Short: "Run a platform service client",
	}
	rootCmd.PersistentFlags().StringVar(&options.url, "url", defaultURL, "service url")
	rootCmd.PersistentFlags().StringVar(&options.username, "username", defaultUsername, "username for authentication")
	{
		cmd := &cobra.Command{
			Use:   "shell",
			Short: "Run an interactive service client",
			Run: func(cmd *cobra.Command, args []string) {
				initShells(cmd)
			},
		}
		rootCmd.AddCommand(cmd)
	}
	_ = rootCmd.Execute()
}

func initShells(cmd *cobra.Command) {
	shell := ishell.New()
	aShell, err := analytics.CreateAnalyticsClient(options.url, options.username)
	if err != nil {
		cmd.Println(err)
		return
	}
	authShell, err := auth.CreateAuthClient(options.url)
	if err != nil {
		cmd.Println(err)
		return
	}
	buddyShell, err := buddy.CreateBuddyClient(options.url)
	if err != nil {
		cmd.Println(err)
		return
	}

	matchmakingShell, err := client.CreateClient(options.url)
	if err != nil {
		cmd.Println(err)
		return
	}
	shell.AddCmd(matchmakingShell)
	shell.AddCmd(buddyShell)
	shell.AddCmd(authShell)
	shell.AddCmd(aShell)
	shell.Run()
}
