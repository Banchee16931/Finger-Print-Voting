<script lang="ts">
    import Hero from "$lib/components/Hero.svelte"
    import TextInput from "$lib/components/TextInput.svelte"
    import Option from '$lib/components/Option.svelte';
    import CardGrid from "$lib/components/CardGrid.svelte";
    import Dialog from "$lib/components/Dialog.svelte";

    export let form;

    let errorMsg = ""
    $: if (form?.error) {
        errorMsg = form.error
    }

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
        party_colour: "#000000",
        photo: ""
    }

    $: candidatesAsJSON = JSON.stringify(candidates)

    let hideNewCandidate = true;

    function NewCandidate() {
        beingCreatedCandidate = {
            first_name: "",
            last_name: "",
            party: "",
            party_colour: "#000000",
            photo: ""
        }

        hideNewCandidate = false
    }

    function AddNewCandidate() {
        hideNewCandidate = true
        if (beingCreatedCandidate.first_name === ""
        || beingCreatedCandidate.last_name === ""
        || beingCreatedCandidate.party === ""
        || beingCreatedCandidate.party_colour === ""
        || beingCreatedCandidate.photo === "") {
            errorMsg = "new candidate was missing required data"
            return
        }
        candidates.push(beingCreatedCandidate)
        console.log(candidates)
        candidates = candidates
    }

    const partyToColour = new Map<string, string>([
        ["Conservative and Unionist Party", "#0087DC"],
        ["Labour Party", "#E4003B"],
        ["Scottish National Party", "#FDF38E"],
        ["Co-operative Party", "#3F1D70"],
        ["Liberal Democrats", "#FAA61A"],
        ["Democratic Unionist Party", "#D46A4C"],
        ["Sinn Féin", "#326760"],
        ["Plaid Cymru", "#005B54"],
        ["Social Democratic and Labour Party", "#2AA82C"],
        ["Ulster Unionist Party", "#48A5EE"],
        ["Green Party of England and Wales", "#02A95B"],
        ["Scottish Greens", "#00B140"],
        ["Alliance Party of Northern Ireland", "#F6CB2F"],
        ["Traditional Unionist Voice", "#0C3A6A"],
        ["People Before Profit", "#E91D50"],
        ["Alba Party", "#005EB8"],
        ["Reclaim Party", "#C03F31"],
        ["Reclaim", "#C03F31"],
        ["Liberal Party", "#EB7A43"],
        ["Reform UK", "#12B6CF"],
        ["Social Democratic Party (SDP)", "#D25469"],
        ["Official Monster Raving Loony Party (OMRLP)", "#FFF000"],
        ["British Democrats (BDP)", "#1C1CF0"],
        ["Breakthrough Party", "#F38B3D"],
        ["Women's Equality Party (WEP)", "#64B69A"],
        ["Animal Welfare Party (AWP)", "#EE3263"],
        ["Climate Party", "#36d0b6"],
        ["Harmony Party UK", "#D60600"],
        ["National Flood Prevention Party", "#DCDCDC"],
        ["Populist Party (UK)", "#D60600"],
        ["Trade Unionist and Socialist Coalition (TUSC)", "#ec008c"],
        ["Independent", "#DCDCDC"]
    ])
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
            <div class="colourHighlight" style={`background: ${beingCreatedCandidate.party_colour}`}/>
        </div>
        <fieldset class = "multi-input" style="margin-top: 20px;">
            <TextInput name="first_name" bind:value={beingCreatedCandidate.first_name} label="First Name" required type="text" style="width: 250px;"/>
            <TextInput name="last_name" bind:value={beingCreatedCandidate.last_name} label="Last Name" required type="text" style="width: 250px;"/>
        </fieldset>
        <datalist id="partylist">
            {#each partyToColour.keys() as party}
                <option value={party}/>
            {/each}
        </datalist>
        <div style="display: flex; flex-direction: row; align-items: center; gap: 10px; width: 100%;">
            <TextInput bind:value={beingCreatedCandidate.party} id="party" name="party" list="partylist" on:change={
                () => {
                    let newColour = partyToColour.get(beingCreatedCandidate.party)
                    if (newColour) {
                        beingCreatedCandidate.party_colour = newColour
                    }
                }
            } type="text" label="Party"/>
            <input name="party_colour" bind:value={beingCreatedCandidate.party_colour} required type="color" style="margin-bottom: 3px; margin-left: auto;"/>
        </div>
        <FileInput required name="identification" bind:value={beingCreatedCandidate.photo} label="Photo" accept=".png,.jpeg,.jpg" style="padding: 20px; padding-left: 50px; padding-right: 50px;"/>
        <button class="button color-text-inverted background-color-primary" style="display: block; margin-left: auto; margin-right: auto;"
        on:click={() => {
            AddNewCandidate()
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
                <div class="colourHighlight" style={`background: ${candidate.party_colour}`}/>
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
        {#if errorMsg !== "" && !showSuccess}
            <span class="error" style="width: 250px;">Error: {errorMsg}</span>
        {/if}
    </form>
    {#if showSuccess}
        <span class="success" style="width: 250px;">Created New Election</span>
    {/if}
</div>

<style lang="scss">
    @use "sass:color";
    @import "../../lib/scss/mixins";

    .colourHighlight {
        position: absolute;
        bottom: 0;
        left: 0;
        width: 100%;
        height: 5px;
    }
    
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
        height: 180px;

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
        background-color: var(--background-tertiary);
        border: 1px solid var(--outline);
        display: flex;
        flex-direction: row;
        gap: 20px;
        height: 180px;

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