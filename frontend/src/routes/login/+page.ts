import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

// forces the user out of the login page if they are already logged in
export const load = (async (e) => {
    let data = await e.parent()

    if (data.user.level !== null) {
        throw redirect(303, "/")
    }

    return {};
}) satisfies PageLoad;