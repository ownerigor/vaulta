package cmd

import (
	"time"

	"github.com/ownerigor/vaulta/pkg/models"
	"github.com/ownerigor/vaulta/pkg/msg"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute scheduled backups based on the saved configuration",
	Run: func(cmd *cobra.Command, args []string) {
		RunScheduler()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runBackup(cfg *models.BackupConfig) {
	msg.Info("Running backup for database: %s", cfg.DBPath)

	//TODO: Issue #1
	time.Sleep(3 * time.Second)

	msg.Info("Backup completed successfully!")
	msg.Info("Backup restored at: %s", cfg.BackupPath)
}

func RunScheduler() {
	cfg := &models.BackupConfig{}
	if err := cfg.Load(); err != nil {
		msg.Err("Failed to load configuration: %v", err)
		return
	}

	msg.Info("Vaulta scheduler started")
	msg.Info("Database path: %s", cfg.DBPath)
	msg.Info("Interval (days): %d", cfg.IntervalDays)
	msg.Info("Backup time: %s", cfg.BackupHour)
	msg.Info("Backup destination: %s", cfg.BackupPath)

	for {
		now := time.Now()
		currentTime := now.Format("15:04")

		if currentTime == cfg.BackupHour {
			msg.Info("It's time to run the backup!")

			runBackup(cfg)

			next := time.Now().Add(time.Duration(cfg.IntervalDays*24) * time.Hour)
			msg.Info("Next backup scheduled for: %s", next.Format("2006-01-02 15:04"))
			time.Sleep(time.Duration(cfg.IntervalDays*24) * time.Hour)
		}
		time.Sleep(1 * time.Minute)
	}
}
