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

    let path = e.request.url.substring(e.request.url.indexOf("api/"))

    console.log("new path: ", path)
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    let res = await e.fetch(`${BACKEND_ENDPOINT}/${path}`, e.request)
    console.log("response: ", res)
    if (!res.ok) {
        let err: CommonError = await res.json()
        throw err
    }

    return new Response("", { status: 201 })
};



