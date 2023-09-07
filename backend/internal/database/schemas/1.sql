CREATE TABLE IF NOT EXISTS voter_details (
    username text PRIMARY KEY,
    email text NOT NULL,
    phone_no text NOT NULL,
    fingerprint text NOT NULL,
    authority_location text NOT NULL,
    FOREIGN KEY (username) REFERENCES users(username)
);