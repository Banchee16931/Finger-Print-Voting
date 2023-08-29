import { error } from "@sveltejs/kit"
import { NewError } from "$lib/types/CommonError"
import { validateRegistrationRequest, type RegistrationRequest } from '$lib/types';
import type { RequestHandler } from './$types';

let registeredUsers: RegistrationRequest[] = []

export const POST: RequestHandler = async (e) => {
    const request: RegistrationRequest = await e.request.json()
    if (!validateRegistrationRequest(request) || request?.first_name == "fail") {
        throw error(400,  NewError('invalid registration data'));
    }

    registeredUsers.push(request)

    await new Promise(resolve => setTimeout(resolve, 2500));
    return new Response("", { status: 201 });
};