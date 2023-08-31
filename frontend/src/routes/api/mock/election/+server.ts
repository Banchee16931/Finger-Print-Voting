import type { RequestHandler } from '@sveltejs/kit';
import type { Election, ElectionRequest, Candidate } from '$lib/types';
import { elections } from "../data"
import { Authority } from '../handleAuth';
import { NewError } from '$lib/types/CommonError';

let electionCount = 0;

export const POST: RequestHandler = async (e) => {
    let req = e.request
    
    let authHeader = req.headers.get("Autherization")

    if (!authHeader || Authority(authHeader) !== "admin") {
        return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
    }

    const request: ElectionRequest = await e.request.json()

    let candidates: Candidate[] = []
    request.candidates.forEach((candidate, id) => {
        let newCandidate: Candidate = {
            candidate_id: id,
            first_name: candidate.first_name,
            last_name: candidate.last_name,
            party: candidate.party,
            photo: candidate.photo
        }
        candidates.push(newCandidate)
    })

    let newElection: Election = {
        election_id: electionCount,
        start: request.start,
        end: request.end,
        location: request.location,
        candidates: candidates
    }

    electionCount = electionCount + 1

    elections.push(newElection)

    await new Promise(resolve => setTimeout(resolve, 2500));
    return new Response("", { status: 201 });
};