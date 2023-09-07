import { NewError } from "./CommonError"

export type RegistrationRequest = {
    first_name: string | undefined,
    last_name: string | undefined,
    email: string | undefined,
    phone_no: string | undefined,
    proof_of_identity: string | undefined,
    fingerprint: string | undefined,
    location: string | undefined
}

export const validateRegistrationRequest = function(req: RegistrationRequest) {
    console.log("request: ", req)
    if (req.email === undefined 
    || req.first_name === undefined
    || req.last_name === undefined
    || req.email === undefined
    || req.phone_no === undefined
    || req.proof_of_identity === undefined
    || req.fingerprint === undefined
    || req.location === undefined) {
        return NewError("missing data")
    }

    if (req.proof_of_identity === "") {
        return NewError("missing identity of proof")
    }

    if (req.fingerprint === "") {
        return NewError("missing fingerprint")
    }
}