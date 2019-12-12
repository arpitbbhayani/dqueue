package cmd

import (
	"github.com/arpitbbhayani/dqueue/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run dqueue server",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		server.Run(configPath)
	},
}
