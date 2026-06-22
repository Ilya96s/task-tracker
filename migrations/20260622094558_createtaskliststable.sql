-- +goose Up
CREATE TABLE task_lists (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE task_lists;;
