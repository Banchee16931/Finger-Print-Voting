import { fail } from '@sveltejs/kit';
import type { Actions, RequestEvent } from '@sveltejs/kit';
import { validateLoginRequest, type LoginRequest } from "$lib/types/loginRequest";
import { NewError, type CommonError } from "$lib/types/CommonError"
import { redirect } from '@sveltejs/kit';
import { invalidate } from '$app/navigation';

/** @type {import('./$types').Actions} */
export const actions: Actions = {
    login: async (e: RequestEvent) => {     
        let formData = await e.request.formData()

        const data = new URLSearchParams()
        for (let field of formData) {
            const [key, value] = field
            data.append(key, value.toString())
        }

        let loginRequest: LoginRequest = {
            username: null,
            password: null
        };

        data.forEach((value, key) => {
            switch (key) {
                case "username": {
                    loginRequest.username = value
                    break;
                }
                case "password": {
                    loginRequest.password = value
                    break;
                }
                default: {
                    throw NewError("couldn't process request")
                }
            }
        })

        if (!validateLoginRequest(loginRequest)) {
            throw NewError("request is missing data")
        }

        let failure: string | null = null
        let failureStatus: number = 500

        await e.fetch("/api/login", {
            method: "POST",
            body: JSON.stringify(loginRequest)
        }).then(async (res) => {
            if (!res.ok) {
                console.log("failing")
                let err: CommonError = await res.json()

                failureStatus = 400
                failure = err.message
            }
        }).catch((err: CommonError) => {
            console.log("failing 2")
            failureStatus = 400
            failure = err.message
        })

        if (failure) {
            console.log("reported failure")
            return fail(failureStatus, {
				error: failure
			});
        }
    },
};