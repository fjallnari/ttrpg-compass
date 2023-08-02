<script lang="ts">
    import Icon from "@iconify/svelte";

    export let name: string;
    export let icon: string;
    export let trackedMetric: string = "";
    export let position: string = "";

</script>

<button class="{position}"
    on:mouseover={() => trackedMetric = name}
    on:focus={() => {}}
    on:click={() => trackedMetric = name}
>
    <Icon icon={icon}
        class="hover:text-goldenrod active:text-goldenrod transition-colors cursor-pointer"
    />
</button>