import type { RequestHandler } from '@sveltejs/kit';
import type { Election, ElectionRequest, Candidate, CandidateVotes, ElectionState } from '$lib/types';
import { elections, electionState, votes, voters, getUser } from "../../data"
import { Authority, StripBearer } from '../../handleAuth';
import { NewError } from '$lib/types/CommonError';

export const GET: RequestHandler = async (e) => {
    console.log("Mocked GET /api/elections/user")
    let req = e.request

    let authHeader = StripBearer(req.headers.get("Autherization"))

    if (!authHeader) {
        return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
    }

    let user = getUser(authHeader)
    if (!user) {
        return new Response(JSON.stringify(NewError("user doesn't exist")), {status: 404});
    }

    let voter = voters.find((voter) => {
        if (voter.username === user?.username) {
            return true
        }

        return false
    })

    if (!voter) {
        return new Response(JSON.stringify(NewError("unregistered voter", {noUserAccount: true})), {status: 400});
    }

    console.log("voter: ", voter)

    let filteredElections = elections.sort((prev, current) => {
        return new Date(current.start.replace("-", ",")).getTime() - new Date(prev.start.replace("-", ",")).getTime()
    }).filter((election) => {
        if (election.location == voter?.location) {
            return true
        }

        return false
    })

    if (filteredElections.length === 0) {
        return new Response(JSON.stringify(NewError("no running elections for this user")), {status: 404});
    }

    let election = filteredElections[0]

    return new Response(JSON.stringify(election), { status: 200 });
}