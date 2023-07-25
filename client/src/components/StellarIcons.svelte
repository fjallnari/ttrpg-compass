<script lang="ts">
    import Icon from "@iconify/svelte";

    export let color: string = "[#E5E7EB]";
    export let trackedMetric: string = "";

</script>

<div class="text-{color} text-2xl opacity-80">
    <button class="absolute left-[50%] translate-x-[-50%] bottom-[84%]"
        on:mouseover={() => trackedMetric = "Complexity"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:brain" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
    <button class="absolute left-[50%] translate-x-[-50%] bottom-0"
        on:mouseover={() => trackedMetric = "Combat"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:sword-cross" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
    <button class="absolute left-[27%] translate-x-[-50%] bottom-6"
        on:mouseover={() => trackedMetric = "Exploration"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:ruler-square-compass" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
    <button class="absolute left-[74%] translate-x-[-50%] bottom-6"
        on:mouseover={() => trackedMetric = "Narrative"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:scroll-text" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
    <button class="absolute left-[19%] translate-x-[-50%] bottom-[41%]"
        on:mouseover={() => trackedMetric = "Customization"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:puzzle" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
    <button class="absolute left-[81%] translate-x-[-50%] bottom-[41%]"
        on:mouseover={() => trackedMetric = "Balance"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:scale-unbalanced" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
    <button class="absolute left-[74%] translate-x-[-50%] bottom-[74%]"
        on:mouseover={() => trackedMetric = "Progression"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:stairs" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
    <button class="absolute left-[27%] translate-x-[-50%] bottom-[74%]"
        on:mouseover={() => trackedMetric = "Versatility"}
        on:focus={() => {}}
    >
        <Icon icon="mdi:shape" 
            class="hover:text-goldenrod transition-colors cursor-pointer"
        />
    </button>
</div>