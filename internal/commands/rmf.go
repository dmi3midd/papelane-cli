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

// simple rmf command for removing a file
// example: rmf file.txt (in vfs current directory)
// in development: rmf path/to/file.txt (in vfs absolute path)
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
				return fmt.Errorf("file '%s' does not exist in the current directory", targetDir)
			}
			return fmt.Errorf("failed to get file: %v", err)
		}
		if err := fileRepo.Delete(ctx, candidate.Id); err != nil {
			return fmt.Errorf("failed to delete file: %v", err)
		}
		if err := client.DeleteFile(candidate.TgFileId); err != nil {
			return fmt.Errorf("failed to delete file from Telegram: %v", err)
		}

		return nil
	},
}
