<script lang="ts">
    import Hero from "../../lib/components/Hero.svelte"
    import TextInput from "../../lib/components/TextInput.svelte"
	import FileInput from '$lib/components/FileInput.svelte';
	import TickAnimation from '$lib/components/TickAnimation.svelte';
	import type { RegistrationRequest } from '$lib/types/registrationRequest';
    import { validateRegistrationRequest } from '$lib/types/registrationRequest';
	import { NewError, errorPrefix, unidentifiedErrorPrefix } from '$lib/types/CommonError';
    import type { CommonError } from '$lib/types/CommonError';

    // The file given by the user that stores the proof of identification
    let identificationFiles: FileList| null = null
    // The file given by the user that stores their fingerprint
    let fingerprintFiles: FileList| null = null

    // If the registered screen/animation should appear
    let isRegistered: boolean = false;

    // Stops the user interacting with the elements as the form is being processed
    let working: boolean = false

    // Shown as red text at bottom of form to alert the user of any issues that have occured with their request
    let submissionError: string | null = null

    // Message that reports 
    const genericErrorMessage = "failed to submit registration request"

    // Submits the form's data to the backend for processing
    function submit(e: SubmitEvent) {
        submissionError = null // resetting the submission error text

        // lets the submit be handled asynchronously
        const submitAsync = async function() {
            if (e.target === null) { // if submit triggered without event data
                console.log(errorPrefix, "submit triggered without event data")
                throw NewError(genericErrorMessage)
            }

            // converting form data into a processable input
            const formData = new FormData(e.target as HTMLFormElement)
            const data = new URLSearchParams()
            for (let field of formData) {
                const [key, value] = field
                
                if (value instanceof File) { // if data is a file
                    // converts image files into smaller versions so they can be transfered to the backend
                    await shrinkImage(value)
                    .then((encodedImage) => {
                        data.append(key, encodedImage)
                    })
                    .catch((err: CommonError) => {
                        console.log(errorPrefix, "when shrinking image: ", err.message)
                        throw NewError(genericErrorMessage)
                    })
                    .catch((err: any) => {
                        console.log(unidentifiedErrorPrefix, "when shrinking image: ", err)
                        throw NewError(genericErrorMessage)
                    })
                } else if (typeof value === "string") { // if text box entry
                    data.append(key, value)
                }
            }

            // sending the registration data to the backend
            return register(data).then((res) => {
                if (!res.ok) { // failed to register
                    res.json()
                    .then((err: CommonError) => {                      
                        throw NewError(err.message)
                    })
                    .catch((err: any) => {
                        console.log(errorPrefix, "failed to decode json: ", err)
                        throw NewError(genericErrorMessage)
                    })
                }
            })
            .catch((err: any) => {
                console.log(unidentifiedErrorPrefix, "registration failed")
                throw err
            })
        }

        // stops user input during the processing of their registration
        working = true

        submitAsync()
        .then(() => {
            isRegistered = true // play registered animation
        })
        .catch((err: CommonError) => {
            console.log(errorPrefix, err.message)
            submissionError = err.message // change error message displayed to user
            working = false // since their was an error allow the user to try submitting the form again
        })
        .catch((err: any) => {
            console.log(unidentifiedErrorPrefix, err)
            submissionError = genericErrorMessage // change error message displayed to user
            working = false // since their was an error allow the user to try submitting the form again
        })
    }

    // Shrinks an image to a size that can be transfered to the backend
    function shrinkImage(file: File): Promise<string> {
        const promise = new Promise<string>((res, rej) => {
            if(!file.type.match(/image.*/)) { // ensure it's an image
                console.log(errorPrefix, "file was not an image")
                rej(NewError(genericErrorMessage))
            }
            // load the image
            var reader = new FileReader();
            reader.onload = function (readerEvent) {
                var image = new Image();
                image.onload = function () {
                    // resize the image
                    var canvas = document.createElement('canvas');
                    const max_size = 544;
                    let width = image.width;
                    let height = image.height;

                    // transform the width and height to below or equal the max size
                    if (width > height && width > max_size) {
                        width = max_size;
                        height *= max_size / width; // keeping the image's proportion
                    } else if (height > max_size) {
                        height = max_size;
                        width *= max_size / height; // keeping the image's proportion
                    }

                    // resize image
                    canvas.width = width;
                    canvas.height = height;
                    canvas.getContext('2d')?.drawImage(image, 0, 0, width, height);

                    // encode it as a base64 encoded png
                    let encoded = canvas.toDataURL("image/jpeg")
                    console.log("encoded data: ", encoded)

                    // return encoded out of promise
                    res(encoded.replace('data:', '').replace(/^.+,/, ''));
                }

                if (typeof readerEvent.target?.result == 'string') {
                    image.src = readerEvent.target.result; // set the image
                } else {
                    console.log(errorPrefix, "image file did not load correctly")
                    rej(NewError(genericErrorMessage))
                }
            }

            reader.readAsDataURL(file) // input image file into reader
        })

        return promise
    }

    // Submits a registration request to the backend
    function register(data: URLSearchParams): Promise<Response> {
        // default registration request
        let request: RegistrationRequest = {
            first_name: undefined,
            last_name: undefined,
            email: undefined,
            phone_no: undefined,
            proof_of_identification: undefined,
            fingerprint: undefined,
        };

        // fills in each element of the request based on the form data
        data.forEach((value, key) => {
            switch (key) {
                case "firstname": {
                    request.first_name = value
                    break;
                }
                case "surname": {
                    request.last_name = value
                    break;
                }
                case "email": {
                    request.email = value
                    break;
                }
                case "telephone": {
                    request.phone_no = value
                    break;
                }
                case "identification": {
                    request.proof_of_identification = value
                    break;
                }
                case "fingerprint": {
                    request.fingerprint = value
                    break;
                }
                default: {
                    console.log(errorPrefix, "form contained unexpected data")
                    throw NewError(genericErrorMessage)
                }
            }
        })

        // checks all the different attributes are filled
        if (!validateRegistrationRequest(request)) {
            throw NewError("missing data in registration")
        }

        // perform registration request to backend
        return fetch("/api/register", { 
            method:"POST", 
            body: JSON.stringify(request), 
            headers: { 'content-type': 'application/json'} ,
            signal: AbortSignal.timeout(3000),
        });
    }
