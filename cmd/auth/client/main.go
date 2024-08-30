package main

import (
	"context"

	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"

	auth "github.com/moke-game/platform/services/auth/client"
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
		Short:   "auth client",
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
			initShell(cmd)
		},
	}

	rootCmd.AddCommand(sGrpc)
	_ = rootCmd.ExecuteContext(context.Background())
}

func initShell(cmd *cobra.Command) {
	shell := ishell.New()
	authShell, err := auth.CreateAuthClient(options.host)
	if err != nil {
		cmd.Println(err)
		return
	}
	shell.AddCmd(authShell)
	shell.Run()
}
