export function Authority(key: string): "admin" | "user" | undefined {
    if (key.startsWith("Bearer {"), key.endsWith("}")) {
        let authorizationID = key.replace("Bearer {", "").replace("}", "")
        if (authorizationID !== "admin" && authorizationID !== "user") {
            return undefined
        }
        return authorizationID
    }
}