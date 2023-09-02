import { error } from "@sveltejs/kit"
import { NewError } from "$lib/types/CommonError"
import type { VoteRequest } from '$lib/types';
import type { RequestHandler } from './$types';
import { StripBearer } from '../handleAuth';
import { getUser, registrations, users, voters, votes } from "../data";

export const POST: RequestHandler = async (e) => {
    let req = e.request

    console.log("mocking vote")

    let authHeader = StripBearer(req.headers.get("Autherization"))

    if (!authHeader) {
        return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
    }

    let user = getUser(authHeader)
    if (user === null) {
        return new Response(JSON.stringify(NewError("user doesn't exist")), {status: 404});
    }

    let voter = voters.find((voter) => {
        if (voter.username === user?.username) {
            return true
        }

        return false
    })
    if (!voter) {
        return new Response(JSON.stringify(NewError("user is not a voter")), {status: 404});
    }

    const request: VoteRequest = await e.request.json()

    if (votes.find((vote) => {
        if (vote.election_id == request.election_id && vote.username == user?.username) {
            return true
        }
        return false
    })) {
        return new Response(JSON.stringify(NewError("already voted")), {status: 400});
    }

    if (request.fingerprint !== voter?.fingerprint) {
        return new Response(JSON.stringify(NewError("invalid fingerprint")), {status: 401});
    }

    votes.push({
        username: user.username,
        election_id: request.election_id,
        candidate_id: request.candidate_id
    })

    voters.splice(voters.findIndex((findVoter) => {
        if (user?.username == findVoter.username) {
            return true
        }

        return false
    }), 1)

    users.splice(voters.findIndex((findUser) => {
        if (user?.username == findUser.username) {
            return true
        }

        return false
    }), 1)

    await new Promise(resolve => setTimeout(resolve, 2500));
    return new Response("", { status: 201 });
};