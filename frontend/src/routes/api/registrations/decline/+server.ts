import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../../mock/mockEndpoint";

export const POST: RequestHandler = async (e) => {
    console.log("/registrations/accept")
    let authorization = e.cookies.get("session")
    if (authorization) {
        console.log("autherisation exists")
        e.request.headers.set("Autherization", authorization)
    } else {
        return new Response("", { status: 401 })
    }
    
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/registrations`, e.request)
    console.log("registrations response: ", res)

    return res
};