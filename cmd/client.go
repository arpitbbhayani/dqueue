package cmd

import (
	"github.com/arpitbbhayani/dqueue/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run dqueue client",
	Run: func(cmd *cobra.Command, args []string) {
		protocol, _ := cmd.Flags().GetString("protocol")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		client.Run(protocol, host, port)
	},
}
