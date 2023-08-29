import type { LoginRequest } from '$lib/types';
import type { RequestHandler } from './$types';
import { NewError } from '$lib/types/CommonError';

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

    return new Response(JSON.stringify(NewError("invalid credentials long on purpose hwhehwheheheheheheheheheheh")), { status: 401});
};