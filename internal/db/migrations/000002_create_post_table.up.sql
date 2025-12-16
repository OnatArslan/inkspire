CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE
    posts (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        title VARCHAR(50) UNIQUE NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    )
