import type { RequestHandler } from './$types';
import type { UserDetails } from "$lib/types"
import { NewError } from '$lib/types/CommonError';
import { StripBearer } from '../handleAuth';
import { getUser, users } from "../data"

export const GET: RequestHandler = async (e) => {
    
    let req = e.request

    let authHeader = StripBearer(req.headers.get("Autherization"))

    if (!authHeader) {
        return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
    }

    let user = getUser(authHeader)
    if (!user) {
        return new Response(JSON.stringify(NewError("user doesn't exist")), {status: 404});
    }
    return new Response(JSON.stringify(user), {status: 200});
};
