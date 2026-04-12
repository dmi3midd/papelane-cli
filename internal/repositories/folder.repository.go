package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"papelane-cli/internal/domain"
	"path"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	ErrFolderNotFound = errors.New("folder not found")
)

// FolderRepository provides methods to interact with the folders in the database.
type FolderRepository struct {
	db *sqlx.DB
}

func NewFolderRepository(db *sqlx.DB) *FolderRepository {
	return &FolderRepository{db: db}
}

func (r *FolderRepository) GetById(ctx context.Context, id string) (*domain.Folder, error) {
	op := "FolderRepository.GetById"
	query := `
	SELECT id, name, parent_id, created_at, updated_at 
	FROM folders 
	WHERE id = $1
	`
	var folder domain.Folder
	err := r.db.GetContext(ctx, &folder, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, ErrFolderNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &folder, nil
}

func (r *FolderRepository) GetByNameAndParentId(ctx context.Context, name string, parentId string) (*domain.Folder, error) {
	op := "FolderRepository.GetByNameAndParentId"
	query := `
	SELECT id, name, parent_id, created_at, updated_at 
	FROM folders 
	WHERE name = $1 AND parent_id = $2
	`
	var folder domain.Folder
	err := r.db.GetContext(ctx, &folder, query, name, parentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, ErrFolderNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &folder, nil
}

func (r *FolderRepository) GetByParentId(ctx context.Context, parentId string) ([]domain.Folder, error) {
	op := "FolderRepository.GetByParentId"
	query := `
	SELECT id, name, parent_id, created_at, updated_at 
	FROM folders 
	WHERE parent_id = $1
	`
	var folders []domain.Folder
	err := r.db.SelectContext(ctx, &folders, query, parentId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return folders, nil
}

func (r *FolderRepository) GetByPath(ctx context.Context, fullPath string) (*domain.Folder, error) {
	op := "FolderRepository.GetByPath"

	if fullPath == "root" {
		return &domain.Folder{
			Id:       "root",
			Name:     "root",
			ParentId: "",
		}, nil
	}

	segments := strings.Split(path.Clean(fullPath), string(os.PathSeparator))
	if len(segments) > 0 && segments[0] == "root" {
		segments = segments[1:]
	}

	currParentId := "root"
	var folder *domain.Folder
	var err error

	for _, segment := range segments {
		if segment == "" {
			continue
		}
		folder, err = r.GetByNameAndParentId(ctx, segment, currParentId)
		if err != nil {
			if errors.Is(err, ErrFolderNotFound) {
				return nil, fmt.Errorf("%s: %w", op, ErrFolderNotFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		currParentId = folder.Id
	}

	if folder == nil {
		return nil, fmt.Errorf("%s: %w", op, ErrFolderNotFound)
	}

	return folder, nil
}

func (r *FolderRepository) Create(ctx context.Context, folder *domain.Folder) error {
	op := "FolderRepository.Create"
	query := `
	INSERT INTO folders (id, name, parent_id, created_at, updated_at) 
	VALUES (:id, :name, :parent_id, :created_at, :updated_at)
	`
	if _, err := r.db.NamedExecContext(ctx, query, folder); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *FolderRepository) Update(ctx context.Context, folder *domain.Folder) error {
	op := "FolderRepository.Update"
	query := `
	UPDATE folders 
	SET name = :name, parent_id = :parent_id, updated_at = :updated_at 
	WHERE id = :id
	`
	result, err := r.db.NamedExecContext(ctx, query, folder)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, ErrFolderNotFound)
	}
	return nil
}

func (r *FolderRepository) Delete(ctx context.Context, id string) error {
	op := "FolderRepository.Delete"
	query := `
	DELETE FROM folders 
	WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, ErrFolderNotFound)
	}
	return nil
}
