-- +goose Up
INSERT OR IGNORE INTO folders (id, name, parent_id) VALUES ('root', 'root', NULL);

-- +goose Down
DELETE FROM folders WHERE id = 'root';
