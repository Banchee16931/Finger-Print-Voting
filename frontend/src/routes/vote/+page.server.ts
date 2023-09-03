import type { Election } from '$lib/types';
import { NewError, type CommonError } from '$lib/types/CommonError';
import { fail, type Actions, type RequestEvent } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async (e) => {
    let parent = await e.parent()

    if (parent.user.level === null) {
        return {election: null}
    }

    console.log("username: ", parent.user.username)
    console.log("path: ", `/api/elections/${parent.user.username}`)
    let response = await e.fetch(`/api/elections/${parent.user.username}`, { method: "GET" })
    let election: Election | null = null
    if (!response.ok) {
        let err: CommonError = await response.json()

        if (err.metadata.noUserAccount !== true) {
            throw NewError(`getting current election failed: ${response.statusText}, ${await response.text()}`)
        }
    } else {
        election = await response.json()
    }

    console.log("election: ", election)

    return {
        election: election,
    };
}) satisfies PageServerLoad;

/** @type {import('./$types').Actions} */
export const actions: Actions = {
    vote: async (e: RequestEvent) => {
        console.log("action: vote")
        let formData = await e.request.formData()

        let voteRequest: {
            election_id: number | undefined
            fingerprint: string | undefined
            candidate_id: number | undefined
        } = {
            fingerprint:  undefined,
            candidate_id: undefined,
            election_id: undefined
        };


        for (let field of formData) {
            const [key, value] = field

            console.log("key: ", key)
            console.log("value type: ", typeof value)

            switch (key) {
                case "fingerprint": {
                    if (typeof value === "string") {
                        voteRequest.fingerprint = value
                    }
                    break;
                }
                case "election_id": {
                    if (typeof value === "string") {
                        voteRequest.election_id = parseInt(value)
                    }
                    break;
                }
                case "candidate_id": {
                    if (typeof value === "string") {
                        voteRequest.candidate_id = parseInt(value)
                    }
                    break;
                }
                default: {
                    return fail(500, {
                        error: "form contained unexpected data: ["+key+"]",
                        registered: false
                    });
                }
            }
        }

        // checks all the different attributes are filled
        if ((voteRequest.fingerprint === undefined)
                || (voteRequest.candidate_id === undefined)
                || (voteRequest.election_id === undefined)) {
            return fail(500, {
                error: "missing data in vote",
                registered: false
            });
        }

        let failure: string | null = null
        let failureStatus: number = 500

        let res = await e.fetch("/api/votes", { 
            method:"POST", 
            body: JSON.stringify(voteRequest), 
            headers: { 'content-type': 'application/json'} ,
            signal: AbortSignal.timeout(3000),
        }).catch((err: CommonError) => {
            console.log("caught common error")
            failureStatus = 400
            failure = err.message
        })

        if (!res) {
            return fail(failureStatus, {
				error: "did not get a response",
                registered: false
			});
        }

        if (!res.ok) {
            console.log("caught res not ok")
            let err: CommonError = await res.json()

            failureStatus = 400
            failure = err.message
        }

        if (typeof failure === "string") {
            return fail(failureStatus, {
				error: failure,
                registered: false
			});
        }

        return { registered: true }
    }
}
