<script lang="ts">
    import Symbol from "./Symbol.svelte"
    
    import type { colors, backgroundColors } from "./colours"

    export let name: string

    export let label: string | undefined = undefined

    export let options: string[]

    export let working: boolean = false;

    export let buttonColor: colors = "text-inverted";
    export let buttonBackgroundColor: backgroundColors = "primary";

   export let value: string | undefined = undefined;

    function show_hide(e: MouseEvent) {
        var click = document.getElementById(name);
        if (click) {
            console.log("dispatching event")
            var clickEvent = document.createEvent ('MouseEvents');
            clickEvent.initEvent("mousedown", true, true);
            click.dispatchEvent(clickEvent)
        }
    }
</script>

<div class="input">
    {#if label}
        <label for={name}>{label}</label>
    {/if}
    <div class="inner">
        <input
            id={name}
            name={name}
            type="text"
            value={value}
            {...$$restProps}
            on:dragenter
            on:dragover
            on:dragleave
            on:click
            on:keydown
            on:keyup
            on:keypress
            style="display: none;"
            />
        <select bind:value={value} on:mousedown={() => {console.log("yay")}}>
            {#each options as option}
                <option value={option}>{option}</option>
            {/each}
        </select>

        <button style="pointer-events: none;" id={name+"-button"} class={`color-${buttonColor} background-color-${buttonBackgroundColor}  input-button`} 
            class:input-button-appear={!working} on:mousedown={(e) => show_hide(e)}>
            {#if !working}   
                <Symbol name="arrow_drop_down" color={buttonColor} variant="rounded"/>
            {/if}
        </button>

        <button class={`background-color-${buttonBackgroundColor} input-button`} class:input-button-appear={working}>
            {#if working}
                <i class="spinner"/>
            {/if}
        </button>
    </div>
</div>