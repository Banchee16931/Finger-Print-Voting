import { NewError } from '$lib/types/CommonError';

/** @type {import('@sveltejs/kit').HandleServerError} */
export async function handleError(error: any) {
	console.log("handled server error: ", error)

	return {
		message: "Unknown Error",
        metadata: undefined
	};
}