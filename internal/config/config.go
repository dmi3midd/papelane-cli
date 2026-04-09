package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	ApiId         string `yaml:"apiId"`
	ApiHash       string `yaml:"apiHash"`
	ChatId        int    `yaml:"chatId"`
	BotToken      string `yaml:"botToken"`
	Port          int    `yaml:"port"`
	StopAlways    bool   `yaml:"stopAlways"`
	Image         string `yaml:"image"`
	ContainerName string `yaml:"containerName"`
	Volume        string `yaml:"volume"`
	DbPath        string `yaml:"dbPath"`
}

func getAppDir() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(cfgDir, "papelane-cli")
	return appDir, nil
}

func WriteOut(c *Config) error {
	appDir, err := getAppDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return err
	}

	viper.SetConfigName("papelane.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(appDir)
	viper.AddConfigPath(".")

	viper.Set("apiId", c.ApiId)
	viper.Set("apiHash", c.ApiHash)
	viper.Set("chatId", c.ChatId)
	viper.Set("botToken", c.BotToken)
	viper.Set("port", c.Port)
	viper.Set("stopAlways", c.StopAlways)
	viper.Set("image", c.Image)
	viper.Set("containerName", c.ContainerName)
	viper.Set("volume", c.Volume)
	viper.Set("dbPath", c.DbPath)

	configPath := filepath.Join(appDir, "papelane.config.yaml")
	err = viper.WriteConfigAs(configPath)
	if err != nil {
		return fmt.Errorf("error while writing config: %w", err)
	}
	return nil
}

func ReadIn() error {
	appDir, err := getAppDir()
	if err != nil {
		return err
	}

	viper.SetConfigName("papelane.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(appDir)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error while loading config: %w", err)
	}
	return nil
}
