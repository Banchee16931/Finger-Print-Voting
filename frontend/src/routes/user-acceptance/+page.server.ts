import type { Registrant, UserAcceptanceRequest } from '$lib/types';
import { NewError, type CommonError } from '$lib/types/CommonError';
import type { PageServerLoad, RequestEvent } from './$types';
import { fail } from '@sveltejs/kit';
import type { Action } from '@sveltejs/kit'

let updateCounter = 0

export const load = (async (e) => {
    console.log("loading reg")
    let res = await e.fetch("/api/registrations", { method: "GET" })
    if (res.ok) {
        let registrations: Registrant[] = await res.json()
        return {registrations: registrations};
    }

    throw NewError("no registrations")
}) satisfies PageServerLoad;

/** @type {import('./$types').Actions} */
export const actions: Action = {
    accept: async (e: RequestEvent) => {
        console.log("Accept!")
        let formData = await e.request.formData()
        const data = new URLSearchParams()
        let newUserDetails: {
            id: string | null
            username: string | null
            password: string | null
            confirmPassword: string | null
            selected: Registrant | null
        } = {
            id: null,
            username: null,
            password: null,
            confirmPassword: null,
            selected: null
        };

        for (let field of formData) {
            const [key, value] = field
            data.append(key, value.toString())
        }

        

        data.forEach((value, key) => {
            switch (key) {
                case "id": {
                    newUserDetails.id = value
                    break;
                }
                case "username": {
                    newUserDetails.username = value
                    break;
                }
                case "password": {
                    newUserDetails.password = value
                    break;
                }
                case "confirm-password": {
                    newUserDetails.confirmPassword = value
                    break;
                }
                case "selected": {
                    newUserDetails.selected = JSON.parse(value)
                }
                default: {
                    return fail(500, {
                        error: "failed form submission"
                    })
                }
            }
        })

        if (newUserDetails.selected === null) {
            return fail(500, {
                error: "oh no"
            })
        }

        if (newUserDetails.id === null
                || newUserDetails.username === null 
                || newUserDetails.password === null
                || newUserDetails.confirmPassword === null) {
            return fail(500, {
                selectedUser: newUserDetails.selected,
                error: "failed form submission"
            })
        }

        if (newUserDetails.password !== newUserDetails.confirmPassword) {
            return fail(400, {
                selectedUser: newUserDetails.selected,
                error: "passwords don't match"
            })
        } 
        
        let acceptance: UserAcceptanceRequest = {
            registrant_id: parseInt(newUserDetails.id), 
            accepted: true,
            username: newUserDetails.username,
            password: newUserDetails.password
        }
        let res = await e.fetch("/api/registrations/acceptance", {
            method: "POST",
            body: JSON.stringify(acceptance)
        })

        if (!res.ok) {
            let err: CommonError = await res.json()
            return fail(res.status, 
                {
                    selectedUser: newUserDetails.selected,
                    error: err.message
                }
            )
        }

        updateCounter = updateCounter + 1
        return {
            selectedUser: newUserDetails.selected,pleaseUpdate: updateCounter}
    }
}
