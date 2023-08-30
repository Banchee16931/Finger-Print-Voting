CREATE TABLE IF NOT EXISTS voter_details (
    username text PRIMARY KEY,
    email text NOT NULL,
    phone_no text NOT NULL,
    fingerprint text NOT NULL UNIQUE,
    authority_location text NOT NULL
);