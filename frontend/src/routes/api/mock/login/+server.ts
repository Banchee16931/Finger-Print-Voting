import type { LoginRequest } from '$lib/types';
import type { RequestHandler } from './$types';
import { NewError } from '$lib/types/CommonError';
import { users } from '../data';

export const POST: RequestHandler = async (e) => {
    const request: LoginRequest = await e.request.json()

    await new Promise(resolve => setTimeout(resolve, 500));

    if (request.username == "admin") {
        e.cookies.set("session", "Bearer {admin}", { path: "/" })
        return new Response("", { status: 200 });
    } else if (request.username == "user") {
        e.cookies.set("session", "Bearer {user}", { path: "/" })
        return new Response("", { status: 200 });
    }

    let foundIndex = users.findIndex((value) => {
        if (value.username === request.username && value.password === request.password) {
            return true
        }
        return false
    })

    if (foundIndex === -1) {
        return new Response(JSON.stringify(NewError("invalid credentials")), { status: 401});
    }

    e.cookies.set("session", `Bearer {${request.username}}`, { path: "/" })
    console.log("Setting auth cookie to: ", request.username)
    return new Response("", { status: 200 });
};