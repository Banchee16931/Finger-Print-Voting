import type { RequestHandler } from "@sveltejs/kit";
import { mockPath } from "../../mock/mockEndpoint";
import { BACKEND_ENDPOINT } from "$env/static/private";
import { NewError, type CommonError } from "$lib/types/CommonError";

export const GET: RequestHandler = async (e) => {
    console.log("GET ", e.request.url)
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

    return await e.fetch(`${BACKEND_ENDPOINT}/elections/user`, e.request)
}


