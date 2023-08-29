<script>
    import Hero from "../../lib/components/Hero.svelte"
    import TextInput from "../../lib/components/TextInput.svelte"

    export let form;
</script>

<Hero>
    <h1>
        Login
    </h1>
</Hero>
<div class="center-container spaced-container">
    <form class="form card" method="POST" action="?/login">
        <span class="circle-one"/>
        <span class="circle-two"/>
        
        <TextInput name="username" label="Username" required type="text" style="width: 250px;"/>
        <TextInput name="password" label="Password" required type="password" style="width: 250px;"/>

        
        <button class="button button-background-color-primary color-text-inverted" type="submit">Login</button>
        {#if form?.error}
            <span class="error" style="width: 250px;">{form?.error}</span>
        {/if}
    </form>
</div>