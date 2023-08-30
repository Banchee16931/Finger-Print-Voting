<script lang="ts">
    import Symbol from "./Symbol.svelte";
    
    export let start: boolean = true;

    export let size: string = "1rem";
</script>
{#if start}
<div class="tick" {...$$restProps} style={`width: ${size}; height: ${size};`}>
    <Symbol name="check_circle" variant="outlined" color="success" style={`font-size: ${size};`}/>
</div>
{/if}

<style>  
    .tick {
        text-align: center;
        display: flex;
        align-items: center;
        justify-content: center;
        animation: grow 0.8s ease-in-out forwards;
    }


    @keyframes grow {
        from {
            scale: 0;
        }
        to {
            scale: 1;
        }
    }

</style>