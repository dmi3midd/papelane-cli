package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"papelane-cli/internal/domain"

	"github.com/jmoiron/sqlx"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

// FileRepository provides methods to interact with the files in the database.
type FileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) GetById(ctx context.Context, id string) (*domain.File, error) {
	op := "FileRepository.GetById"
	query := `
	SELECT id, tg_msg_id, tg_file_id, parent_id, name, size, mime_type, created_at, updated_at 
	FROM files 
	WHERE id = $1
	`
	var file domain.File
	err := r.db.GetContext(ctx, &file, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, ErrFileNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &file, nil
}

func (r *FileRepository) GetByParentId(ctx context.Context, parentId string) ([]domain.File, error) {
	op := "FileRepository.GetByParentId"
	query := `
	SELECT id, tg_msg_id, tg_file_id, parent_id, name, size, mime_type, created_at, updated_at 
	FROM files 
	WHERE parent_id = $1
	`
	var files []domain.File
	err := r.db.SelectContext(ctx, &files, query, parentId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return files, nil
}

func (r *FileRepository) Create(ctx context.Context, file *domain.File) error {
	op := "FileRepository.Create"
	query := `
	INSERT INTO files (id, tg_msg_id, tg_file_id, parent_id, name, size, mime_type, created_at, updated_at) 
	VALUES (:id, :tg_msg_id, :tg_file_id, :parent_id, :name, :size, :mime_type, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, file)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *FileRepository) Update(ctx context.Context, file *domain.File) error {
	op := "FileRepository.Update"
	query := `
	UPDATE files 
	SET tg_msg_id = :tg_msg_id, tg_file_id = :tg_file_id, parent_id = :parent_id, name = :name, size = :size, mime_type = :mime_type, created_at = :created_at, updated_at = :updated_at
	WHERE id = :id
	`
	result, err := r.db.NamedExecContext(ctx, query, file)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, ErrFileNotFound)
	}
	return nil
}

func (r *FileRepository) Delete(ctx context.Context, id string) error {
	op := "FileRepository.Delete"
	query := `
	DELETE FROM files 
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
		return fmt.Errorf("%s: %w", op, ErrFileNotFound)
	}
	return nil
}
