-- +goose Up
CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,

    task_list_id BIGINT NOT NULL,

    name TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_tasks_task_list
        FOREIGN KEY (task_list_id)
        REFERENCES task_lists(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tasks;
