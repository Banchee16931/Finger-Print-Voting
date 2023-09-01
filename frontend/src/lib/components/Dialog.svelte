

<script lang="ts">
    export let hide = false
    export let title: string
</script>

{#if !hide}
    <div class="bg" class:hide={hide}>
        <div class="dialog" {...$$restProps}>
            <div class="topbar">
                <h3 class="color-text-inverted" style="margin-top: 5px;">
                    {title}
                </h3>
                <button class="xButton color-text-inverted" on:click={() => hide = true}>X</button>
            </div>
            <slot/>
        </div>
    </div>
{/if}

<style lang="scss">
    .hide {
        display: none !important;
    }

    .bg {
        position: fixed;
        overflow: auto;
        z-index: 1000;
        top: 0;
        left: 0;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: 100vw;
        height: 100vh;
        background: rgba(0, 0, 0, 0.66);
    }

    .topbar {
        display: flex;
        flex-direction: row;
        align-items: center;
        position: relative;
        left: 0;
        right: 0;
        top: 0;
        width: 100%;
        padding-left: 15px;
        padding-right: 15px;
        padding-top: 5px;
        background: var(--primary)
    }

    .dialog {
        position: absolute;
        padding: 0;
        margin: 0;
        background: white;
        overflow: clip;
        margin-top: 75px;
    }

    .xButton {
        width: 35px;
        height: 35px;
        font-size: 25px;
        margin-left: auto;
        text-align: center;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: transparent;
        border: none;
        float: right;
    }

    .xButton:hover:not(:active) {
        font-size: 28px;
    }
</style>