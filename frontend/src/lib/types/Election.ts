import type { Candidate, CandidateRequest, CandidateVotes } from "./Candidate"

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

export type ElectionState = {
	election_id: number
	start:      string  
	end:        string  
	location:   string
	result: CandidateVotes[]
}