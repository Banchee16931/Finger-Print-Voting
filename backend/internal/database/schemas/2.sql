CREATE TABLE IF NOT EXISTS registrants (
    registrant_id SERIAL PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    phone_no text NOT NULL,
    fingerprint text NOT NULL,
    proof text NOT NULL,
    authority_location text NOT NULL
);