package cmd

import (
	"github.com/arpitbbhayani/dqueue/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run dqueue client",
	Run: func(cmd *cobra.Command, args []string) {
		client.Run()
	},
}
