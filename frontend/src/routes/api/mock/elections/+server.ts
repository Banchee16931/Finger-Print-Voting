import type { RequestHandler } from '@sveltejs/kit';
import type { Election, ElectionRequest, Candidate, CandidateVotes, ElectionState } from '$lib/types';
import { elections, electionState, votes } from "../data"
import { Authority } from '../handleAuth';
import { NewError } from '$lib/types/CommonError';

let electionCount = 5;

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
            party_colour: candidate.party_colour,
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

export const GET: RequestHandler = async (e) => {
    let returnElections = [...electionState]
    elections.forEach((election) => {
        let results = new Map<number, CandidateVotes>([])
        election.candidates.forEach((candidate) => {
            results.set(candidate.candidate_id, {
                first_name: candidate.first_name,
                last_name: candidate.last_name,
                party: candidate.party,
                party_colour: candidate.party_colour,
                votes: 0
            })
        })
        votes.forEach((vote) => {
            if (vote.election_id === election.election_id) {
                let foundCandidate = election.candidates.find((candidate) => {
                    if (candidate.candidate_id === vote.candidate_id) {
                        return true
                    }

                    return false
                })
                if (foundCandidate) {
                    let votedCandidate = results.get(vote.candidate_id)
                    if (votedCandidate !== undefined) {
                        votedCandidate.votes = votedCandidate.votes + 1
                        results.set(vote.candidate_id, votedCandidate)
                    } else {
                        results.set(vote.candidate_id, {
                            first_name: foundCandidate.first_name,
                            last_name: foundCandidate.last_name,
                            party: foundCandidate.party,
                            party_colour: foundCandidate.party_colour,
                            votes: 1
                        })
                    }
                }
            }
        })
        let newElection: ElectionState = {
            election_id: election.election_id,
            start: election.start,
            end: election.end,
            location: election.location,
            result: Array.from(results.values())
        }
        returnElections.push(newElection)
    })
    
    return new Response(JSON.stringify(returnElections), { status: 200 });
}