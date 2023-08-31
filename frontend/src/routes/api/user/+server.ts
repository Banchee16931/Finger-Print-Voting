import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../mock/mockEndpoint";
import { NewError } from '$lib/types/CommonError';

export const GET: RequestHandler = async (e) => {
    console.log("/user")
    let authorization = e.cookies.get("session")
    if (authorization) {
        console.log("autherisation exists")
        e.request.headers.set("Autherization", authorization)
    } else {
        return new Response(JSON.stringify(NewError("invalid credentials")), { status: 401 })
    }
    
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/user`, e.request)
    console.log("user response: ", res)

    return res
};