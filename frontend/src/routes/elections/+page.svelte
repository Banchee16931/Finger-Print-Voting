<script lang="ts">
	import Card from "$lib/components/Card.svelte";
    import CardList from "$lib/components/CardList.svelte";
	import ElectionList from "$lib/components/ElectionList.svelte";
	import Hero from "$lib/components/Hero.svelte";
	import TextInput from "$lib/components/TextInput.svelte";
    import type { CandidateVotes, ElectionState } from "$lib/types"
	import type { PageData } from "./$types";

    export let data: PageData

    let searchTerm: string | null = null

    $: if (searchTerm?.trim() === "") {
        searchTerm = null
    }

    $: filteredElections = data.elections.filter((election) => {
        if (!searchTerm) {
            return true
        }
        let shouldKeep = true
        searchTerm?.toLowerCase().split(" ").forEach((term) => {
            let passedTerm = false
            if (String(election.election_id) == term
                    || election.location.toLowerCase().includes(term)) {
                        passedTerm = true
            }

            
            election.result.forEach((result) => {
                if (term) {
                    if (result.first_name.toLowerCase().includes(term)
                            || result.last_name.toLowerCase().includes(term)
                            || result.party.toLowerCase().includes(term)) {
                                passedTerm = true
                    }
                }
            })


            if (election.start.toLowerCase().startsWith(term)) {
                passedTerm = true
            }

            if (election.end.toLowerCase().startsWith(term)) {
                passedTerm = true
            }

            if (!passedTerm) {
                shouldKeep = false
            }
        })
        
        return shouldKeep
    }).sort((prev, current) => {
        return new Date(current.start.replace("-", ",")).getTime() - new Date(prev.start.replace("-", ",")).getTime()
    })

    $: console.log("filtered elections: ", filteredElections)

    $: happeningElections = filteredElections.filter((election) => {
        let now = new Date().getTime()
        if (new Date(election.end.replace("-", ",")).getTime() > now) {
            return true
        }

        return false
    }).sort((prev, current) => {
        return new Date(current.start.replace("-", ",")).getTime() - new Date(prev.start.replace("-", ",")).getTime()
    })

    $: completedElections = filteredElections.filter((election) => {
        let now = new Date().getTime()
        if (new Date(election.end.replace("-", ",")).getTime() <= now) {
            return true
        }

        return false
    }).sort((prev, current) => {
        return new Date(current.end.replace("-", ",")).getTime() - new Date(prev.end.replace("-", ",")).getTime()
    })
</script>

<Hero sticky>
    <div class="heroFormat">
        <h1 style="margin-top: 40px">
            Elections
        </h1>
        <TextInput name="searchValue" bind:value={searchTerm} disableFocusHighlight type="search" buttonBackgroundColor="background-secondary" 
        style="width: 250px; margin-left: auto; margin-bottom: 10px;"/>
    </div>
</Hero>

<div class="spaced-container body-container" style="margin-top: 10px;">
    {#if happeningElections.length > 0}
    <h3>Ongoing Elections</h3>
    <ElectionList elections={happeningElections}/>
    {/if}
    {#if completedElections.length > 0}
    <h3>Finished Elections</h3>
    <ElectionList elections={completedElections}/>
    {/if}
</div>

<style lang="scss">
    .heroFormat {
        display: flex; 
        flex-direction: row; 
        align-items: end;
    }
</style>