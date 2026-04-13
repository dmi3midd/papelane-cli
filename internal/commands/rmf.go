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

// simple rmf command for removing a file
// need to add complex logic for removing a file like:
// rmf root/folder1/folder2/file.txt (abs path) or rmf file.txt (relative path)
// also need(optional) to add a flag for going to the parent directory after removing the file
// now it removes the file in the current directory and if the flag is set, it goes to the parent directory
var rmfCmd = &cobra.Command{
	Use:   "rmf",
	Short: "Remove a file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		ids := strings.Split(currentDirId, string(os.PathSeparator))
		targetDir := args[0]

		ctx := cmd.Context()
		candidate, err := fileRepo.GetByNameAndParentId(ctx, targetDir, ids[len(ids)-1])
		if err != nil {
			if errors.Is(err, repositories.ErrFileNotFound) {
				return fmt.Errorf("File '%s' does not exist in the current directory", targetDir)
			}
			return fmt.Errorf("Failed to get file: %v", err)
		}
		if err := fileRepo.Delete(ctx, candidate.Id); err != nil {
			return fmt.Errorf("Failed to delete file: %v", err)
		}

		return nil
	},
}
