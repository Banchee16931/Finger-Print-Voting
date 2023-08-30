export type CommonError = {
    message: string
    metadata: any | undefined,
}

export function NewError(message: string, metadata: any | undefined = undefined): CommonError {
    let error: CommonError = {
        message: message,
        metadata: metadata
    }

    return error
}

export const errorPrefix = "Error: "
export const unidentifiedErrorPrefix = "Unidentified Error: "