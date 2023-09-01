import { fail } from '@sveltejs/kit';
import type { Actions, RequestEvent } from '@sveltejs/kit';
import { validateLoginRequest, type LoginRequest } from "$lib/types/loginRequest";
import { NewError, type CommonError } from "$lib/types/CommonError"
import { redirect } from '@sveltejs/kit';
import { invalidate } from '$app/navigation';
import type { Candidate, CandidateRequest, Election, ElectionRequest } from '$lib/types';
import { start } from 'repl';

let electionsCreated = 0;

/** @type {import('./$types').Actions} */
export const actions: Actions = {
    newElection: async (e: RequestEvent) => {
        console.log("new election")
        let formData = await e.request.formData()

        let election: {
            start: string | null;
            end: string | null;
            location: string | null;
        } = {
            start: null,
            end: null,
            location: null,
        }
        let candidates: CandidateRequest[] | null = null;

        for (let field of formData) {
            const [key, value] = field
            console.log("handled key: ", key)
            console.log("value: ", value)
            switch (key) {
                case "start": {
                    if (typeof value === "string") {
                        console.log("start")
                        election.start = value
                    }
                    break;
                }
                case "end": {
                    if (typeof value === "string") {
                        console.log("end")
                        election.end = value
                    }
                    break;
                }
                case "location": {
                    if (typeof value === "string") {
                        console.log("location")
                        election.location = value
                    }
                    break;
                }
                case "candidates": {
                    if (typeof value === "string") {
                        let newCandidates: CandidateRequest[] = JSON.parse(value)
                        candidates = newCandidates
                    }
                    break;
                }
                default: {
                    throw NewError("couldn't process request")
                }
            }
        }

        if (election.start === null
            || election.end === null
            || election.location === null
            || candidates === null) {
            throw NewError("request is missing data")
        }

        let newElection: ElectionRequest = {
            start: election.start,
            end: election.end,
            location: election.location,
            candidates: candidates
        }

        let startDate = new Date(newElection.start.replace("-", ","))
        let endDate = new Date(newElection.end.replace("-", ","))
        if (startDate.getTime() > endDate.getTime()) {
            return fail(400, {
				error: "start date is after end date"
			});
        }

        let now = new Date()
        let today = new Date(now.getFullYear(), now.getMonth(), now.getDate())
        if (startDate.getTime() < today.getTime()) {
            return fail(400, {
				error: "date has already passed"
			});
        }

        if (newElection.candidates.length < 2) {
            return fail(400, {
				error: "less than two candidates"
			});
        }

        let failure: string | null = null
        let failureStatus: number = 500

        await e.fetch("/api/elections", {
            method: "POST",
            body: JSON.stringify(newElection)
        }).then(async (res) => {
            if (!res.ok) {
                let err: CommonError = await res.json()

                failureStatus = 400
                failure = err.message
            }
        }).catch((err: CommonError) => {
            failureStatus = 400
            failure = err.message
        })

        if (failure) {
            return fail(failureStatus, {
				error: failure
			});
        }

        electionsCreated = electionsCreated + 1

        return {
            electionsCreated: electionsCreated
        }
    },
};