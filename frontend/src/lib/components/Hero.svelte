<script lang="ts">
    export let sticky = false
</script>

<div class="hero" class:stickyHero={sticky}>
    <div class="spaced-container">
        <slot/>
    </div>
</div>