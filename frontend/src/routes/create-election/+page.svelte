<script lang="ts">
    import Hero from "$lib/components/Hero.svelte"
    import TextInput from "$lib/components/TextInput.svelte"
    import Option from '$lib/components/Option.svelte';
    import CardGrid from "$lib/components/CardGrid.svelte";
    import Dialog from "$lib/components/Dialog.svelte";

    export let form;

    let electionsCreated = 0;

    let showSuccess = false
    $: if (form?.electionsCreated && form?.electionsCreated > electionsCreated) {
        console.log("created")
        showSuccess = true
        setTimeout(() => {
            showSuccess = false
        }, 3000);
        electionsCreated = form.electionsCreated
    }

    $: console.log("showSuccess: ", showSuccess)

    let options: string[] = []
    import localAuthorities from "../../../../local-authorities.json"
	import type { Candidate, CandidateRequest } from "$lib/types";
	import Card from "$lib/components/Card.svelte";
	import FileInput from "$lib/components/FileInput.svelte";
    localAuthorities.forEach((authority) => {
        if (authority.local_authority.length !== 0) {
            options.push(authority.local_authority[0])
        }
    })
    options.sort()

    let candidates: CandidateRequest[] = []

    let beingCreatedCandidate: CandidateRequest = {
        first_name: "",
        last_name: "",
        party: "",
        photo: ""
    }

    $: candidatesAsJSON = JSON.stringify(candidates)

    let hideNewCandidate = true;

    function NewCandidate() {
        beingCreatedCandidate = {
            first_name: "",
            last_name: "",
            party: "",
            photo: ""
        }

        hideNewCandidate = false
    }
</script>

<Hero>
    <h1>
        Create New Election
    </h1>
</Hero>


<Dialog bind:hide={hideNewCandidate} title="Create New Candidate" style="border-radius: 10px;">
    <Card>
        <div class="candidateCard card" style="height: 150px; width: 330px; margin-left: auto; margin-right: auto;">
            {#if beingCreatedCandidate.photo !== ""}
            <img src={beingCreatedCandidate.photo} alt={`Photo of ${beingCreatedCandidate.first_name} ${beingCreatedCandidate.last_name}`} style="background: white; width: 80px; height: 80px;"/>
            {:else}
            <div style="background: white; width: 80px; height: 80px;"/>
            {/if}
            <div style="margin-top: 5px;">
                <strong>{beingCreatedCandidate.first_name} {beingCreatedCandidate.last_name}</strong>
                <div>{beingCreatedCandidate.party}</div>
            </div>
        </div>
        <fieldset class = "multi-input" style="margin-top: 20px;">
            <TextInput name="first_name" bind:value={beingCreatedCandidate.first_name} label="First Name" required type="text" style="width: 250px;"/>
            <TextInput name="last_name" bind:value={beingCreatedCandidate.last_name} label="Last Name" required type="text" style="width: 250px;"/>
        </fieldset>
        <TextInput name="party" bind:value={beingCreatedCandidate.party} label="Party" required type="text" style="width: 530px;"/>
        <FileInput required name="identification" bind:value={beingCreatedCandidate.photo} label="Photo" accept=".png,.jpeg,.jpg" style="padding: 20px; padding-left: 50px; padding-right: 50px;"/>
        <button class="button color-text-inverted background-color-primary" style="display: block; margin-left: auto; margin-right: auto;"
        on:click={() => {
            console.log("adding candidate")
            candidates.push(beingCreatedCandidate)
            console.log(candidates)
            candidates = candidates
            hideNewCandidate = true
        }}>Add Candidate</button>
    </Card>
</Dialog>

<div class="center-container spaced-container">
    <h3>Candidates</h3>
    <CardGrid style="Width: 100%; margin-top: 10px;">
        {#each candidates as candidate}
            <div class="candidateCard card">
                <img src={candidate.photo} alt={`Photo of ${candidate.first_name} ${candidate.last_name}`} style="background: white; width: 80px; height: 80px;"/>
                <div style="margin-top: 5px;">
                    <strong>{candidate.first_name} {candidate.last_name}</strong>
                    <div>{candidate.party}</div>
                </div>
                <button class="xButton" type="button" style="position: absolute; top: 5px; right: 10px; font-size: 20px;">X</button>
            </div>
        {/each}

        <button class="background-color-background card addNewCard" on:click={() => NewCandidate()}>+</button>
    </CardGrid>
    <form class="form card" style="width: 100%" method="POST" action="?/newElection">
        <span class="circle-one"/>
        <span class="circle-two"/>

        <input name="candidates" value={candidatesAsJSON} style="display: none;"/>
        
        <fieldset class = "multi-input" style="margin-top: 20px;">
            <TextInput name="start" label="Start Date" required type="date" style="width: 250px;"/>
            <TextInput name="end" label="End Date" required type="date" style="width: 250px;"/>
        </fieldset>

        <Option style="width: 530px;" name="location" options={options} label="Local Authority"/>

        
        <button class="button button-background-color-primary color-text-inverted" type="submit">Create New Election</button>
        {#if form?.error && !showSuccess}
            <span class="error" style="width: 250px;">{form?.error}</span>
        {/if}
    </form>
    {#if showSuccess}
        <span class="success" style="width: 250px;">Created New Election</span>
    {/if}
</div>

<style lang="scss">
    @use "sass:color";
    @import "../../lib/scss/mixins";

    .button {
        font-size: 1rem;
        padding: 5px;
        width: 200px;
    }

    .addNewCard {
        font-size: 35px;
        display: flex;
        align-items: center;
        justify-content: center;
        opacity: 50%;

    }

    .addNewCard:hover {
        font-size: 38px;
    }

    .xButton {
        padding: 0;
        margin: 0;
        background-color: transparent;
        color: var(--text);
        border: none;
    }

    .xButton:hover {
        scale: 1.1;
    }

    .candidateCard {
        display: flex;
        flex-direction: row;
        gap: 20px;

        @include lg-and-up {
            flex-direction: column;
            gap: 0;
        }
    }

    .success {
        position: relative;
        margin-left: auto;
        margin-right: auto;
        margin-top: 20px;
        text-align: center;
        width: 100%;
        border-radius: 5px;
        z-index: 100;
        padding: 5px;
        padding-top: 5px;
        padding-bottom: 5px;
        background-color: color.adjust(green, $lightness: +54%);
        border: 1px solid color.adjust(green, $lightness: -5%);
        color: color.adjust(green, $lightness: -5%);
        animation: fadeOut 2.8s ease-in-out forwards;
    }

    @keyframes fadeOut { 0% { opacity: 1; visibility: visible; } 100% { opacity: 0; visibility: hidden;  }} 

    
    
</style>