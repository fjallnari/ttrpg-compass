<script lang="ts">
	import Icon from "@iconify/svelte";

    export let cardPage = 0;

</script>

<div class="flex justify-center items-center gap-4 text-opacity-90">
    {#each Array(3) as _, index}
        <button class="{cardPage === index ? 'text-goldenrod': ''} cursor-pointer text-4xl transition-colors" on:click={() => {cardPage = index}}>
            <Icon icon="mdi:rhombus" />
        </button>
    {/each}
</div>