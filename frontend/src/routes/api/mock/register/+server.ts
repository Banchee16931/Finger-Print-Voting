import { error } from "@sveltejs/kit"
import { NewError } from "$lib/types/CommonError"
import { validateRegistrationRequest, type RegistrationRequest, type Registrant } from '$lib/types';
import type { RequestHandler } from './$types';
import { registrations } from "../data";

let registeredUsersCounter: number = 2

export const POST: RequestHandler = async (e) => {
    const request: RegistrationRequest = await e.request.json()
    if (!validateRegistrationRequest(request) || request?.first_name == "fail"
    || request.email === undefined 
    || request.first_name === undefined
    || request.last_name === undefined
    || request.email === undefined
    || request.phone_no === undefined
    || request.proof_of_identity === undefined
    || request.fingerprint === undefined
    || request.location === undefined) {
        throw error(400,  NewError('invalid registration data'));
    }

    let newReg: Registrant = {
        registrant_id: registeredUsersCounter,
        first_name: request.first_name,
        last_name: request.last_name,
        email: request.email,
        phone_no: request.phone_no,
        fingerprint: request.fingerprint,
        proof_of_identity: request.proof_of_identity,
        location: request.location
    }

    registeredUsersCounter = registeredUsersCounter + 1

    registrations.push(newReg)

    await new Promise(resolve => setTimeout(resolve, 2500));
    return new Response("", { status: 201 });
};