CREATE TYPE order_status AS ENUM('NEW', 'PROCESSING', 'INVALID', 'PROCESSED');

CREATE TABLE IF NOT EXISTS user_order (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES "user"(id),
    status order_status NOT NULL DEFAULT 'NEW',
    number TEXT UNIQUE NOT NULL,
    accrual NUMERIC(10,2) NOT NULL DEFAULT 0,
    uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);