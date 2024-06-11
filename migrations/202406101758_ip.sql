-- +goose Up
-- +goose StatementBegin
alter table sessions
    add ip inet;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table sessions
    drop column ip;
-- +goose StatementEnd