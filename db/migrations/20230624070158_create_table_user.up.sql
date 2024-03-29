CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    username VARCHAR(100) NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS auths (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    token VARCHAR NOT NULL,
    auth_type VARCHAR,
    expired_at TIMESTAMPTZ NOT NULL,
    user_id BIGINT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
