<script>
    export let href = ""
    export let alt = ""
</script>

{#if href != "" && alt != ""}
<a class="card link-card animated-card no-format-link" href={href}>
    <span class="circle-one"/>
    <span class="circle-two"/>
    <div class="content">
        <h5><slot/></h5>
        <span class="ceveron">›</span>
        <span class="subtitle"><slot name="subtitle"/></span>
    </div>
</a>
{:else}
<div class="card">
    <span class="circle-one"/>
    <span class="circle-two"/>
    <div class="content">
        <h5><slot/></h5>
        <span class="subtitle"><slot name="subtitle"/></span>
    </div>
</div>
{/if}