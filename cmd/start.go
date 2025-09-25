package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ownerigor/vaulta/pkg/models"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Configure and start backup scheduling",
	Run: func(cmd *cobra.Command, args []string) {
		runStart()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runStart() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the database path: ")
	dbPath := readLine(reader)

	fmt.Print("Enter the number of days between backups: ")
	var interval int
	fmt.Scanln(&interval)

	fmt.Print("Enter the backup time (HH:MM): ")
	backupHour := readLine(reader)

	fmt.Print("Enter the path where backups will be stored: ")
	backupPath := readLine(reader)

	cfg := &models.BackupConfig{
		DBPath:       dbPath,
		IntervalDays: interval,
		BackupHour:   backupHour,
		BackupPath:   backupPath,
	}

	if err := cfg.Save(); err != nil {
		fmt.Println("Failed to save configuration: ", err)
	}
}

func readLine(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}
