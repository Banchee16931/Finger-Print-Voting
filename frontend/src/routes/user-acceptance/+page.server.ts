import type { Registrant } from '$lib/types';
import { NewError } from '$lib/types/CommonError';
import type { PageServerLoad } from './$types';

export const load = (async (e) => {
    console.log("loading reg")
    let res = await e.fetch("/api/registrations", { method: "GET" })
    if (res.ok) {
        let registrations: Registrant[] = await res.json()
        return {registrations: registrations};
    }

    throw NewError("no registrations")
}) satisfies PageServerLoad;