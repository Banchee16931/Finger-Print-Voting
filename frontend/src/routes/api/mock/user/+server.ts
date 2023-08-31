import type { RequestHandler } from './$types';
import type { UserDetails } from "$lib/types"
import { NewError } from '$lib/types/CommonError';
import { StripBearer } from '../handleAuth';
import { users } from "../data"

export const GET: RequestHandler = async (e) => {
    
    let req = e.request

    let authHeader = StripBearer(req.headers.get("Autherization"))

    if (authHeader) {
        console.log("auth header", authHeader)
        if (authHeader == "admin") {
            console.log("admin")
            let data: UserDetails = {
                admin: true,
                username: "admin",
                first_name: "Aron",
                last_name: "Access"
            }
            return new Response(JSON.stringify(data), { status: 200 });
        } else if (authHeader == "user") {
            console.log("user")
            let data: UserDetails = {
                admin: false,
                username: "user",
                first_name: "Norman",
                last_name: "Normal"
            }
            return new Response(JSON.stringify(data), { status: 200 });
        }

        let foundUser = users.find((value) => {
            if (value.username === authHeader) {
                return true
            }
            return false
        })

        if (foundUser) {
            console.log("found user")
            let returnUser: UserDetails = {
                admin: foundUser.admin,
                username: foundUser.username,
                first_name: foundUser.first_name,
                last_name: foundUser.last_name
            }

            return new Response(JSON.stringify(returnUser), { status: 200 });
        }
    }

    return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
};