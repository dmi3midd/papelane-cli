package commands

import (
	"fmt"
	"os"
	"papelane-cli/internal/config"
	"papelane-cli/internal/domain"
	"strings"

	"github.com/spf13/cobra"
)

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
				return fmt.Errorf("Failed to get files: %w", err)
			}
			folders, err = folderRepo.GetByParentId(ctx, ids[len(ids)-1])
			if err != nil {
				return fmt.Errorf("Failed to get folders: %w", err)
			}
		}

		if filesFlag {
			files, err = fileRepo.GetByParentId(ctx, ids[len(ids)-1])
			if err != nil {
				return fmt.Errorf("Failed to get files: %w", err)
			}
		}

		if dirsFlag {
			folders, err = folderRepo.GetByParentId(ctx, ids[len(ids)-1])
			if err != nil {
				return fmt.Errorf("Failed to get folders: %w", err)
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
