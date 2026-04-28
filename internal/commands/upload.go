package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dmi3midd/papelane-cli/internal/config"
	"github.com/dmi3midd/papelane-cli/internal/domain"

	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

// upload command for uploading a file to the current directory
// example: upload file.txt (on your local machine)
// in development: upload /path/to/file.txt (on your local machine)
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to the current directory",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDirId := config.CurrentDirConfig.GetString("currentDirId")
		ids := strings.Split(currentDirId, string(os.PathSeparator))
		filePath := args[0]
		ostat, err := os.Stat(filePath)
		if err != nil {
			return fmt.Errorf("Failed to get file info: %v", err)
		}
		if !ostat.Mode().IsRegular() {
			return fmt.Errorf("Path is not a regular file: %s", filePath)
		}
		uploadedFile, err := client.UploadFile(filePath, ostat)
		if err != nil {
			return fmt.Errorf("Failed to upload file: %v", err)
		}
		ctx := cmd.Context()
		file := domain.File{
			Id:        xid.New().String(),
			TgMsgId:   uploadedFile.TgMsgId,
			TgFileId:  uploadedFile.TgFileId,
			ParentId:  ids[len(ids)-1],
			Name:      uploadedFile.Name,
			Size:      uploadedFile.Size,
			MimeType:  uploadedFile.MimeType,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err = fileRepo.Create(ctx, &file); err != nil {
			return fmt.Errorf("Failed to create file record: %v", err)
		}

		fmt.Printf("File uploaded successfully: %s\n", uploadedFile.Name)
		return nil
	},
}
