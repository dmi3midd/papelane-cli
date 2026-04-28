package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dmi3midd/papelane-cli/internal/config"
	"github.com/dmi3midd/papelane-cli/internal/repositories"

	"github.com/spf13/cobra"
)

// download command for downloading a file from the current directory.
// example: download file.txt (default path is ~/Downloads)
// example: download file.txt --out /path/to/download/or/other/folder
// in development: download root/folder1/folder2/file.txt (abs path)
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a file from the current directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		ids := strings.Split(currentDirId, string(os.PathSeparator))
		target := args[0]
		ctx := cmd.Context()

		candidate, err := fileRepo.GetByNameAndParentId(ctx, target, ids[len(ids)-1])
		if err != nil {
			if errors.Is(err, repositories.ErrFileNotFound) {
				return fmt.Errorf("File not found in the current directory")
			}
			return fmt.Errorf("Failed to get file info: %v", err)
		}

		out, err := cmd.Flags().GetString("out")
		if err != nil {
			return fmt.Errorf("Failed to get out path: %v", err)
		}
		var dest string
		if out == "" {
			dest, err = os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("Failed to get home path: %v", err)
			}
			dest = filepath.Join(dest, "Downloads")
		} else {
			dest = out
		}

		if err := client.DownloadFile(candidate.TgFileId, dest); err != nil {
			return fmt.Errorf("failed to download file from Telegram: %v", err)
		}
		return nil
	},
}
