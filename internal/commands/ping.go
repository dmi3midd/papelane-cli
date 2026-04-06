package commands

import (
	"papelane-cli/internal/telegrampkg"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if the Telegram Bot API is ready",
	Long:  `Check if the Telegram Bot API is ready by sending a request to the /getMe endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		telegrampkg.Ping()
	},
}
