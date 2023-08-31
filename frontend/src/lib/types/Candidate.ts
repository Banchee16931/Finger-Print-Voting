export type Candidate = {
	candidate_id: number;
	first_name:   string;
	last_name:    string;
	party: string;
	photo:       string;
}

export type CandidateRequest = {
	first_name:   string;
	last_name:    string;
	party: string;
	photo:       string;
}