</script>

<Hero>
    <h1>
        Register
    </h1>
</Hero>
<div class="center-container  spaced-container">
    <!-- successful registration animation -->
    {#if isRegistered}
        <div class="registered-animation">
            <TickAnimation start={true} size="10rem;"/>
            <h1>You are Registered!</h1>
            <p>Await for us to contact you with your login details</p>
            <a class="home-button-link no-format-link" href="/" style="margin-left: auto; margin-right: auto;">
                <button class="home-button button button-background-color-primary color-text-inverted">
                    Return Home
                </button>
            </a>
        </div>
    {/if}
    <!-- div will dissapear if registration animation is playing -->
    <div class:dissapear={isRegistered}>
        <!-- registration form -->
        <form class="form card" on:submit|preventDefault={(e => submit(e))}>

            <!-- card detail elements -->
            <span class="circle-one"/>
            <span class="circle-two"/>
            
            <!-- inputs -->
            <fieldset class = "multi-input">
                <TextInput name="firstname" label="Firstname" required style="width: 250px;" working={working}/>
                <TextInput name="surname" label="Surname" required style="width: 250px;" working={working}/>
            </fieldset>
            <fieldset class = "multi-input">
                <TextInput name="email" type="email" label="Email" required style="width: 250px;" working={working}/>
                <TextInput name="telephone" type="tel" label="Telephone No" required style="width: 250px;" working={working}/>
            </fieldset>

            <!-- file inputs -->
            <FileInput label="Proof of Identification" required name="identification" bind:identificationFiles accept=".png,.jpeg,.jpg" style="padding: 20px; padding-left: 50px; padding-right: 50px;" working={working}/>
            <FileInput label="Fingerprint" required name="fingerprint" bind:fingerprintFiles accept=".png,.jpeg,.jpg"style="padding: 20px; padding-left: 50px; padding-right: 50px;" working={working}/>
            <button class="button button-background-color-primary color-text-inverted" type="submit">Register</button>

            <!-- error message displayed to user -->
            {#if submissionError}
                <span class="background-color-danger" style="padding-top: 10px;">{submissionError}</span>
            {/if}
        </form>
    </div>
</div>

<style lang="scss">
    // causes an element to fade out
    @keyframes dissapear {
        from {
            opacity: 100%;
        }
        to {
            opacity: 0;
        }
    }

    .dissapear {
        animation: dissapear 0.2s ease-in-out forwards;
    }

    // causes an element to fade in
    @keyframes appear {
        from {
            opacity: 0%;
        }
        to {
            opacity: 100%;
        }
    }

    .registered-animation {
        position: absolute; 
        top: 150px; 
        z-index: 100;

        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding-bottom: 80px;

        h1 {
            text-align: center;
            animation: appear 1.5s ease-in-out forwards;
            margin-top: 15px;
            margin-bottom: 0px;
        }

        p {
            margin-top: 0px;
            animation: appear 1.5s ease-in-out forwards;
        }

        .home-button-link {
            display: flex;
            align-items: center;
            justify-content: center;
            .home-button {
                position: absolute;
                top: 270px;
                animation: appear 1.5s ease-in-out forwards;
                padding: 5px;
                font-size: 20px;
                padding-left: 30px;
                padding-right: 30px;
            }
        }
    
    }

    
</style>