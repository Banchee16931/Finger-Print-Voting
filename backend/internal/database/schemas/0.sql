CREATE TABLE IF NOT EXISTS users (
    username text PRIMARY KEY,
    encrypted_password text NOT NULL,
    is_admin boolean NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL
);