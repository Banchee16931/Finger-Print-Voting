import type { RequestHandler } from './$types';

export const DELETE: RequestHandler = async (e) => {
    console.log("deleting authorization key")
    e.cookies.delete("session", {path: "/"})

    return new Response("", { status: 200 });
};