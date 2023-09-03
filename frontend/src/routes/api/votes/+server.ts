import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../mock/mockEndpoint";
import { NewError, type CommonError } from '$lib/types/CommonError';

export const POST: RequestHandler = async (e) => {
    let authorization = e.cookies.get("session")
    if (authorization) {
        console.log("autherisation exists")
        e.request.headers.set("Autherization", authorization)
    } else {
        return new Response(JSON.stringify(NewError("invalid credentials")), { status: 401 })
    }

    console.log("/votes: ", e.request)
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/votes`, e.request)
    console.log("response: ", res)
    if (!res.ok) {
        let err: CommonError = await res.json()
        throw err
    }

    return new Response("", { status: 201 })
};