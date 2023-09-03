<script lang="ts">
    import CardList from "$lib/components/CardList.svelte"
    import Hero from "$lib/components/Hero.svelte"
    import TextInput from "$lib/components/TextInput.svelte"
	import type { PageData } from "./$types";
    import type { UserAcceptanceRequest, Registrant, LoginRequest } from "$lib/types"
    import { NewError } from "$lib/types/CommonError";
	import Dialog from "$lib/components/Dialog.svelte";

    export let data: PageData

    let errorMsg: string = ""

    // Makes it so the first letter and only the first letter is capatalised in a word
    function nameCapatilsation(name: string): string {
        if (name === undefined) {
            return ""
        }

        if (name.length == 0) {
            return name
        }
        
        return name.charAt(0).toUpperCase()+name.slice(1)
    }

    function accept(e: SubmitEvent & {currentTarget: EventTarget & HTMLFormElement}, id: number) {
        if (e.target) {
            const formData = new FormData(e.currentTarget)
            const data = new URLSearchParams()
            for (let field of formData) {
                const [key, value] = field
                data.append(key, value.toString())
            }

            let newUserDetails: {
                username: string | null
                password: string | null
                confirmPassword: string | null
            } = {
                username: null,
                password: null,
                confirmPassword: null
            };

            data.forEach((value, key) => {
                switch (key) {
                    case "username": {
                        newUserDetails.username = value
                        break;
                    }
                    case "password": {
                        newUserDetails.password = value
                        break;
                    }
                    case "confirm-password": {
                        newUserDetails.confirmPassword = value
                        break;
                    }
                    default: {
                        throw NewError("couldn't process request")
                    }
                }
            })

            if (newUserDetails.username === null 
                    || newUserDetails.password === null
                    || newUserDetails.confirmPassword === null) {
                throw NewError("request is missing data")
            }

            if (newUserDetails.password !== newUserDetails.confirmPassword) {
                errorMsg = "passwords don't match"
                return
            } 
            
            let acceptance: UserAcceptanceRequest = {
                registrant_id: id, 
                accepted: true,
                username: newUserDetails.username,
                password: newUserDetails.password
            }
            fetch("/api/registrations/acceptance", {
                method: "POST",
                body: JSON.stringify(acceptance)
            }).finally(() => {
                updateRegistrantsList()
                hideForm = true
            })
            console.log("accept: ", id)
        }
    }

    function decline(id: number) {
        console.log("decline")
        let acceptance: UserAcceptanceRequest = {
            registrant_id: id, 
            accepted: false,
            username: undefined,
            password: undefined
        }
        fetch("/api/registrations/acceptance", {
            method: "POST",
            body: JSON.stringify(acceptance)
        }).finally(() => {
            updateRegistrantsList()
        })
        
        console.log("declined: ", id)
    }

    function updateRegistrantsList() {
        // update registrants list
        fetch("/api/registrations", { method: "GET" }).then((res) => {
            if (res.ok) {
                res.json().then((registrations: Registrant[]) => {
                    console.log("updated registrants: ", registrations.length)
                    data.registrations = registrations
                })
            }
        })
    }

    let selectedRegistrant: Registrant | null = null

    let hideForm = true

    function showForm() {
        console.log("show form")
        errorMsg = ""
        hideForm = false
    }
</script>

<Hero>
    <h1>
        User Acceptance
    </h1>
</Hero>

{#if selectedRegistrant}
<Dialog bind:hide={hideForm} title="Create New Users Details" style="border-radius: 10px;">
    <form class="form card dialogCard" on:submit|preventDefault={(e) => {if (selectedRegistrant) {accept(e, selectedRegistrant?.registrant_id)}}}>
        <span class="circle-one"/>
        <span class="circle-two"/>
        
        <div class="dialogContent">
            <div class="column" style="margin-right: auto;">
                <TextInput name="username" label="Username" required type="text" style="width: 250px;"/>
                <TextInput name="password" label="Password" required type="password" style="width: 250px;"/>
                <TextInput name="confirm-password" label="Confirm Password" required type="password" style="width: 250px;"/>
            </div>
            <div class="column" style="margin-left: auto;">
                <h4 style="margin-top: 5px;">{nameCapatilsation(selectedRegistrant.first_name)} {nameCapatilsation(selectedRegistrant.last_name)}</h4>
                <span><strong>Local Authority: </strong>{selectedRegistrant.location}</span>
                <span><strong>Phone No: </strong>{selectedRegistrant.phone_no}</span>
                <span><strong>Email: </strong>{selectedRegistrant.email}</span>
            </div>
        </div>
        <button class="button button-background-color-primary color-text-inverted" type="submit">Add User</button>
        
        {#if errorMsg !== ""}
        <span class="error" style="width: 250px;">{errorMsg}</span>
        {/if}
    </form>
</Dialog>
{/if}

<div class="spaced-container body-container">
    <CardList>
        {#each data.registrations as registrant}
            <div class="card">
                <div class="details">
                    <div class="imageContainer">
                        <img src={registrant.proof_of_identity} alt="fingerprint" class="sizing"/>
                    </div>
                    <div class="column">
                        <h4 style="margin-top: 5px;">{nameCapatilsation(registrant.first_name)} {nameCapatilsation(registrant.last_name)}</h4>
                        <span><strong>Local Authority: </strong>{registrant.location}</span>
                        <span><strong>Phone No: </strong>{registrant.phone_no}</span>
                        <span><strong>Email: </strong>{registrant.email}</span>
                        
                    </div>
                    <div class="buttonPanel">
                        <button class="button background-color-primary color-text-inverted"
                            on:click={() => {
                                selectedRegistrant = registrant
                                showForm()
                                }}>Accept</button>
                        <button class="button color-text"
                            on:click={() => decline(registrant.registrant_id)}>Decline</button>
                    </div>
                </div>
                <span class="circle-one"/>
                <span class="circle-two"/>
            </div>
        {/each}
    </CardList>
</div>

<style lang="scss">
    $imageSizing: 200px;

    .dialogContent {
        display: flex;
        flex-direction: row;
        gap: 30px;
        width: 100%;
    }

    .form {
        margin: 0;
    }

    .card {
        padding: 20px;
    }

    .dialogCard {
        border-radius: 0;
    }

    .imageContainer {
        min-width: $imageSizing;
        max-width: $imageSizing;
        height: $imageSizing;
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--background-tertiary);
        margin-top: 5px;
    }

    .column {
        display: flex;
        flex-direction: column;
        width: 100%;
    }

    .details {
        display: flex;
        flex-direction: row;
        gap: 20px;
    }

    .spaced-container {
        padding-top: 10px;;
    }


    .buttonPanel {
        z-index: 100;
        display: flex;
        flex-direction: row;
        gap: 5px;
        height: 30px;
        margin-top: 5px;
        margin-left: auto;
        margin-right: 5px;
    }

    .sizing {
        display: block;
        width: auto;
        max-width: $imageSizing;
        height: auto;
        max-height: $imageSizing;
    }
</style>