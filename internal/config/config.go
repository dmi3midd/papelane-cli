package config

import (
	"fmt"

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
}

func WriteOut(c *Config) error {
	viper.SetConfigName("papelane.config")
	viper.SetConfigType("yaml")
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

	err := viper.SafeWriteConfigAs("papelane.config.yaml")
	if err != nil {
		return fmt.Errorf("error while writing config: %w", err)
	}
	return nil
}

func ReadIn() error {
	viper.SetConfigName("papelane.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error while loading config: %w", err)
	}
	return nil
}
