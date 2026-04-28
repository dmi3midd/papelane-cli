package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clean cache command
// it cleans the cache of the Local Telegram Bot API (docker volume)
// example: cleanc
var cleancCmd = &cobra.Command{
	Use:   "cleanc",
	Short: "Clean the cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.CleanCache(); err != nil {
			return fmt.Errorf("failed to clean the cache: %v", err)
		}
		return nil
	},
}
