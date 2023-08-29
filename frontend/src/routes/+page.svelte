<script lang="ts">
    import Hero from "../lib/components/Hero.svelte"
    import CardGrid from "../lib/components/CardGrid.svelte"
    import Card from "../lib/components/Card.svelte"
    import type { PageData } from "./$types";

    export let data: PageData

    // The access level of the user
    let userLevel: null | "admin" | "user" = null

    $: if (data.user) {
        userLevel = data.user.level
    }
</script>

<Hero>
    <h1>
        Welcome to The UK Local Election System
    </h1>
    <p>
        Vote for or find the results of elections in a secure, safe and reliable way
    </p>
</Hero>

<div class="spaced-container" style="padding-top: 25px;">
    <CardGrid>
        <Card href="/elections" alt="election results">
            Elections
            <span slot="subtitle">View the progress of different elections</span>
        </Card>
        {#if userLevel !== "admin"}
            <Card href="/how-to-vote" alt="how to vote?">
                How the Vote?
                <span slot="subtitle">Read a quick help guide on how to vote for your election and how we keep your data safe</span>
            </Card>
        {/if}
        {#if userLevel === null}
            <Card href="/register" alt="register">
                Register
                <span slot="subtitle">To vote you must have an account so make sure you register</span>
            </Card>
        {:else if userLevel === "user"}
            <Card href="/vote" alt="how to vote?">
                Vote
                <span slot="subtitle">Vote for an election</span>
            </Card>
        {:else if userLevel === "admin"}
            <Card href="/create-election" alt="register">
                Create Election
                <span slot="subtitle">Create a new election for your users to vote in</span>
            </Card>
            <Card href="/user-acceptance" alt="register">
                User Acceptance
                <span slot="subtitle">View user registrations and accept or deny their request</span>
            </Card>
        {/if}
    </CardGrid>
</div>

<style lang="scss">
</style>