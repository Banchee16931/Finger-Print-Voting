CREATE TABLE IF NOT EXISTS elections (
    election_id serial PRIMARY KEY,
    election_start text NOT NULL,
    election_end text NOT NULL,
    authority_location text NOT NULL
);