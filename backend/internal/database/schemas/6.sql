CREATE TABLE IF NOT EXISTS result (
    result_id serial PRIMARY KEY,
    election_id int NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    party text NOT NULL,
    votes int NOT NULL
);