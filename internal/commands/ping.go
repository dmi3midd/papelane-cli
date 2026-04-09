package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"papelane-cli/internal/config"
	"papelane-cli/internal/telegrampkg"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if the Telegram Bot API is ready",
	Long:  `Check if the Telegram Bot API is ready by sending a request to the /getMe endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.ReadIn(); err != nil {
			log.Fatalf("Failed to read config: %v", err)
		}

		client = telegrampkg.NewClient(
			viper.GetString("botToken"),
			fmt.Sprintf("http://localhost:%d", viper.GetInt("port")),
		)

		err := client.Ping()
		if err != nil {
			log.Fatalf("Error while execute ping cmd: %v", err)
		}
		fmt.Println("Telegram Bot API is ready")
	},
}
