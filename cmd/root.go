package cmd

import (
	"github.com/ownerigor/vaulta/pkg/msg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vaulta",
	Short: "Vaulta is a simple backup tool",
	Run: func(cmd *cobra.Command, args []string) {
		msg.Info("Welcome to the Vaulta CLI!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
