import type { Candidate, CandidateRequest } from "./Candidate"

export type Election = {
	election_id: number       
	start:      string  
	end:        string  
	location:   string     
	candidates: Candidate[]
}

export type ElectionRequest = {       
	start:      string  
	end:        string  
	location:   string     
	candidates: CandidateRequest[]
}
