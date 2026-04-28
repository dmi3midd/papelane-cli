package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/dmi3midd/papelane-cli/internal/config"
	"github.com/dmi3midd/papelane-cli/internal/domain"

	"github.com/spf13/cobra"
)

// ls command for listing the contents of the vfs current directory
// example: ls -f (list files only)
// example: ls -d (list directories only)
// example: ls (list both files and directories)
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List directories in the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		ids := strings.Split(currentDirId, string(os.PathSeparator))
		ctx := cmd.Context()
		var (
			folders []domain.Folder
			files   []domain.File
			err     error
		)
		if !filesFlag && !dirsFlag {
			files, err = fileRepo.GetByParentId(ctx, ids[len(ids)-1])
			if err != nil {
				return fmt.Errorf("failed to get files: %w", err)
			}
			folders, err = folderRepo.GetByParentId(ctx, ids[len(ids)-1])
			if err != nil {
				return fmt.Errorf("failed to get folders: %w", err)
			}
		}

		if filesFlag {
			files, err = fileRepo.GetByParentId(ctx, ids[len(ids)-1])
			if err != nil {
				return fmt.Errorf("failed to get files: %w", err)
			}
		}

		if dirsFlag {
			folders, err = folderRepo.GetByParentId(ctx, ids[len(ids)-1])
			if err != nil {
				return fmt.Errorf("failed to get folders: %w", err)
			}
		}

		for _, folder := range folders {
			fmt.Printf("🗁  %s\n", folder.Name)
		}
		for _, file := range files {
			fmt.Printf("🗈  %s\n", file.Name)
		}
		return nil
	},
}
