<script lang="ts">
	import { invalidateAll } from "$app/navigation";
	import CandidateCard from "$lib/components/CandidateCard.svelte";
	import CardGrid from "$lib/components/CardGrid.svelte";
	import Dialog from "$lib/components/Dialog.svelte";
	import FileInput from "$lib/components/FileInput.svelte";
    import Hero from "$lib/components/Hero.svelte";
	import TickAnimation from "$lib/components/TickAnimation.svelte";
	import type { Candidate } from "$lib/types";
	import { redirect } from "@sveltejs/kit";
	import type { ActionData, PageData } from "./$types";

    export let data: PageData
    $: election = data.election

    let hideVoteDialog = true

    let working = false

     /** @type {import('./$types').ActionData} */
     export let form: ActionData;

    $: if (form?.registered !== undefined) {
        if (form?.registered === false) {
            working = false
        }
    }

    let selectedCandidate: Candidate;

    function VoteScreenForCandidate(idx: number) {
        if (election) {
            selectedCandidate = election.candidates[idx]
            hideVoteDialog = false
        }
    }

    async function  logout(e: Event) {
        await fetch("/api/delete-authorization", { method: "DELETE" })
        invalidateAll()
        redirect (303, "/")
    }
</script>

{#if election}
<Hero>
    <h1 style="margin-bottom: 0;">
        {election?.location} Election
    </h1>
    <h2 style="margin-top: 0;">
        {election?.start.split("-").reverse().join("/")} - {election?.end.split("-").reverse().join("/")}
    </h2>
    <p><strong>Election ID: </strong> #{election?.election_id}</p>
    <p>Click on the Candidate you want to vote for</p>
</Hero>

{#if selectedCandidate}
    <Dialog bind:hide={hideVoteDialog} title="Submit your Vote" style="border-radius: 10px;">
        <form class="form card" method="POST" action="?/vote" on:submit={() => working = true} style="margin-top: 0; border-radius: 0;">
            <span class="circle-one"/>
            <span class="circle-two"/>
            <CandidateCard candidate={selectedCandidate}  style="width: 330px; margin-left: auto; margin-right: auto; margin-bottom: 50px;"/>
            <FileInput required name="fingerprint" label="Fingerprint" accept=".png,.jpeg,.jpg" 
                style="padding: 20px; padding-left: 50px; padding-right: 50px;" working={working}/>
            <input name="election_id" value={election?.election_id} style="display: none;"/>
            <input name="candidate_id" value={selectedCandidate.candidate_id} style="display: none;"/>
            <button class="button color-text-inverted background-color-primary" style="display: block; margin-left: auto; margin-right: auto;"
            type="submit">Vote</button>
        </form>
    </Dialog>
{/if}

<div class="spaced-container body-container">
    <!-- error message displayed to user -->
    {#if form?.error}
        <span class="inPageError" style="width: 250px;">Error: {form?.error}</span>
    {/if}
    <h3>Candidates</h3>
    <CardGrid style="Width: 100%; margin-top: 10px;">
        {#each election?.candidates as candidate, idx}
            <CandidateCard candidate={candidate} windowSizeReactive id={candidate.candidate_id} hoverEffect on:click={() => VoteScreenForCandidate(idx)}/>
        {/each}
    </CardGrid>
</div>
{:else}
    <div class="spaced-container body-container">
        <div class="center-container registered-animation">
            <TickAnimation start={true} size="10rem;"/>
            <h1>You have Voted!</h1>
            <p>Look at the results of the election on the Elections Page</p>
            <button class="home-button button button-background-color-primary color-text-inverted" style="margin-left: auto; margin-right: auto;"
            on:click={(e) => {
                logout(e)
            }}>
                Logout
            </button>
        </div>
    </div>
{/if}

<style lang="scss">
    @use "sass:color";

    .inPageError {
        display: block;
        margin-top: 20px;
        text-align: center;
        border-radius: 5px;
        z-index: 100;
        padding: 5px;
        padding-top: 5px;
        padding-bottom: 5px;
        margin-left: auto;
        margin-right: auto;
        background-color: color.adjust(#E50000, $lightness: +54%);
        border: 1px solid color.adjust(red, $lightness: -5%);
        color: color.adjust(red, $lightness: -5%);
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
        top: 250px; 
        left: 0;
        right: 0;
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
</style>
