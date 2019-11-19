package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run dqueue client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runs dqueue client")
	},
}
