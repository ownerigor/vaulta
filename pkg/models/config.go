package models

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFile = "config.json"

type BackupConfig struct {
	DBPath       string `json:"db_path"`
	IntervalDays int    `json:"interval_days"`
	BackupHour   string `json:"backup_hour"`
	BackupPath   string `json:"backup_path"`
}

func (cfg *BackupConfig) Save() error {
	marshal, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFile, marshal, 0600); err != nil {
		return err
	}

	fmt.Println("Configuration saved succesfully to", configFile)
	return nil
}

func (cfg *BackupConfig) Load() error {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, cfg); err != nil {
		return err
	}

	fmt.Println("Configuration loaded from", configFile)
	return nil
}
