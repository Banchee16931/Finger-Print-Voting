import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../mock/mockEndpoint";
import type { CommonError } from '$lib/types/CommonError';

export const POST: RequestHandler = async (e) => {
    console.log("/register: ", e.request)
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/register`, e.request)
    console.log("response: ", res)
    if (!res.ok) {
        let err: CommonError = await res.json()
        throw err
    }

    return new Response("", { status: 201 })
};