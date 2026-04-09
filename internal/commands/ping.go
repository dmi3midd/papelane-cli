package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if the Telegram Bot API is ready",
	Long:  `Check if the Telegram Bot API is ready by sending a request to the /getMe endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		if client == nil {
			log.Fatalf("Client is not initialized")
		}
		err := client.Ping()
		if err != nil {
			log.Fatalf("Error while execute ping cmd: %v", err)
		}
		fmt.Println("Telegram Bot API is ready")
	},
}
