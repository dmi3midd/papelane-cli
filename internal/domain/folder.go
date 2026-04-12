package domain

import (
	"context"
	"time"
)

// Folder represents a folder in the virtual filesystem.
// It contains metadata about the folder.
type Folder struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	ParentId  string    `json:"parentId" db:"parent_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// FolderRepository defines the interface for folder repository operations.
type FolderRepository interface {
	// GetById returns a folder by its ID.
	// It returns the ErrFolderNotFound error if the folder does not exist.
	GetById(ctx context.Context, id string) (*Folder, error)
	// GetByNameAndParentId returns a folder by its name and parent ID.
	// It returns the ErrFolderNotFound error if the folder does not exist.
	GetByNameAndParentId(ctx context.Context, name string, parentId string) (*Folder, error)
	// GetByParentId returns a list of folders by their parent ID.
	// It returns an empty slice if no folders are found with the given parent ID.
	GetByParentId(ctx context.Context, parrentId string) ([]Folder, error)
	// GetByPath returns a folder by its path.
	// It returns the ErrFolderNotFound error if the folder does not exist.
	GetByPath(ctx context.Context, path string) (*Folder, error)
	// Create adds a new folder to the database.
	Create(ctx context.Context, folder *Folder) error
	// Update modifies an existing folder in the database.
	// It returns the ErrFolderNotFound error if the folder does not exist.
	Update(ctx context.Context, folder *Folder) error
	// Delete removes a folder from the database by its ID.
	// It returns the ErrFolderNotFound error if the folder does not exist.
	Delete(ctx context.Context, id string) error
}
