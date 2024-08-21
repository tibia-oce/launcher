package config

import (
	"fmt"
	"os"
	"path/filepath"

	"launcher/internal/logger"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string
	LogLevel string
}

func LoadConfig(appName string) *Config {
	configDir := configDirectory(appName)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("Config file not found; using defaults")
		} else {
			logger.Error(fmt.Errorf("Error reading config file: %v", err))
		}
	}

	return &Config{
		AppName:  appName,
		LogLevel: viper.GetString("logger.level"),
	}
}

func (c *Config) SaveConfig() {
	configPath := filepath.Join(configDirectory(c.AppName), "config.toml")
	if err := viper.WriteConfigAs(configPath); err != nil {
		logger.Error(fmt.Errorf("Error writing config: %v", err))
	}
}

func configDirectory(appName string) string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logger.Error(fmt.Errorf("Error getting config directory: %v", err))
		return ""
	}
	return filepath.Join(configDir, appName)
}

func (c *Config) SetEnableLocal(value bool) {
	logger.Info(fmt.Sprintf("Setting enableLocal to %v", value))
	viper.Set("enableLocal", value)
	c.SaveConfig()
}

func (c *Config) LocalEnabled() bool {
	return viper.GetBool("enableLocal")
}
