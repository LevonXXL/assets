-- +goose Up
-- +goose StatementBegin
ALTER TABLE sessions ADD CONSTRAINT fk_sessions_users FOREIGN KEY (uid) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE assets ADD CONSTRAINT fk_assets_users FOREIGN KEY (uid) REFERENCES users (id) ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sessions DROP CONSTRAINT IF EXISTS fk_sessions_users;
ALTER TABLE sessions DROP CONSTRAINT IF EXISTS fk_assets_users;
-- +goose StatementEnd