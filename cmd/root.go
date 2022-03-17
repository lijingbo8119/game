package cmd

import (
	"game/client"
	"game/server"

	"github.com/spf13/cobra"
)

func Execute() {
	var cmdServer = &cobra.Command{
		Use:   "server",
		Short: "game server",
		Long:  `game server.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			server.Main()
		},
	}

	var cmdClient = &cobra.Command{
		Use:   "client",
		Short: "game client",
		Long:  `game client.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			client.Main()
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdServer, cmdClient)
	rootCmd.Execute()
}
