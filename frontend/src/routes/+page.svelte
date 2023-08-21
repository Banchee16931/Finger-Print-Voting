<script>
    import Hero from "../lib/components/Hero.svelte"
    import CardGrid from "../lib/components/CardGrid.svelte"
    import Card from "../lib/components/Card.svelte"
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
        <Card href="/how-to-vote" alt="how to vote?">
            How the Vote?
            <span slot="subtitle">Read a quick help guide on how to vote for your election and how we keep your data safe</span>
        </Card>
        <Card href="/register" alt="register">
            Register
            <span slot="subtitle">To vote you must have an account so make sure you register</span>
        </Card>
        <Card href="/election-results" alt="election results">
            Election Results
            <span slot="subtitle">Get the results of all elections you have and haven’t participated in</span>
        </Card>
    </CardGrid>
</div>

<style lang="scss">
</style>