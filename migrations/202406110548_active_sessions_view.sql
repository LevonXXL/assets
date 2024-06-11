-- +goose Up
-- +goose StatementBegin
CREATE VIEW active_sessions AS
    SELECT DISTINCT ON (uid) uid, created_at, id, ip
    FROM sessions
    ORDER BY uid, created_at DESC;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS active_sessions;
-- +goose StatementEnd