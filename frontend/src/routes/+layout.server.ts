import type { LayoutServerLoad } from './$types';
import { user } from "$lib/auth";
import type { UserData, UserDetails } from '$lib/types';
import { redirect, type ServerLoadEvent } from '@sveltejs/kit';

let oldUser: UserData | null = null

export const load = (async (e: ServerLoadEvent) => {
    let authorization = e.cookies.get("session")
    user.set({
        level: null,
        username: "",
        first_name: "",
        last_name: "",
    })

    let userRes: UserDetails | null = null
    if (authorization) {
        console.log("session changes")
        let res = await e.fetch("/api/user", { method: "GET" })
        if (res.ok) {
            console.log("ok")
            userRes = await res.json()
            console.log("user res: ", userRes)
        } else {
            console.log("not ok")
        }
    }

    let userData: UserData | null = null

    if (userRes) {
        console.log("is admin: ", userRes.admin)
        userData = {
            level: userRes.admin ? "admin" : "user",
            username: userRes.username,
            first_name: userRes.first_name,
            last_name: userRes.last_name
        }
    }

    if (userData === null) {
        userData = {
            level: null,
            last_name: "",
            first_name: "",
            username: "",
        }
    }

    if (oldUser?.level !== null && userData.level == null) {
        oldUser = userData
        throw redirect(307, "/")
    }

    oldUser = userData

    return {
        user: userData
    };
}) satisfies LayoutServerLoad;