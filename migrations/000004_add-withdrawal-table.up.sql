CREATE TABLE IF NOT EXISTS user_withdrawal
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES "user"(id),
    "order" TEXT NOT NULL REFERENCES user_order(number),
    "sum" NUMERIC(10,2) NOT NULL,
    processed_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);