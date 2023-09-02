import type { RequestHandler } from '@sveltejs/kit';
import type { Election, ElectionRequest, Candidate, CandidateVotes, ElectionState } from '$lib/types';
import { elections, electionState, votes, voters } from "../../data"
import { Authority } from '../../handleAuth';
import { NewError } from '$lib/types/CommonError';

export const GET: RequestHandler = async (e) => {
    console.log("Mocked GET /api/elections/[username]")
    let req = e.request

    let username = e.params.username
    if (!username) {
        return new Response(JSON.stringify(NewError("missing username param")), {status: 500});
    }

    let authHeader = req.headers.get("Autherization")
    console.log("auth header: ", authHeader)

    if (!authHeader || Authority(authHeader) === undefined) {
        return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
    }

    let voter = voters.find((voter) => {
        if (voter.username === username) {
            return true
        }

        return false
    })

    if (!voter) {
        return new Response(JSON.stringify(NewError("unregistered voter", {noUserAccount: true})), {status: 400});
    }

    elections.sort((prev, current) => {
        return new Date(current.start.replace("-", ",")).getTime() - new Date(prev.start.replace("-", ",")).getTime()
    }).filter((election) => {
        if (election.location == voter?.location) {
            return true
        }

        return false
    })

    if (elections.length === 0) {
        return new Response(JSON.stringify(NewError("no running elections for this user")), {status: 404});
    }

    let election = elections[0]

    return new Response(JSON.stringify(election), { status: 200 });
}