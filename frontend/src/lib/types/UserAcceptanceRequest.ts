import type { Voter } from "./Voter"

export type UserAcceptanceRequest = {
	registrant_id: number   
    accepted: boolean
    username: string | undefined
    password: string | undefined
}