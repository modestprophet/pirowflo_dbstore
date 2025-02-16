package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".pirowfloconfig.json"

type Config struct {
	DBURL        string `json:"db_url"`
	MqServerURL  string `json:"mq_server_url"`
	MqClientID   string `json:"mq_client_id"`
	MqDeviceName string `json:"mq_device_name"`
	MqUser       string `json:"mq_user"`
	MqPassword   string `json:"mq_password"`
	MqTopic      string `json:"mq_topic"`
}

func Read() (*Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config filepath: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config not found: %w", err) //maybe return an empty default config here nil:&Config{}
		}
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("error decoding config json: %w", err)
	}

	return &cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error obtaining user home directory path: %w", err)
	}
	return filepath.Join(home, configFileName), nil
}
