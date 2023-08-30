import { NewError } from '$lib/types/CommonError';

/** @type {import('@sveltejs/kit').HandleClientError} */
export async function handleError(error: any) {
	console.log("handled client error: ", error)

	return {
		message: "Unknown Error",
        metadata: undefined
	};
}