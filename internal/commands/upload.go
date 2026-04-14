package commands

import (
	"fmt"
	"os"
	"papelane-cli/internal/config"
	"papelane-cli/internal/domain"
	"strings"
	"time"

	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

// simple upload command for uploading a file to the current directory
// need to add complex logic for uploading a file like:
// upload root/folder1/folder2/file.txt (abs path) or upload file.txt (relative path)
// now it uploads the file to the current directory
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
