CREATE TABLE IF NOT EXISTS user_balance (
    user_id uuid PRIMARY KEY REFERENCES "user"(id),
    current NUMERIC(10,2) NOT NULL DEFAULT 0,
    withdrawn NUMERIC(10,2) NOT NULL DEFAULT 0
);