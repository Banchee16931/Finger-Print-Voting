export function Authority(key: string): "admin" | "user" | undefined {
    if (key.startsWith("Bearer {"), key.endsWith("}")) {
        let authorizationID = key.replace("Bearer {", "").replace("}", "")
        if (authorizationID === "admin") {
            return "admin"
        }
        return "user"
    }
    return undefined
}

export function StripBearer(key: string | undefined | null): string | undefined {
    if (key) {
        if (key.startsWith("Bearer {"), key.endsWith("}")) {
            let authorizationID = key.replace("Bearer {", "").replace("}", "")
            return authorizationID
        } 
    }
    
    return undefined
}