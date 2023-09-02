export type Vote = {
	username:   string
	election_id:  number  
	candidate_id: number  
}

export type VoteRequest = {
	election_id:  number  
	candidate_id: number
	fingerprint: string
}
