-- +goose Up
CREATE TABLE files (
    id          TEXT PRIMARY KEY,
    tg_msg_id   INTEGER,
    tg_file_id  TEXT NOT NULL,
    parent_id   TEXT,
    name        TEXT NOT NULL,
    size        INTEGER,
    mime_type   TEXT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES folders (id) ON DELETE CASCADE
);

CREATE INDEX idx_files_parent ON files(parent_id);

-- +goose Down
DROP TABLE files;
