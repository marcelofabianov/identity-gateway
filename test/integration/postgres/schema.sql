-- init_db.sql

-- Create the table realms
CREATE TABLE IF NOT EXISTS realms (
    id UUID PRIMARY KEY,
    identity_provider_id UUID UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    version INT NOT NULL
);

-- Create indexes for the realms table
CREATE INDEX IF NOT EXISTS idx_realms_identity_provider_id ON realms (identity_provider_id);

CREATE INDEX IF NOT EXISTS idx_realms_name ON realms (name);

-- Create the table users
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

-- Create indexes for the users table
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);

CREATE INDEX IF NOT EXISTS idx_users_document_registry ON users (document_registry);

CREATE INDEX IF NOT EXISTS idx_users_enabled ON users (enabled);
