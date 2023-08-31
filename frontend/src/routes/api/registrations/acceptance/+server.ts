import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../../mock/mockEndpoint";
import { invalidate, invalidateAll } from '$app/navigation';

export const POST: RequestHandler = async (e) => {
    console.log("/registrations/acceptance")
    let authorization = e.cookies.get("session")
    if (authorization) {
        console.log("autherisation exists")
        e.request.headers.set("Autherization", authorization)
    } else {
        return new Response("", { status: 401 })
    }
    
    
    let res:Response
    if (import.meta.env.DEV) {
        res =  await e.fetch(mockPath(e.request.url), e.request)
    } else {
        res = await e.fetch(`${BACKEND_ENDPOINT}/registrations/acceptance`, e.request)
        console.log("registrations response: ", res)
    }

    return res
};