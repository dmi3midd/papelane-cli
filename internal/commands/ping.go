package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// ping command for checking if the Telegram Bot API is ready
// example: ping
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if the Telegram Bot API is ready",
	Long:  `Check if the Telegram Bot API is ready by sending a request to the /getMe endpoint.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if client == nil {
			log.Fatalf("Client is not initialized")
		}
		err := client.Ping()
		if err != nil {
			return fmt.Errorf("failed to ping: %v", err)
		}
		fmt.Println("Telegram Bot API is ready")
		return nil
	},
}
