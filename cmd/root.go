package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dqueue",
		Short: "Persistent and fault tolerant queue",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	serverCmd.Flags().String("config", "/etc/dqueue/config.yml", "configuration file")
	rootCmd.AddCommand(serverCmd)

	clientCmd.Flags().String("protocol", "http", "which protocol to be used for communication")
	rootCmd.AddCommand(clientCmd)
}
