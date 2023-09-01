
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../mock/mockEndpoint";
import { NewError, type CommonError } from '$lib/types/CommonError';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async (e) => {
    console.log("POST /elections")
    let authorization = e.cookies.get("session")
    if (authorization) {
        console.log("autherisation exists")
        e.request.headers.set("Autherization", authorization)
    } else {
        return new Response(JSON.stringify(NewError("invalid credentials")), { status: 401 })
    }

    console.log("POST /elections: ", e.request)
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/elections`, e.request)
    console.log("response: ", res)
    if (!res.ok) {
        let err: CommonError = await res.json()
        throw err
    }

    return new Response("", { status: 201 })
};

export const GET: RequestHandler = async (e) => {
    console.log("GET /elections: ", e.request)
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/elections`, e.request)
    console.log("response: ", res)
    if (!res.ok) {
        let err: CommonError = await res.json()
        throw err
    }

    return new Response("", { status: 201 })
};



