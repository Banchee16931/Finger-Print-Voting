<script>
    export let href = ""
    export let alt = ""
</script>

{#if href != "" && alt != ""}
<a class="card no-format-link" href={href} alt="other">
    <span class="circle-one"/>
    <span class="circle-two"/>
    <div class="content">
        <h5 class><slot/></h5>
        <span class="ceveron">›</span>
        <span class="subtitle"><slot name="subtitle"/></span>
    </div>
</a>
{:else}
<div class="card">
    <span class="circle-one"/>
    <span class="circle-two"/>
    <div class="content">
        <h5 class><slot/></h5>
        <span class="ceveron">›</span>
        <span class="subtitle"><slot name="subtitle"/></span>
    </div>
</div>
{/if}