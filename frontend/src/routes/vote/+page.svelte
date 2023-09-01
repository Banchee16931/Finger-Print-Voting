<script lang="ts">
	import { invalidateAll } from "$app/navigation";
	import CandidateCard from "$lib/components/CandidateCard.svelte";
	import CardGrid from "$lib/components/CardGrid.svelte";
	import Dialog from "$lib/components/Dialog.svelte";
	import FileInput from "$lib/components/FileInput.svelte";
    import Hero from "$lib/components/Hero.svelte";
	import TickAnimation from "$lib/components/TickAnimation.svelte";
	import type { Candidate, Election } from "$lib/types";

    let election: Election = {
        election_id: 0,
        start: "2023-10-2",
        end: "2023-10-3",
        location: "Bedfordshire",
        candidates: [
            {
                candidate_id: 0,
                photo: "",
                first_name: "Hi",
                last_name: "Bye",
                party: "Labour",
                party_colour: "#ff0000",
            },
            {
                candidate_id: 1,
                photo: "",
                first_name: "Other",
                last_name: "Guy",
                party: "Conservitive",
                party_colour: "#0000ff",
            },
            {
                candidate_id: 2,
                photo: "",
                first_name: "Green",
                last_name: "Guy",
                party: "Green",
                party_colour: "#00ff00",
            }
        ]
    }

    let hideVoteDialog = true

    function VoteScreenForCandidate(idx: number) {
        selectedCandidate = election.candidates[idx]
        hideVoteDialog = false
    }

    function Vote(id: number) {
        console.log("voted for: ", id)
    }

    let selectedCandidate: Candidate;

    let voted = false

    async function  logout(e: Event) {
        await fetch("/api/delete-authorization", { method: "DELETE" })
        invalidateAll()
    }
</script>

<Hero>
    <h1>
        Vote
    </h1>
    <p>Click on the Candidate you want to vote for</p>
</Hero>

{#if selectedCandidate}
    <Dialog bind:hide={hideVoteDialog} title="Submit your Vote" style="border-radius: 10px;">
        <form class="form card" style="margin-top: 0; border-radius: 0;">
            <span class="circle-one"/>
            <span class="circle-two"/>
            <CandidateCard candidate={selectedCandidate}  style="width: 330px; margin-left: auto; margin-right: auto; margin-bottom: 50px;"/>
            <FileInput required name="fingerprint" bind:value={selectedCandidate.photo} label="Fingerprint" accept=".png,.jpeg,.jpg" 
            style="padding: 20px; padding-left: 50px; padding-right: 50px;"/>
            <button class="button color-text-inverted background-color-primary" style="display: block; margin-left: auto; margin-right: auto;"
            on:click={() => {
                Vote(selectedCandidate.candidate_id)
            }}>Vote</button>
        </form>
    </Dialog>
{/if}

<div class="spaced-container">
    {#if voted}
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
    {/if}
    <div  class:dissapear={voted}>
        <h3>Candidates</h3>
        <CardGrid style="Width: 100%; margin-top: 10px;">
            {#each election.candidates as candidate, idx}
                <CandidateCard candidate={candidate} windowSizeReactive id={candidate.candidate_id} hoverEffect on:click={() => VoteScreenForCandidate(idx)}/>
            {/each}
        </CardGrid>
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