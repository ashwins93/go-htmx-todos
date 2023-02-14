-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
  id TEXT,
  task TEXT,
  completed_at INTEGER,
  created_at INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todos;
-- +goose StatementEnd
