-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL,
    password TEXT NOT NULL,
    document_registry VARCHAR(14) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    version INT DEFAULT 1 NOT NULL,
    realm_id UUID,
    FOREIGN KEY (realm_id) REFERENCES realms (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);

CREATE INDEX IF NOT EXISTS idx_users_document_registry ON users (document_registry);

CREATE INDEX IF NOT EXISTS idx_users_enabled ON users (enabled);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_enabled;

DROP INDEX IF EXISTS idx_users_document_registry;

DROP INDEX IF EXISTS idx_users_email;

ALTER TABLE users DROP CONSTRAINT IF EXISTS users_realm_id_fkey;

DROP TABLE IF EXISTS users;
-- +goose StatementEnd
