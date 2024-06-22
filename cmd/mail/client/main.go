package main

import (
	"github.com/spf13/cobra"
)

const (
	defaultMail     = "localhost:8081"
	defaultUsername = "test"
)

var options struct {
	mail     string
	username string
}

func main() {

	rootCmd := &cobra.Command{
		Use:   "cond_cli",
		Short: "Run a mail CLI",
	}

	rootCmd.PersistentFlags().StringVar(&options.mail, "mail", defaultMail, "mail service (<host>:<port>)")
	rootCmd.PersistentFlags().StringVar(&options.username, "username", defaultUsername, "username for authentication")
	{
		shell := &cobra.Command{
			Use:   "shell",
			Short: "Run an interactive mail service client",
			Run: func(cmd *cobra.Command, args []string) {
				//TODO add client
			},
		}
		rootCmd.AddCommand(shell)
	}
	_ = rootCmd.Execute()
}
