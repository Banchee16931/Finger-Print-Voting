import type { RequestHandler } from './$types';
import type { UserDetails } from "$lib/types"
import { NewError } from '$lib/types/CommonError';
import { Authority } from '../handleAuth';

export const GET: RequestHandler = async (e) => {
    
    let req = e.request

    let authHeader = req.headers.get("Autherization")

    if (authHeader) {
        console.log("auth header", authHeader)
        let authorizationID = Authority(authHeader)
        if (authorizationID == "admin") {
            console.log("admin")
            let data: UserDetails = {
                admin: true,
                username: "admin",
                first_name: "Aron",
                last_name: "Access"
            }
            return new Response(JSON.stringify(data), { status: 200 });
        } else if (authorizationID == "user") {
            console.log("user")
            let data: UserDetails = {
                admin: false,
                username: "user",
                first_name: "Norman",
                last_name: "Normal"
            }
            return new Response(JSON.stringify(data), { status: 200 });
        }
    }

    return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
};