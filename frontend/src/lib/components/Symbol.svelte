<script lang="ts">
	import type { colors } from "./colours";

    export let variant: "outlined" | "rounded" | "sharp" = "outlined";

    export let color: colors = "text";

    export let name = "";
</script>

<span class={`symbol color-${color} material-symbols-${variant}`} {...$$restProps}>
    {name}
</span>