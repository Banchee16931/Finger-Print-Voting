export type UserData = {
    // Note: This just changes what webpages are displayed on the homescreen.
    // Election, user and other data is gotten via the backend which handle 
    // authentification.
    // This means no matter what this is set to user's won't be able to 
    // access anything they aren't supposed to
    level: null | "admin" | "user" // Note this just changes what webpages are displayed on the homescreen.
    username: string
    first_name: string
    last_name: string
}