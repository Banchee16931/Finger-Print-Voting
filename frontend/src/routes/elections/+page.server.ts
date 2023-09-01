import type { ElectionState } from "$lib/types";
import type { CommonError } from "$lib/types/CommonError";
import type { PageServerLoad } from "./$types";


export const load = (async (e) => {
    console.log("loading reg")
    let res = await e.fetch("/api/elections", { method: "GET" })
    if (!res.ok) {
        let elections: CommonError = await res.json()
        throw elections
    }

    let elections: ElectionState[] = await res.json()
    return {elections: elections};
}) satisfies PageServerLoad;