package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleancCmd = &cobra.Command{
	Use:   "cleanc",
	Short: "Clean the cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := client.CleanCache(); err != nil {
			return fmt.Errorf("failed to clean the cache: %v", err)
		}
		fmt.Println("Cache cleaned successfully")
		return nil
	},
}
