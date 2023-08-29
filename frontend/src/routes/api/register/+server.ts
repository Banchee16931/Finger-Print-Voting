import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../mock/mockEndpoint";

export const POST: RequestHandler = async (e) => {
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    return await e.fetch(`${BACKEND_ENDPOINT}/register`, e.request)
};