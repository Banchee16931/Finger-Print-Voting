<script lang="ts">
    import Footer from "$lib/components/Footer.svelte";
    import Header from "../lib/components/Header.svelte"
    import "../lib/scss/styling.scss";
	import type { LayoutData } from "./$types";

    export let data: LayoutData;


    let resetTimer: () => {};
    function activityDetected() {
        if (resetTimer) {
            resetTimer()
        }
    }
</script>

<div class="interaction-detecter" on:mousemove={activityDetected} on:click={activityDetected} on:keypress={activityDetected} role="none">
    {#if data.user}
        <Header bind:resetInactivityTimer={resetTimer} userDetails={data.user}/>
    {/if}
    <slot/>
    <Footer/>
</div>

<style>
    .interaction-detecter {
        width: 100%;
        height: 100%;
    }
</style>
