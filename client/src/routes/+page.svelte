<script lang="ts">
	import Cards from "$lib/Cards.svelte";
	import Search from "$lib/Search.svelte";
	import ScrollButton from "$lib/ScrollButton.svelte";
	import type TTRPGSystem from "../interfaces/TTRPGSystem";
	import CardPageSetter from "$lib/CardPageSetter.svelte";
	import { cursor, foundSystems, selectedSystem, serverURL } from "../stores";
	import { onMount } from "svelte";
	import Icon from "@iconify/svelte";
	import type Filters from "../interfaces/Filters";

    export let data;

    let cardPage: number = 0;
	let topElem: HTMLImageElement;
	let y: number = 0;
    let searchValue = '';
    let systems: TTRPGSystem[] = [];
    let genres: string[] = [];
    let families: string[] = [];

    let cardsHeight: number = 0;
    let loadingMore: boolean = false;

    let filters: Filters = {
        genre: '*',
        family: '*',
    };

    // trigger loadMore() when user scrolls to bottom
    $: if (y !== 0 && y >= cardsHeight - 800 && !loadingMore && $cursor !== 0 && !$selectedSystem) {
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
        const res = await fetch(`${$serverURL}/api/systems/all/${$cursor}`);
        let data = await res.json();
       
        cursor.set(data.cursor);
        genres = data.genres;
        families = data.families;
        return data.systems;
    }

    onMount(async () => {
        if (data.API_URL) {
            serverURL.set(data.API_URL);
        }
        systems = await loadMore();
    })

    const deselectSystem = () => {
        cardPage = 0;
        selectedSystem.set(undefined);
        foundSystems.set([]);
    }

    const resetFilters = () => {
        filters = { 
            genre: '*', 
            family: '*' 
        };
        selectedSystem.set(undefined);
        foundSystems.set([]);
    }

</script>

<svelte:window bind:scrollY={y} />

<main class="flex flex-col flex-wrap justify-center items-center w-11/12 p-6 m-auto">
    <div class="flex flex-col justify-center items-center w-full gap-8 relative">
        <img src="/favicon_bgless.svg" class="w-48" alt="compass" bind:this={topElem}>
        <Search bind:searchValue bind:filters {genres} {families} />
        <CardPageSetter bind:cardPage />
        {#if $selectedSystem}
            <div class="flex flex-row gap-2 mb-[-1.5rem] text-xl justify-center items-center">
                <h3 class="text-goldenrod font-cinzel">Systems similar to:</h3>
                <h3 class="text-inherit font-poiret-one">
                    {`${$selectedSystem.Title}`}
                </h3>
                <button class="text-eggshell font-poiret-one" on:click={() => deselectSystem()}>
                    <Icon icon="mdi:close-thick" class="text-xl" />
                </button>
            </div>
        {/if}
        {#if Object.entries(filters).map(([_, v], i) => v !== '*').includes(true)}
            <div class="flex flex-row flex-wrap gap-2 mb-[-1.5rem] text-xl justify-center items-center">
                <Icon 
                    icon="mdi:filter-multiple" 
                    class="text-2xl text-goldenrod"
                />
                {#each Object.entries(filters) as [key, value]}
                    {#if value !== '*'}
                        <div class="flex flex-row gap-2">
                            <h3 class="text-goldenrod font-cinzel">
                                {`${key}:`}
                            </h3>
                            <h3 class="text-inherit font-poiret-one">
                                {` ${value.replaceAll('_', ' ')}`}
                            </h3>
                        </div>
                    {/if}
                {/each}
                <button class="text-inherit font-poiret-one" on:click={() => resetFilters()}>
                    <Icon icon="mdi:close-thick" class="text-xl" />
                </button>
            </div>
        {/if}
        <Cards {systems} {cardPage} bind:height={cardsHeight} bind:filters />
        {#if y > 100}
            <ScrollButton {topElem}/>
        {/if}
    </div>
</main>
