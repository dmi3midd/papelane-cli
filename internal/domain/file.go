package domain

import (
	"context"
	"time"
)

// File represents a file in the virtual filesystem.
// It contains metadata about the file, especially Telegram file information such as message ID and file ID.
type File struct {
	Id        string    `json:"id" db:"id"`
	TgMsgId   int       `json:"tgMsgId" db:"tg_msg_id"`
	TgFileId  string    `json:"tgFileId" db:"tg_file_id"`
	ParentId  string    `json:"parentId" db:"parent_id"`
	Name      string    `json:"name" db:"name"`
	Size      int       `json:"size" db:"size"`
	MimeType  string    `json:"mimeType" db:"mime_type"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// FileRepository defines the interface for file repository operations.
type FileRepository interface {
	// GetById returns a file by its ID.
	// It returns the ErrFileNotFound error if the file does not exist.
	GetById(ctx context.Context, id string) (*File, error)
	// GetByNameAndParentId returns a file by its name and parent ID.
	// It returns the ErrFileNotFound error if the file does not exist.
	GetByNameAndParentId(ctx context.Context, name string, parentId string) (*File, error)
	// GetByParentId returns a list of files by their parent ID.
	// It returns an empty slice if no files are found with the given parent ID.
	GetByParentId(ctx context.Context, parentId string) ([]File, error)
	// Create adds a new file to the database.
	Create(ctx context.Context, file *File) error
	// Update modifies an existing file in the database.
	// It returns the ErrFileNotFound error if the file does not exist.
	Update(ctx context.Context, file *File) error
	// Delete removes a file from the database by its ID.
	// It returns the ErrFileNotFound error if the file does not exist.
	Delete(ctx context.Context, id string) error
}
