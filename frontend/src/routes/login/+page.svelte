<script lang="ts">
    import Hero from "../../lib/components/Hero.svelte"
    import TextInput from "../../lib/components/TextInput.svelte"
    import type { ActionData } from "./$types";

    /** @type {import('./$types').ActionData} */
    export let form: ActionData;
</script>

<Hero>
    <h1>
        Login
    </h1>
</Hero>
<div class="center-container spaced-container body-container">
    <form class="form card" method="POST" action="?/login">
        <span class="circle-one"/>
        <span class="circle-two"/>
        
        <TextInput name="username" label="Username" required type="text" style="width: 250px;"/>
        <TextInput name="password" label="Password" required type="password" style="width: 250px;"/>

        
        <button class="button button-background-color-primary color-text-inverted" type="submit">Login</button>
        {#if form?.error}
            <span class="error" style="width: 250px;">Error: {form?.error}</span>
        {/if}
    </form>
</div>