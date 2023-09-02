import type { Registrant, UserAcceptanceRequest, UserStore, Voter } from "$lib/types"
import { NewError } from '$lib/types/CommonError';
import { Authority } from '../../handleAuth';
import { registrations, users, voters } from '../../data';
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

        let voter: Voter = {
            username: userAcceptance.username,
            password: userAcceptance.password,
            first_name: reg.first_name,
            last_name: reg.last_name,
            phone_no: reg.phone_no,
            email: reg.email,
            fingerprint: reg.fingerprint,
            location: reg.location
        }
        users.push(newUser)
        voters.push(voter)
    }

    registrations.splice(removeIndex, 1)
    
    return new Response(JSON.stringify(registrations), {status: 200});
};