package commands

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/dmi3midd/papelane-cli/internal/config"
	"github.com/dmi3midd/papelane-cli/internal/domain"
	"github.com/dmi3midd/papelane-cli/internal/repositories"

	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

// simple mkdir command for creating a new directory
// example: mkdir folder (in vfs current directory)
// in development: mkdir path/to/folder (in vfs absolute path)
var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Create a new directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirName := config.CurrentDirConfig.GetString("currentDirName")
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		newDir := args[0]
		ctx := cmd.Context()
		ids := strings.Split(currentDirId, string(os.PathSeparator))
		parentId := ids[len(ids)-1]

		// check if the directory already exists in the current directory
		_, err := folderRepo.GetByNameAndParentId(ctx, newDir, parentId)
		if err == nil {
			return fmt.Errorf("directory '%s' already exists in the current directory", newDir)
		}
		if err != nil && !errors.Is(err, repositories.ErrFolderNotFound) {
			return fmt.Errorf("failed to get directory: %v", err)
		}

		// create a new directory
		folder := &domain.Folder{
			Id:        xid.New().String(),
			Name:      newDir,
			ParentId:  parentId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := folderRepo.Create(ctx, folder); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}

		// if the flag is set, go to the newly created directory
		if goToCurrDirFlag {
			newCurrDirName := path.Join(currentDirName, newDir)
			newCurrDirId := path.Join(currentDirId, folder.Id)
			if err := config.WriteOutCurrDirCfg(&config.CurrDirConfig{
				CurrentDirName: newCurrDirName,
				CurrentDirId:   newCurrDirId,
			}); err != nil {
				return err
			}
			fmt.Printf("directory '%s' created successfully\n", newCurrDirName)
		} else {
			fmt.Printf("directory '%s' created successfully\n", path.Join(currentDirName, newDir))
		}

		return nil
	},
}
