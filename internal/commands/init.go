package commands

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dmi3midd/papelane-cli/internal/config"

	"github.com/spf13/cobra"
)

// init command for initializing the application
// example: init --apid <api_id> --apih <api_hash> --token <bot_token> --cid <chat_id> --port <port> --sa <stop_always>
// port and sa flags are unneccessary, they're just for docker and set default values if not provided
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes application using your api hash and api id.",
	Run: func(cmd *cobra.Command, args []string) {
		apiId, err := cmd.Flags().GetString("apid")
		apiHash, err := cmd.Flags().GetString("apih")
		chatId, err := cmd.Flags().GetInt("cid")
		botToken, err := cmd.Flags().GetString("token")
		port, err := cmd.Flags().GetInt("port")
		stopAlways, err := cmd.Flags().GetBool("sa")
		if err != nil {
			log.Fatalf("Error while execute init cmd: %v", err)
		}

		cfgDir, err := os.UserConfigDir()
		if err != nil {
			log.Fatalf("Error getting user config dir: %v", err)
		}
		appDir := filepath.Join(cfgDir, "papelane-cli")
		if err := os.MkdirAll(appDir, 0755); err != nil {
			log.Fatalf("Error creating app config dir: %v", err)
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
			log.Fatalf("Error while execute init cmd: %v", err)
		}

		err = config.WriteOutCurrDirCfg(&currDirCfg)
		if err != nil {
			log.Fatalf("Error while execute init cmd: %v", err)
		}
		err = client.Ping()
		if err != nil {
			log.Fatalf("Error while execute init cmd: %v", err)
		}
	},
}
