-- +goose Up
-- +goose StatementBegin

CREATE TABLE tasks(
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT,
  status VARCHAR(50) DEFAULT 'pending',
  user_id INT REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks
-- +goose StatementEnd
