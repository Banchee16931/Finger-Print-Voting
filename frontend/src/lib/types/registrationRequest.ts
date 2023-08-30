export type RegistrationRequest = {
    first_name: string | undefined,
    last_name: string | undefined,
    email: string | undefined,
    phone_no: string | undefined,
    proof_of_identification: string | undefined,
    fingerprint: string | undefined,
}

export const validateRegistrationRequest = function(req: RegistrationRequest) {
    if (req.email === null 
    || req.first_name === null
    || req.last_name === null
    || req.email === null
    || req.phone_no === null
    || req.proof_of_identification === null
    || req.fingerprint === null) {
        return false
    }

    return true
}