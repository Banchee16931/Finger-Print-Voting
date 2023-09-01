CREATE TABLE IF NOT EXISTS candidates (
    candidate_id serial PRIMARY KEY,
    election_id int NOT NULL,
    first_name text NOT NULL,  
    last_name text NOT NULL,   
    party text NOT NULL,       
    party_colour text NOT NULL,
    photo text NOT NULL        
);