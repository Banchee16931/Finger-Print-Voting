<script lang="ts">
    import Hero from "../../lib/components/Hero.svelte"
    import TextInput from "../../lib/components/TextInput.svelte"
	import FileInput from '$lib/components/FileInput.svelte';
    import Option from '$lib/components/Option.svelte';
	import TickAnimation from '$lib/components/TickAnimation.svelte';
	import type { ActionData } from "./$types";

    /** @type {import('./$types').ActionData} */
    export let form: ActionData;

    $: if (form?.registered !== undefined) {
        if (form?.registered === false) {
            working = false
        }
    }

    // The file given by the user that stores the proof of identification
    let identificationFiles: FileList| null = null
    // The file given by the user that stores their fingerprint
    let fingerprintFiles: FileList| null = null

    // Options
    let options: string[] = []
    import localAuthorities from "../../../../local-authorities.json"
    localAuthorities.forEach((authority) => {
        if (authority.local_authority.length !== 0) {
            options.push(authority.local_authority[0])
        }
    })
    options.sort()

    // If the registered screen/animation should appear
    let isRegistered: boolean = false;

    $: if (form) {
        if (form.registered != undefined) {
            isRegistered = form.registered
        }
    }

    // Stops the user interacting with the elements as the form is being processed
    let working: boolean = false
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
        <form class="form card" method="POST" action="?/register" on:submit={() => working = true}>

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

            <Option name="location" options={options} label="Local Authority" working={working}/>

            <!-- file inputs -->
            <FileInput label="Proof of Identification" required name="identification" bind:identificationFiles accept=".png,.jpeg,.jpg" style="padding: 20px; padding-left: 50px; padding-right: 50px;" working={working}/>
            <FileInput label="Fingerprint" required name="fingerprint" bind:fingerprintFiles accept=".png,.jpeg,.jpg"style="padding: 20px; padding-left: 50px; padding-right: 50px;" working={working}/>
            <button class="button button-background-color-primary color-text-inverted" type="submit">Register</button>

            <!-- error message displayed to user -->
            {#if form?.error}
                <span class="error" style="width: 250px;">{form?.error}</span>
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