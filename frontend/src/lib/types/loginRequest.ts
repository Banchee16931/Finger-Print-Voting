export type LoginRequest = {
    username: string | null,
    password: string | null,
}

export const validateLoginRequest = function(req: LoginRequest) {
    if (req.username === null 
    || req.password === null) {
        return false
    }

    return true
}