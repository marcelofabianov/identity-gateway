-- init_db.sql
CREATE TABLE IF NOT EXISTS realms (
    id UUID PRIMARY KEY,
    identity_provider_id UUID UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    version INT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_realms_identity_provider_id ON realms (identity_provider_id);

CREATE INDEX IF NOT EXISTS idx_realms_name ON realms (name);
