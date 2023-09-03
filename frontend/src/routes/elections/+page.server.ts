import type { ElectionState } from "$lib/types";
import { NewError, type CommonError } from "$lib/types/CommonError";
import type { PageServerLoad } from "./$types";


export const load = (async (e) => {
    console.log("loading reg")
    let res = await e.fetch("/api/elections", { method: "GET" })
    if (!res.ok) {
        if (res.status == 404) {
            return {elections: []};
        }
        throw NewError("failed to load elections", res)
    }

    let elections: ElectionState[] = await res.json()
    return {elections: elections};
}) satisfies PageServerLoad;