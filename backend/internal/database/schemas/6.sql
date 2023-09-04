CREATE TABLE IF NOT EXISTS votes (
    username text NOT NULL,
    election_id int NOT NULL,
    candidate_id int NOT NULL,
    PRIMARY KEY(username, election_id, candidate_id),
    FOREIGN KEY (username) REFERENCES users(username),
    FOREIGN KEY (candidate_id) REFERENCES candidates(candidate_id),
    FOREIGN KEY (election_id) REFERENCES elections(election_id)
);