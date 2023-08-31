<script lang="ts">
	import Card from "$lib/components/Card.svelte";
    import CardList from "$lib/components/CardList.svelte";
	import ElectionList from "$lib/components/ElectionList.svelte";
	import Hero from "$lib/components/Hero.svelte";
	import TextInput from "$lib/components/TextInput.svelte";
    import type { CandidateVotes, ElectionState } from "$lib/types"

    let elections: ElectionState[] = [
        {
            election_id: 0,
            start: "2023-10-2",
            end: "2023-10-3",
            location: "Bedfordshire",
            result: [
                {
                    first_name: "Hi",
                    last_name: "Bye",
                    party: "Labour",
                    party_colour: "#ff0000",
                    votes: 15
                },
                {
                    first_name: "Other",
                    last_name: "Guy",
                    party: "Conservitive",
                    party_colour: "#0000ff",
                    votes: 5
                },
                {
                    first_name: "Green",
                    last_name: "Guy",
                    party: "Green",
                    party_colour: "#00ff00",
                    votes: 2
                }
            ]
        },
        {
            election_id: 2,
            start: "2022-10-2",
            end: "2022-10-3",
            location: "Hertfordshire",
            result: [
                {
                    first_name: "Hi",
                    last_name: "Bye",
                    party: "Labour",
                    party_colour: "#ff0000",
                    votes: 9
                },
                {
                    first_name: "Other",
                    last_name: "Guy",
                    party: "Conservitive",
                    party_colour: "#0000ff",
                    votes: 1
                },
                {
                    first_name: "Green",
                    last_name: "Guy",
                    party: "Green",
                    party_colour: "#00ff00",
                    votes: 5
                }
            ]
        },
        {
            election_id: 1,
            start: "2022-4-2",
            end: "2023-4-3",
            location: "Bedfordshire",
            result: [
                {
                    first_name: "Hi",
                    last_name: "Bye",
                    party: "Independant",
                    party_colour: "#eeeeee",
                    votes: 2
                },
                {
                    first_name: "Other",
                    last_name: "Guy",
                    party: "Conservitive",
                    party_colour: "#0000ff",
                    votes: 69
                },
                {
                    first_name: "Green",
                    last_name: "Guy",
                    party: "Green",
                    party_colour: "#00ff00",
                    votes: 15
                }
            ]
        },
        {
            election_id: 3,
            start: "2023-8-31",
            end: "2023-9-1",
            location: "York",
            result: [
                {
                    first_name: "Hi",
                    last_name: "Bye",
                    party: "Independant",
                    party_colour: "#eeeeee",
                    votes: 2
                },
                {
                    first_name: "Other",
                    last_name: "Guy",
                    party: "Conservitive",
                    party_colour: "#0000ff",
                    votes: 69
                },
                {
                    first_name: "Green",
                    last_name: "Guy",
                    party: "Green",
                    party_colour: "#00ff00",
                    votes: 15
                }
            ]
        }
    ]

    let searchTerm: string | null = null
    $: console.log(searchTerm)

    $: if (searchTerm?.trim() === "") {
        searchTerm = null
    }

    $: filteredElections = elections.filter((election) => {
        if (!searchTerm) {
            return true
        }
        let shouldKeep = true
        searchTerm?.toLowerCase().split(" ").forEach((term) => {
            let passedTerm = false
            console.log("hello: ", term)
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
        
        console.log("should keep: ", shouldKeep)
        return shouldKeep
    }).sort((prev, current) => {
        return new Date(current.start.replace("-", ",")).getTime() - new Date(prev.start.replace("-", ",")).getTime()
    })

    $: happeningElections = filteredElections.filter((election) => {
        let now = new Date().getTime()
        if (new Date(election.end.replace("-", ",")).getTime() > now) {
            return true
        }

        return false
    })

    $: completedElections = filteredElections.filter((election) => {
        let now = new Date().getTime()
        if (new Date(election.end.replace("-", ",")).getTime() <= now) {
            return true
        }

        return false
    })
</script>

<Hero>
    <div class="heroFormat">
        <h1 style="margin-top: 40px">
            Elections
        </h1>
        <TextInput name="searchValue" bind:value={searchTerm} disableFocusHighlight type="search" buttonBackgroundColor="background-secondary" 
        style="width: 250px; margin-left: auto; margin-bottom: 10px;"/>
    </div>
</Hero>

<div class="spaced-container" style="margin-top: 10px;">
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