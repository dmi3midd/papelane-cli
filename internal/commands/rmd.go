package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dmi3midd/papelane-cli/internal/config"
	"github.com/dmi3midd/papelane-cli/internal/repositories"

	"github.com/spf13/cobra"
)

// simple rmd command for removing a directory
// example: rmd folder (in vfs current directory)
// in development: rmd path/to/folder (in vfs absolute path)
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
