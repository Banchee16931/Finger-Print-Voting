import { fail } from '@sveltejs/kit';
import type { Actions, RequestEvent } from '@sveltejs/kit';
import { validateLoginRequest, type LoginRequest } from "$lib/types/loginRequest";
import type { CommonError } from "$lib/types/CommonError"
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
                    return fail(500, {
                        error: "missing required data"
                    });
                }
            }
        })

        if (!validateLoginRequest(loginRequest)) {
            return fail(500, {
				error: "missing required data"
			});
        }

        let res = await e.fetch("/api/login", {
            method: "POST",
            body: JSON.stringify(loginRequest)
        });

        if (res && !res.ok) {
            if (res.status === 401) {
                return fail(401, {
                    error: "invalid credentials"
                });
            }

            let err: CommonError = await res.json()

            return fail(res.status, {
				error: err.message
			});
        }

        return {thing: true}
    },
};