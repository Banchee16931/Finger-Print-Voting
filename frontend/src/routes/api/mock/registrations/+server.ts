import type { RequestHandler } from './$types';
import { validateRegistrationRequest, type UserDetails, type RegistrationRequest, type Registrant } from "$lib/types"
import { NewError } from '$lib/types/CommonError';
import { Authority } from '../handleAuth';
import { registrations } from '../data';


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
        return new Response(JSON.stringify(NewError("invalid registration data")), {status: 400});
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

export const GET: RequestHandler = async (e) => {
    let req = e.request

    let authHeader = req.headers.get("Autherization")

    if (!authHeader || Authority(authHeader) !== "admin") {
        return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
    }

    return new Response(JSON.stringify(registrations), {status: 200});
};