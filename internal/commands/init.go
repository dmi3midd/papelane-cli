package commands

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/dmi3midd/papelane-cli/internal/config"
	"github.com/dmi3midd/papelane-cli/internal/logger"

	"github.com/spf13/cobra"
)

var (
	ErrGetFlags         = errors.New("failed to get flags")
	ErrGetUserConfigDir = errors.New("failed to get user config directory")
	ErrCreateAppDir     = errors.New("failed to create application config directory")
	ErrWriteGlobalCfg   = errors.New("failed to write global config")
	ErrWriteCurrDirCfg  = errors.New("failed to write current directory config")
)

// init command for initializing the application
// example: init --apid <api_id> --apih <api_hash> --token <bot_token> --cid <chat_id> --port <port> --sa <stop_always>
// port and sa flags are unneccessary, they're just for docker and set default values if not provided
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes application using your api hash and api id.",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiId, err := cmd.Flags().GetString("apid")
		if err != nil {
			logger.Log.Error("Error getting apid flag", "error", err)
			return ErrGetFlags
		}

		apiHash, err := cmd.Flags().GetString("apih")
		if err != nil {
			logger.Log.Error("Error getting apih flag", "error", err)
			return ErrGetFlags
		}

		chatId, err := cmd.Flags().GetInt("cid")
		if err != nil {
			logger.Log.Error("Error getting cid flag", "error", err)
			return ErrGetFlags
		}

		botToken, err := cmd.Flags().GetString("token")
		if err != nil {
			logger.Log.Error("Error getting token flag", "error", err)
			return ErrGetFlags
		}

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			logger.Log.Error("Error getting port flag", "error", err)
			return ErrGetFlags
		}

		stopAlways, err := cmd.Flags().GetBool("sa")
		if err != nil {
			logger.Log.Error("Error getting sa flag", "error", err)
			return ErrGetFlags
		}

		cfgDir, err := os.UserConfigDir()
		if err != nil {
			logger.Log.Error("Error getting user config dir", "error", err)
			return ErrGetUserConfigDir
		}

		appDir := filepath.Join(cfgDir, "papelane-cli")
		if err := os.MkdirAll(appDir, 0755); err != nil {
			logger.Log.Error("Error creating app config dir", "error", err)
			return ErrCreateAppDir
		}

		cfg := config.Config{
			ApiId:         apiId,
			ApiHash:       apiHash,
			ChatId:        chatId,
			BotToken:      botToken,
			Port:          port,
			StopAlways:    stopAlways,
			Image:         "aiogram/telegram-bot-api:latest",
			ContainerName: "papelane-telegram-bot-api",
			Volume:        "papelane-telegram-bot-api-data",
			DbPath:        filepath.Join(appDir, "papelane.sql"),
		}

		currDirCfg := config.CurrDirConfig{
			CurrentDirName: "root",
			CurrentDirId:   "root",
		}

		err = config.WriteOutGlobalCfg(&cfg)
		if err != nil {
			logger.Log.Error("Error writing global config", "error", err)
			return ErrWriteGlobalCfg
		}

		err = config.WriteOutCurrDirCfg(&currDirCfg)
		if err != nil {
			logger.Log.Error("Error writing current directory config", "error", err)
			return ErrWriteCurrDirCfg
		}

		logger.Log.Info("Initialization successful", "appDir", appDir)
		return nil
	},
}
