<script lang="ts">
	import type { CandidateVotes, ElectionState } from "$lib/types";
	import { votes } from "../../routes/api/mock/data";
	import Card from "./Card.svelte";
	import CardList from "./CardList.svelte";

    export let elections: ElectionState[]

    function total(result: CandidateVotes[]) {
        let count = 0;
        result.forEach((value) => {
            count = count + value.votes
        })

        return count
    }

    function percentage(voteCount: number, result: CandidateVotes[]) {
        let resultTotal = total(result)
        if (resultTotal === 0) {
            return 0
        }

        return (voteCount / resultTotal * 100).toFixed(2)
    }
</script>

<CardList>
    {#each elections as election}
        <Card>
            {election.location}
            <svelte:fragment slot="subtitle">
                <strong>ID:</strong> {election.election_id}<br/>
                <strong>Date:</strong> {election.start} -> {election.end}<br/>
                <strong>Result:</strong><br/>
                <div style="width: 100%; height: 5px; display: flex; flex-direction: row; 
                border: 1px solid var(--outline); background-color:transparent; z-index: 101;">
                    {#each election.result.sort((prev, current) => {
                        return current.votes - prev.votes
                    }) as result}
                        <div style={`width: ${percentage(result.votes, election.result)}%; height: 100%; z-index: 101; background: ${result.party_colour};`}/>
                    {/each}
                </div>
                <table style="margin-top: 10px;">
                    <tr>
                        <th style="text-align: left;">
                            Candidate
                        </th>
                        <th style="padding-left: 30px; text-align: left;">
                            Party
                        </th>
                        <th style="padding-left: 30px; text-align: left;">
                            Votes
                        </th>
                    </tr>
                    {#each election.result.sort((prev, current) => {
                        return current.votes - prev.votes
                    }) as result}
                        <tr>
                            <td style={`border-left: 3px solid ${result.party_colour}; padding-left: 10px;`}>
                                {result.first_name} {result.last_name}
                            </td>
                            <td style="padding-left: 30px;">
                                {result.party}
                            </td>
                            <td style="padding-left: 30px;">
                                {result.votes} : {percentage(result.votes, election.result)}%
                            </td>
                        </tr>
                    {/each}
                </table>
            </svelte:fragment>
        </Card>
        
    {/each}
</CardList>