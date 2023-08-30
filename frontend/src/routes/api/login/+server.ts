import type { RequestHandler } from './$types';
import { BACKEND_ENDPOINT } from '$env/static/private';
import { mockPath } from "../mock/mockEndpoint";

export const POST: RequestHandler = async (e) => {
    if (import.meta.env.DEV) {
        return await e.fetch(mockPath(e.request.url), e.request)
    }

    console.log("fetching login stuff: ", BACKEND_ENDPOINT)
    let res = await e.fetch(`${BACKEND_ENDPOINT}/login`, e.request)
    console.log("response: ", res)
    let newCookies = res.headers.get("Set-Cookie")
    if (newCookies) {
        newCookies.split(",").forEach((cookie) => {
                const [key, value] = cookie.split(":")
                console.log("cookie: ", key, " : ", value)
                if (key === "Autherization") {
                    console.log("setting session")
                    e.cookies.set("session", value.trim(), {path:"/"})
                }
            }
        )
    }

    console.log("after cookies")
    
    return new Response("", { status: 200});
}