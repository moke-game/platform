package main

import (
	"context"

	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"

	"github.com/moke-game/platform/services/room/client"
)

var options struct {
	host string
	port int
}

const (
	DefaultHost = "localhost"
	DefaultPort = 8888
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "shell",
		Short:   "room client",
		Aliases: []string{"S"},
	}
	rootCmd.PersistentFlags().StringVar(
		&options.host,
		"host",
		DefaultHost,
		"grpc http service (<host>:<port>)",
	)

	rootCmd.PersistentFlags().IntVar(
		&options.port,
		"port",
		DefaultPort,
		"grpc http service (<host>:<port>)",
	)

	sGrpc := &cobra.Command{
		Use:   "shell",
		Short: "Run an interactive grpc client",
		Run: func(cmd *cobra.Command, args []string) {
			shell := ishell.New()
			if sh, err := client.CreateRoomWSClient(options.host, options.port); err != nil {
				panic(err)
			} else {
				shell.AddCmd(sh)
			}
			shell.Run()
		},
	}

	rootCmd.AddCommand(sGrpc)
	_ = rootCmd.ExecuteContext(context.Background())
}
