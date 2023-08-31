import type { Registrant, UserAcceptanceRequest, UserStore } from "$lib/types"
import { NewError } from '$lib/types/CommonError';
import { Authority } from '../../handleAuth';
import { registrations, users } from '../../data';
import type { RequestHandler } from "@sveltejs/kit";

export const POST: RequestHandler = async (e) => {
    console.log("acceptance")
    let req = e.request

    let authHeader = req.headers.get("Autherization")

    if (!authHeader || Authority(authHeader) !== "admin") {
        return new Response(JSON.stringify(NewError("invalid credentials")), {status: 401});
    }

    let userAcceptance: UserAcceptanceRequest = await req.json()

    let removeIndex = registrations.findIndex((value) => {
        if (value.registrant_id == userAcceptance.registrant_id) {
            return true
        }
        return false
    })

    if (userAcceptance.accepted) {
        if (!(userAcceptance.username && userAcceptance.password)) {
            return new Response(JSON.stringify(NewError("missing user data")), {status: 400});
        }
        
        let reg = registrations[removeIndex]
        let newUser: UserStore = {
            admin: false,
            username: userAcceptance.username,
            password: userAcceptance.password,
            first_name: reg.first_name,
            last_name: reg.last_name,
        }
        users.push(newUser)
    }

    registrations.splice(removeIndex, 1)
    
    return new Response(JSON.stringify(registrations), {status: 200});
};