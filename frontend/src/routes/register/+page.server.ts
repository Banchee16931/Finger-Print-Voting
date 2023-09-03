import { fail } from '@sveltejs/kit';
import type { RegistrationRequest } from '$lib/types';
import { validateRegistrationRequest } from '$lib/types';
import type { Actions, RequestEvent } from '@sveltejs/kit';
import { NewError, errorPrefix, unidentifiedErrorPrefix } from '$lib/types/CommonError';
import type { CommonError } from '$lib/types/CommonError';

const genericErrorMessage = "failed to submit registration request"

/** @type {import('./$types').Actions} */
export const actions: Actions = {
    register: async (e: RequestEvent) => {
        console.log("action register")
        let formData = await e.request.formData()

        const data = new URLSearchParams()
        for (let field of formData) {
            const [key, value] = field

            console.log("key: ", key)

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
                    throw NewError(genericErrorMessage)
                }
            }
        })

        // checks all the different attributes are filled
        if (!validateRegistrationRequest(registrationRequest)) {
            console.log("data given: ", registrationRequest)
            throw NewError("missing data in registration")
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

        if (typeof failure === "string") {
            return fail(failureStatus, {
				error: failure,
                registered: false
			});
        }

        return { registered: true }
    }
}
