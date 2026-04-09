-- +goose Up
CREATE TABLE folders (
    id         TEXT PRIMARY KEY,
    name       TEXT NOT NULL,
    parent_id  TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES folders (id) ON DELETE CASCADE
);

CREATE INDEX idx_folders_parent ON folders(parent_id);

-- +goose Down
DROP TABLE folders;
