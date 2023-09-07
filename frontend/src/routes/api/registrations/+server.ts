import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../mock/mockEndpoint";
import { NewError, type CommonError } from '$lib/types/CommonError';

export const POST: RequestHandler = async (e) => {
    console.log("/registrations: ", e.request)
    if (import.meta.env.PROD) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/registrations`, e.request)
    console.log("response: ", res)
    if (!res.ok) {
        let err: CommonError = await res.json()
        throw err
    }

    return new Response("", { status: 201 })
};

export const GET: RequestHandler = async (e) => {
    console.log("/registrations")
    let authorization = e.cookies.get("session")
    if (authorization) {
        console.log("autherisation exists")
        e.request.headers.set("Autherization", authorization)
    } else {
        return new Response(JSON.stringify(NewError("invalid credentials")), { status: 401 })
    }
    
    if (import.meta.env.PROD) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/registrations`, e.request)
    console.log("registrations response: ", res)

    return res
};