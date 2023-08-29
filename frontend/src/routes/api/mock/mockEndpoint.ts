export function mockPath(path: string): string {
    console.log("mocking: ", path)
    return path.replace("api/", "api/mock/")
}