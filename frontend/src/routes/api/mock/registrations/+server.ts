import type { RequestHandler } from './$types';
import type { UserDetails } from "$lib/types"
import { NewError } from '$lib/types/CommonError';
import { Authority } from '../handleAuth';
import { registrations } from '../data';

export const GET: RequestHandler = async (e) => {
    let req = e.request

    let authHeader = req.headers.get("Autherization")

    if (authHeader) {
        console.log("auth header", authHeader)

        if (Authority(authHeader) === "admin") {
            return new Response(JSON.stringify(registrations), {status: 200});
        }
    }

    return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
};