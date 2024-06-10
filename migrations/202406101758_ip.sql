-- +goose Up
-- +goose StatementBegin
alter table sessions
    add ip inet;


-- +goose Down
-- +goose StatementBegin
alter table sessions
    drop column ip;
-- +goose StatementEnd