<script lang="ts">
	import Cards from "$lib/Cards.svelte";
	import Search from "$lib/Search.svelte";
	import ScrollButton from "$lib/ScrollButton.svelte";
	import type TTRPGSystem from "../interfaces/TTRPGSystem";
	import CardPageSetter from "$lib/CardPageSetter.svelte";
	import { cursor, foundSystems, selectedSystem } from "../stores";
	import { onMount } from "svelte";
	import Icon from "@iconify/svelte";

    let cardPage = 0;
	let topElem: HTMLImageElement;
	let y = 0;
    let searchValue = '';
    let systems: TTRPGSystem[] = [];
    let genres: string[] = [];
    let families: string[] = [];

    let cardsHeight = 0;
    let loadingMore = false;

    // trigger loadMore() when user scrolls to bottom
    $: if (y !== 0 && y >= cardsHeight - 800 && !loadingMore && $cursor !== 0) {
        loadingMore = true;

        loadMore().then(async (newSystems) => {
            systems = [...systems, ...newSystems];
            await new Promise((r) => {
                setTimeout(r, 200);
            });
            loadingMore = false;
        });
    }

    const loadMore = async () => {
        const res = await fetch(`http://localhost:5000/api/systems/all/${$cursor}`);
        let data = await res.json();
       
        cursor.set(data.cursor);
        genres = data.genres;
        families = data.families;
        return data.systems;
    }

    onMount(async () => {
        systems = await loadMore();
    })

    const deselectSystem = () => {
        selectedSystem.set(undefined);
        foundSystems.set([]);
    }

</script>

<svelte:window bind:scrollY={y} />

<main class="flex flex-col flex-wrap justify-center items-center w-11/12 p-6 m-auto">
    <div class="flex flex-col justify-center items-center w-full gap-8 relative">
        <img src="/favicon_bgless.svg" class="w-48" alt="compass" bind:this={topElem}>
        <Search bind:searchValue {genres} {families} />
        <CardPageSetter bind:cardPage />
        {#if $selectedSystem}
            <div class="flex flex-row gap-2 mb-[-1.5rem] text-xl justify-center items-center italic font-poiret-one">
                <h3>Systems similar to:</h3>
                <h3 class="text-goldenrod">
                    {`${$selectedSystem.Title}`}
                </h3>
                <button class="text-eggshell font-poiret-one" on:click={() => deselectSystem()}>
                    <Icon icon="mdi:close-thick" class="text-xl" />
                </button>
            </div>

        {/if}
        <Cards {systems} {cardPage} bind:height={cardsHeight} />
        {#if y > 100}
            <ScrollButton {topElem}/>
        {/if}
    </div>
</main>
