-- +goose Up
CREATE TABLE orders (
    uuid UUID PRIMARY KEY,
    user_uuid UUID NOT NULL,
    total_price NUMERIC(15, 2) NOT NULL,
    status TEXT NOT NULL,
    items JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
);

-- +goose Down
DROP TABLE orders;