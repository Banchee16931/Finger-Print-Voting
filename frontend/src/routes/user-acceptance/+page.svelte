<script lang="ts">
    import Card from "$lib/components/Card.svelte"
    import CardList from "$lib/components/CardList.svelte"
    import Hero from "$lib/components/Hero.svelte"
	import type { PageData } from "./$types";

    export let data: PageData

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

    function accept(id: number) {
        console.log("accept: ", id)
    }

    function decline(id: number) {
        console.log("declined: ", id)
    }
</script>

<Hero>
    <h1>
        User Acceptance
    </h1>
</Hero>

<div class="spaced-container">
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
                            on:click={() => {accept(registrant.registrant_id)}}>Accept</button>
                        <button class="button color-text"
                            on:click={() => {decline(registrant.registrant_id)}}>Decline</button>
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

    .card {
        padding: 15px;
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