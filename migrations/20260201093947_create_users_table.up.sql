CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,

    role user_role NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by UUID NULL,

    updated_at TIMESTAMPTZ NULL,
    updated_by UUID NULL,

    deleted_at TIMESTAMPTZ NULL,
    deleted_by UUID NULL
);

CREATE INDEX idx_users_deleted_at ON users(deleted_at);