CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    username VARCHAR(125) NULL,
    email VARCHAR(125) NULL,
    phone VARCHAR(125) NULL,
    address VARCHAR(125) NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NULL
);

ALTER TABLE customers
ADD CONSTRAINT unique_email UNIQUE (email);