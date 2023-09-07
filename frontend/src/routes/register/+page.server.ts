import { fail } from '@sveltejs/kit';
import type { RegistrationRequest } from '$lib/types';
import { validateRegistrationRequest } from '$lib/types';
import type { Actions, RequestEvent } from '@sveltejs/kit';
import { errorPrefix } from '$lib/types/CommonError';
import type { CommonError } from '$lib/types/CommonError';

/** @type {import('./$types').Actions} */
export const actions: Actions = {
    register: async (e: RequestEvent) => {
        console.log("action register")
        let formData = await e.request.formData()

        const data = new URLSearchParams()
        for (let field of formData) {
            const [key, value] = field

            console.log("key: ", key)
            console.log("value: ", value)

            if (value === undefined) {
                console.log(errorPrefix, `${key} is not set`)
                return fail(400, {
                    error: `${key} is not set`,
                    registered: false
                });
            }

            if (typeof value === "string") {
                data.append(key, value)
            }
        }

        let registrationRequest: RegistrationRequest = {
            first_name:  undefined,
            last_name:  undefined,
            email:  undefined,
            phone_no:  undefined,
            proof_of_identity:  undefined,
            fingerprint:  undefined,
            location: undefined
        };

        // fills in each element of the request based on the form data
        data.forEach((value, key) => {
            switch (key) {
                case "firstname": {
                    registrationRequest.first_name = value
                    break;
                }
                case "surname": {
                    registrationRequest.last_name = value
                    break;
                }
                case "email": {
                    registrationRequest.email = value
                    break;
                }
                case "telephone": {
                    registrationRequest.phone_no = value
                    break;
                }
                case "identification": {
                    console.log("set proof of identity")
                    registrationRequest.proof_of_identity = value
                    break;
                }
                case "fingerprint": {
                    registrationRequest.fingerprint = value
                    break;
                }
                case "location": {
                    registrationRequest.location = value
                    break;
                }
                default: {
                    console.log(errorPrefix, "form contained unexpected data")
                    return fail(400, {
                        error: "missing details",
                        registered: false
                    });
                }
            }
        })

        // checks all the different attributes are filled
        let result = validateRegistrationRequest(registrationRequest)
        if (result !== undefined) {
            console.log("data given: ", registrationRequest)
            return fail(400, {
				error: result.message,
                registered: false
			});
        }

        let failure: string | null = null
        let failureStatus: number = 500

        let res = await e.fetch("/api/registrations", { 
            method:"POST", 
            body: JSON.stringify(registrationRequest), 
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

        if (failure !== null) {
            return fail(failureStatus, {
				error: failure,
                registered: false
			});
        }

        console.log("end")

        return { registered: true }
    }
}
