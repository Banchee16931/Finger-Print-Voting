CREATE TABLE IF NOT EXISTS votes (
    username text NOT NULL,
    election_id text NOT NULL,
    candidate_id text NOT NULL,
    PRIMARY KEY(username, election_id, candidate_id)
);