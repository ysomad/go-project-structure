-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS test (
    id serial PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE test;
-- +goose StatementEnd
