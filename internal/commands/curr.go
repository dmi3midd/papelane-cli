package commands

import (
	"fmt"
	"os"
	"papelane-cli/internal/config"

	"github.com/spf13/cobra"
)

var currCmd = &cobra.Command{
	Use:   "curr",
	Short: "Prints the current directory in the Telegram Bot API (Local) storage.",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirName := config.CurrentDirConfig.GetString("currentDirName")
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		if currentDirName == "" || currentDirId == "" {
			fmt.Println("Current directory not set. Please run 'papelane init' to initialize the configuration.")
			os.Exit(1)
		}
		fmt.Printf("Current directory: %s\n", currentDirName)
		return nil
	},
}
