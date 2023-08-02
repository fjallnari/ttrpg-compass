<script lang="ts">
    import StellarIcon from "./StellarIcon.svelte";

    export let trackedMetric: string = "";

</script>

<div class="text-2xl opacity-80">
    <StellarIcon
        name="Complexity"
        icon="mdi:brain"
        position="absolute left-[50%] translate-x-[-50%] bottom-[84%]"
        bind:trackedMetric={trackedMetric}
    />
    <StellarIcon
        name="Combat"
        icon="mdi:sword-cross"
        position="absolute left-[50%] translate-x-[-50%] bottom-0"
        bind:trackedMetric={trackedMetric}
    />
    <StellarIcon
        name="Exploration"
        icon="mdi:ruler-square-compass"
        position="absolute left-[27%] translate-x-[-50%] bottom-6"
        bind:trackedMetric={trackedMetric}
    />
    <StellarIcon
        name="Narrative"
        icon="mdi:scroll-text"
        position="absolute left-[74%] translate-x-[-50%] bottom-6"
        bind:trackedMetric={trackedMetric}
    />
    <StellarIcon
        name="Customization"
        icon="mdi:puzzle"
        position="absolute left-[19%] translate-x-[-50%] bottom-[41%]"
        bind:trackedMetric={trackedMetric}
    />
    
    <StellarIcon
        name="Balance"
        icon="mdi:scale-unbalanced"
        position="absolute left-[81%] translate-x-[-50%] bottom-[41%]"
        bind:trackedMetric={trackedMetric}
    />
    <StellarIcon
        name="Progression"
        icon="mdi:stairs"
        position="absolute left-[74%] translate-x-[-50%] bottom-[74%]"
        bind:trackedMetric={trackedMetric}
    />
    <StellarIcon
        name="Versatility"
        icon="mdi:shape"
        position="absolute left-[27%] translate-x-[-50%] bottom-[74%]"
        bind:trackedMetric={trackedMetric}
    />
</div>