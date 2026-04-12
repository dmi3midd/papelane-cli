package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	CurrentDirConfig *viper.Viper = viper.New()
	GlobalConfig     *viper.Viper = viper.New()
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

type CurrDirConfig struct {
	CurrentDirName string `yaml:"currentDirName"`
	CurrentDirId   string `yaml:"currentDirId"`
}

func getAppDir() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(cfgDir, "papelane-cli")
	return appDir, nil
}

func WriteOutGlobalCfg(c *Config) error {
	appDir, err := getAppDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return err
	}

	GlobalConfig.SetConfigName("papelane.config")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(appDir)
	GlobalConfig.AddConfigPath(".")

	GlobalConfig.Set("apiId", c.ApiId)
	GlobalConfig.Set("apiHash", c.ApiHash)
	GlobalConfig.Set("chatId", c.ChatId)
	GlobalConfig.Set("botToken", c.BotToken)
	GlobalConfig.Set("port", c.Port)
	GlobalConfig.Set("stopAlways", c.StopAlways)
	GlobalConfig.Set("image", c.Image)
	GlobalConfig.Set("containerName", c.ContainerName)
	GlobalConfig.Set("volume", c.Volume)
	GlobalConfig.Set("dbPath", c.DbPath)

	configPath := filepath.Join(appDir, "papelane.config.yaml")
	err = GlobalConfig.WriteConfigAs(configPath)
	if err != nil {
		return fmt.Errorf("error while writing config: %w", err)
	}
	return nil
}

func ReadInGlobalCfg() error {
	appDir, err := getAppDir()
	if err != nil {
		return err
	}

	GlobalConfig.SetConfigName("papelane.config")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(appDir)
	GlobalConfig.AddConfigPath(".")

	if err := GlobalConfig.ReadInConfig(); err != nil {
		return fmt.Errorf("error while loading config: %w", err)
	}
	return nil
}

func WriteOutCurrDirCfg(currDirCfg *CurrDirConfig) error {
	CurrentDirConfig.SetConfigName("currdir.papelane.config")
	CurrentDirConfig.SetConfigType("yaml")
	appDir, err := getAppDir()
	if err != nil {
		return err
	}
	CurrentDirConfig.AddConfigPath(appDir)
	CurrentDirConfig.AddConfigPath(".")

	CurrentDirConfig.Set("currentDirName", currDirCfg.CurrentDirName)
	CurrentDirConfig.Set("currentDirId", currDirCfg.CurrentDirId)

	configPath := filepath.Join(appDir, "currdir.papelane.config.yaml")
	err = CurrentDirConfig.WriteConfigAs(configPath)
	if err != nil {
		return fmt.Errorf("error while writing current directory config: %w", err)
	}
	return nil
}

func ReadInCurrDirCfg() error {
	appDir, err := getAppDir()
	if err != nil {
		return err
	}
	CurrentDirConfig.SetConfigName("currdir.papelane.config")
	CurrentDirConfig.SetConfigType("yaml")
	CurrentDirConfig.AddConfigPath(appDir)
	CurrentDirConfig.AddConfigPath(".")

	if err := CurrentDirConfig.ReadInConfig(); err != nil {
		return fmt.Errorf("error while loading current directory config: %w", err)
	}
	return nil
}
