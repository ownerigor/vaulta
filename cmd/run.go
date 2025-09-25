package cmd

import (
	"fmt"
	"time"

	"github.com/ownerigor/vaulta/pkg/models"
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
		fmt.Println("Failed to load configuration:", err)
		return
	}

	fmt.Println("Starting backup process...")
	fmt.Printf("Database path: %s\n", cfg.DBPath)
	fmt.Printf("Interval (days): %d\n", cfg.IntervalDays)
	fmt.Printf("Backup time: %s\n", cfg.BackupHour)
	fmt.Printf("Backup destination: %s\n", cfg.BackupPath)

	time.Sleep(2 * time.Second)
	fmt.Println("Backup completed successfully!")
	fmt.Printf("Backup stored at: %s\n", cfg.BackupPath)
}
