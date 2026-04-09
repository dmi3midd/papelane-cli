package commands

import (
	"fmt"
	"log"
	"papelane-cli/internal/config"
	"papelane-cli/internal/telegrampkg"

	"github.com/spf13/cobra"
)

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
		}
		err = config.WriteOut(&cfg)
		if err != nil {
			log.Fatalf("Error while execute init cmd: %v", err)
		}
		client = telegrampkg.NewClient(botToken, fmt.Sprintf("http://localhost:%d", port))
		err = client.Ping()
		if err != nil {
			log.Fatalf("Error while execute init cmd: %v", err)
		}
	},
}
