package commands

import (
	"fmt"
	"os"
	"papelane-cli/internal/config"
	"path"

	"github.com/spf13/cobra"
)

// simple mkdir command for creating a new directory
// need to add complex logic for creating a new directory like:
// mkdir root/folder1/folder2/folder3
// in this situation we need to create folder3 in root/folder1/folder2
// and change current directory to root/folder1/folder2/folder3
// and if folder1 or folder2 is not exist, we need to return an error
var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Create a new directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDir := config.CurrentDirConfig.GetString("currentDir")
		if currentDir == "" {
			fmt.Println("Current directory not set. Please run 'papelane init' to initialize the configuration.")
			os.Exit(1)
		}
		newDir := args[0]
		newCurrDir := path.Join(currentDir, newDir)
		if err := config.WriteOutCurrDirCfg(&config.CurrDirConfig{
			CurrentDir: newCurrDir,
		}); err != nil {
			return err
		}
		fmt.Printf("Directory '%s' created successfully.\n", newDir)
		fmt.Printf("Current directory: %s\n", newCurrDir)
		return nil
	},
}
