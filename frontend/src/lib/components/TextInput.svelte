<script lang="ts">
    import Symbol from "./Symbol.svelte"

    export let name: string;

    export let type: "text" | "password" | "email" | "search" | "tel" = "text";

    export let label: string | null = null;

    export let placeholder: string | null = null;

    export let required: boolean = false;

    export let valid: boolean = false;

    export let invalid: boolean | string = false;

    export let pattern: string | null = null;

    export let working: boolean = false;

    export let value: string | null = null;

    export let disableFocusHighlight: boolean = false

    export let buttonColor: colors = "text-inverted";
    export let buttonBackgroundColor: backgroundColors = "primary";

    // this gets around the 2-way databindings by setting the input's value on-demand
    const handleInput = (e: Event) => {
        value = (e.target as HTMLInputElement).value
    }

    // adding enter event
    import { createEventDispatcher } from 'svelte';
	import type { backgroundColors, colors } from "./colours";

    const dispatch = createEventDispatcher();
    
    function enter() {
        dispatch('enter', {value: value});
    }

    // trigger enter event when the Enter key is pressed
    function keyDown(e: KeyboardEvent) {
        if (e.key == "Enter") {
            enter()
        }
    }
</script>

<!--
    @component
    # TextInput
    TextInput allows someone to enter information. This acts as an "input" tag but also allows for a bunch of extra features.

    - (required) Use `name` to define how it would be submitted to a <form> element. This will act as this inputs id.
    - Use `type` to define what type of data you are trying to get from the user.
    - Use `value` to get the value from the input.
    - Use `label` to write some text above the input to describe what it is to the user.
    - Use `placeholder` to give some example text to the user.
    - Use `required` to communicate to the user that this input must be filled for the form to be valid.
    - Use `valid` to communicate to the user that their input is correct.
    - Use `invalid` to communicate to the user that their input is incorrect.
    - Use `pattern` to check the validity of the user's input against a regex match statement.
    - Use `working` to show a spining icon so the user knows that you are processing the input.
    - Use `enter` to bind the the user pressing enter in the input or for "search" type pressing the search button.

    ## Usage
    ```tsx
    <TextInput name="example"/>
    ```

    ### Binding the value
    ```tsx
    <script lang="ts">
        let textInputValue: string | null;
    </script>
    <TextInput name="example" bind:value={textInputValue}/>
    ```
    
    ### Binding enter
    This one will print the value in the input to the console when the enter key or the search button is pressed.
    ```tsx
    <TextInput name="hi" type="search" on:enter={(event)=>{console.log("enter: " + event.detail.value)}}/>
    ```

-->
<div class="input {working?"disabled":""}" class:invalid class:valid>
    {#if label}
        <label for={name}>{label}</label>
    {/if}

    
    <div class="inner">
        <input on:keydown={(e) => keyDown(e)}
            id={name}
            {name}
            disabled={working}
            {placeholder}
            {required}
            {type}
            {pattern}
            {...$$restProps}
            on:change
            on:blur
            on:focus
            on:keypress
            on:paste
            on:input={handleInput}
            class:disable-focus-highlight={disableFocusHighlight}
    />
        <button class={`color-${buttonColor} background-color-${buttonBackgroundColor} search-button input-button`} class:input-button-appear={type==="search" && !working} on:click={() => enter()}>
            {#if type==="search" && !working}
            <Symbol name="search" color={buttonColor} variant="outlined"/>
            {/if}
        </button>

        
        <button class={`background-color-${buttonBackgroundColor} input-button`} class:input-button-appear={working} on:click={() => enter()}>
            {#if working}
            <i class="spinner"/>
            {/if}
        </button>
    </div>

    {#if invalid !== false && invalid !== true}
        <small>{invalid}</small>
    {/if}
    
</div>
