package commands

import (
	"fmt"
	"os"
	"papelane-cli/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

// cd command for changing the vfs current directory
// example: cd folder1/folder2
// example: cd ..
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change the current directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirName := config.CurrentDirConfig.GetString("currentDirName")
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		cdPath := args[0]
		cdPathParts := strings.Split(cdPath, string(os.PathSeparator))
		ids := strings.Split(currentDirId, string(os.PathSeparator))
		names := strings.Split(currentDirName, string(os.PathSeparator))
		for _, part := range cdPathParts {
			if part == "." {
				continue
			}
			if part == ".." {
				if len(ids) > 1 {
					ids = ids[:len(ids)-1]
					names = names[:len(names)-1]
				}
				continue
			}
			if len(ids) > 0 {
				candidate, err := folderRepo.GetByNameAndParentId(cmd.Context(), part, ids[len(ids)-1])
				if err != nil {
					return fmt.Errorf("Folder '%s' not found", part)
				}
				ids = append(ids, candidate.Id)
				names = append(names, candidate.Name)
				continue
			}
		}
		newCurrentDirId := strings.Join(ids, string(os.PathSeparator))
		newCurrentDirName := strings.Join(names, string(os.PathSeparator))
		if err := config.WriteOutCurrDirCfg(&config.CurrDirConfig{
			CurrentDirName: newCurrentDirName,
			CurrentDirId:   newCurrentDirId,
		}); err != nil {
			return fmt.Errorf("Failed to update current directory config: %w", err)
		}
		return nil
	},
}
