package main

import (
	"context"

	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"

	"github.com/moke-game/platform/services/buddy/client"
)

var options struct {
	host    string
	tcpHost string
}

const (
	// DefaultHost default host
	DefaultHost = "localhost:8081"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "shell ",
		Short:   "buddy client",
		Aliases: []string{"cli"},
	}
	rootCmd.PersistentFlags().StringVar(
		&options.host,
		"host",
		DefaultHost,
		"grpc http service (<host>:<port>)",
	)

	sGrpc := &cobra.Command{
		Use:   "shell",
		Short: "Run an interactive grpc client",
		Run: func(cmd *cobra.Command, args []string) {
			shell := ishell.New()
			buddyClient, err := client.CreateBuddyClient(options.host)
			if err != nil {
				cmd.Println(err)
				return
			}
			shell.AddCmd(buddyClient)
			shell.Run()
		},
	}

	rootCmd.AddCommand(sGrpc)
	_ = rootCmd.ExecuteContext(context.Background())
}
