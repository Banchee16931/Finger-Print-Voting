import type { Registrant } from '$lib/types';
import { NewError } from '$lib/types/CommonError';
import type { PageServerLoad } from './$types';

export const load = (async (e) => {
    let res = await e.fetch("/api/registrations", { method: "GET" })
    if (res.ok) {
        console.log("ok")
        let registrations: Registrant[] = await res.json()
        console.log("registrations: ", registrations)
        return {registrations: registrations};
    }
    
    throw NewError("no registrations")
}) satisfies PageServerLoad;