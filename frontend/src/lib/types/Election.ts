import type { Candidate } from "./Candidate"

export type Election = {
	election_id: number       
	start:      Date  
	end:        Date  
	location:   string     
	candidates: Candidate[]
}
