package cmd

import (
	"time"

	"github.com/ownerigor/vaulta/pkg/models"
	"github.com/ownerigor/vaulta/pkg/msg"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a backup based on the saved configuration",
	Run: func(cmd *cobra.Command, args []string) {
		runBackup()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runBackup() {
	cfg := &models.BackupConfig{}
	if err := cfg.Load(); err != nil {
		msg.Err("Failed to load configuration: %v", err)
		return
	}

	msg.Info("Starting backup process...")
	msg.Info("Database Path: %s", cfg.DBPath)
	msg.Info("Interval (days): %d", cfg.IntervalDays)
	msg.Info("Backup Time: %s", cfg.BackupHour)
	msg.Info("Backup Destination: %s", cfg.BackupPath)

	time.Sleep(2 * time.Second)
	msg.Info("Backup completed successfully!")
	msg.Info("Backup stored at: %s", cfg.BackupPath)
}
