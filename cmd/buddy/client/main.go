package main

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/moke-game/platform/services/buddy/client"
)

var options struct {
	host    string
	tcpHost string
}

const (
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
			client.RunBuddyCmd(options.host)
		},
	}

	rootCmd.AddCommand(sGrpc)
	_ = rootCmd.ExecuteContext(context.Background())
}
