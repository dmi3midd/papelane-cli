package commands

import (
	"papelane-cli/internal/config"

	"github.com/spf13/cobra"
)

// toRootCmd is a command for setting the current directory to root
// example: root
var toRootCmd = &cobra.Command{
	Use:   "root",
	Short: "Set the current directory in the storage to root",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.WriteOutCurrDirCfg(&config.CurrDirConfig{
			CurrentDirName: "root",
			CurrentDirId:   "root",
		}); err != nil {
			return err
		}
		return nil
	},
}
