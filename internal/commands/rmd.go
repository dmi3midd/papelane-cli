package commands

import (
	"errors"
	"fmt"
	"os"
	"papelane-cli/internal/config"
	"papelane-cli/internal/repositories"
	"strings"

	"github.com/spf13/cobra"
)

// simple rmd command for removing a directory
// need to add complex logic for removing a directory like:
// rmd root/folder1/folder2/folder3 (abs path) or rmd folder3 (relative path)
// also need(optional) to add a flag for going to the parent directory after removing the directory
// now it removes the directory in the current directory and if the flag is set, it goes to the parent directory
var rmdCmd = &cobra.Command{
	Use:   "rmd",
	Short: "Remove a directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		ids := strings.Split(currentDirId, string(os.PathSeparator))
		targetDir := args[0]

		ctx := cmd.Context()
		candidate, err := folderRepo.GetByNameAndParentId(ctx, targetDir, ids[len(ids)-1])
		if err != nil {
			if errors.Is(err, repositories.ErrFolderNotFound) {
				return fmt.Errorf("Directory '%s' does not exist in the current directory", targetDir)
			}
			return fmt.Errorf("Failed to get directory: %v", err)
		}
		if err := folderRepo.Delete(ctx, candidate.Id); err != nil {
			return fmt.Errorf("Failed to delete directory: %v", err)
		}

		return nil
	},
}
