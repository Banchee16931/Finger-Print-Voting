// See https://kit.svelte.dev/docs/types#app

import type { UserData } from "$lib/auth";
import type { NewError } from "$lib/types/CommonError";

// for information about these interfaces
declare global {
	namespace App {
		interface Error {
			message: string
    		metadata: any | undefined,
		}
		interface Locals {
			user: UserData
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};
