CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    "comments" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        content TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );
