package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "observer [config-file]",
	Args:  cobra.ExactArgs(1),
	Short: "start observing events",
	RunE:  observerHandler,
}

func observerHandler(cmd *cobra.Command, args []string) error {
	return nil
}
