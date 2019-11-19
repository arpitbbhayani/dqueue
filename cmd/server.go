package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run dqueue server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runs dqueue server")
	},
}
