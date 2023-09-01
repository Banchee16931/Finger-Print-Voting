<script lang="ts">
	import type { CandidateRequest } from "$lib/types";
    import { createEventDispatcher } from 'svelte'

    const dispatch = createEventDispatcher()

    export let candidate: CandidateRequest

    export let hoverEffect = false
    
    export let closable = false

    export let id: number | null = null

    const click = (e: MouseEvent) => {
        if (id) {
            dispatch(`click`, {mouseEvent: e});
        } else {
            dispatch(`click`, {mouseEvent: e, id: id});
        }  
    }

    const close = (e: MouseEvent) => {
        if (id) {
            dispatch(`close`, {mouseEvent: e});
        } else {
            dispatch(`close`, {mouseEvent: e, id: id});
        }
    }
</script>

<div class="changeSize candidateCard card" class:hoverEffect={hoverEffect} 
        on:click={click} 
        on:keydown={()=>{}}
        tabindex=0
        role="button" {...$$restProps}>
    {#if candidate.photo !== ""}
        <img src={candidate.photo} alt={`Photo of ${candidate.first_name} ${candidate.last_name}`} style="background: white; width: 80px; height: 80px;"/>
    {:else}
        <div style="background: white; width: 80px; height: 80px;"/>
    {/if}
    <div style="margin-top: 5px;">
        <strong>{candidate.first_name} {candidate.last_name}</strong>
        <div>{candidate.party}</div>
    </div>
    <div class="colourHighlight" style={`background: ${candidate.party_colour}`}/>
    {#if closable}
    <button class="xButton" type="button" style="position: absolute; top: 5px; right: 10px; font-size: 20px;" on:click={close}>X</button>
    {/if}
</div>

<style lang="scss">
    @import "../scss/mixins";

    .hoverEffect:hover {
        scale: 1.1;
    }

    .colourHighlight {
        position: absolute;
        bottom: 0;
        left: 0;
        width: 100%;
        height: 5px;
    }

    .xButton {
        padding: 0;
        margin: 0;
        background-color: transparent;
        color: var(--text);
        border: none;
        font-weight: 700;
    }

    .xButton:hover {
        scale: 1.1;
    }

    .changeSize {
        flex-direction: row !important;
        gap: 20px !important;
        height: 100px !important;
        @include lg-and-up {
            flex-direction: column !important;
            gap: 0 !important;
            height: 180px !important;
        }
    }

    .candidateCard {
        background-color: var(--background-tertiary);
        border: 1px solid var(--outline);
        display: flex;
        height: 180px;

        flex-direction: column;
        gap: 0;
    }
</style>