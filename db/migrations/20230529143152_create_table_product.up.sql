CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    currency VARCHAR(50),
    total_stock INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    category_id BIGINT NOT NULL,
    FOREIGN KEY(category_id) REFERENCES categories(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
