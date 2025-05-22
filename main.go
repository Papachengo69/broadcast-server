package main

import (
	"broadcast-server/client"
	"broadcast-server/server"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "broadcast-server"}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "connect",
		Short: "Connect as client",
		Run: func(cmd *cobra.Command, args []string) {
			client.ConnectToServer()
		},
	})

	rootCmd.Execute()
}
