CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login TEXT UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);