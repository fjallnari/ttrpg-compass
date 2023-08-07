<script lang="ts">
	import Cards from "$lib/Cards.svelte";
	import Search from "$lib/Search.svelte";
	import ScrollButton from "$lib/ScrollButton.svelte";
	import type TTRPGSystem from "../interfaces/TTRPGSystem";
	import CardPageSetter from "$lib/CardPageSetter.svelte";
	import { cursor } from "../stores";
	import { onMount } from "svelte";

    let cardPage = 0;
	let topElem: HTMLImageElement;
	let y = 0;
    let searchValue = '';
    let systems: TTRPGSystem[] = [];
    let genres: string[] = [];

    let cardsHeight = 0;
    let loadingMore = false;

    $: if (y !== 0 && y >= cardsHeight - 700 && !loadingMore && $cursor !== 0) {
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
        const res = await fetch(`http://localhost:5000/api/systems/${$cursor}`);
        let data = await res.json();
       
        cursor.set(data.cursor);
        genres = data.genres;
        return data.systems;
    }

    onMount(async () => {
        systems = await loadMore();
    })

</script>

<svelte:window bind:scrollY={y} />

<main class="flex flex-col flex-wrap justify-center items-center w-11/12 p-6 m-auto">
    <div class="flex flex-col justify-center items-center w-full gap-8 relative">
        <img src="/favicon_bgless.svg" class="w-48" alt="compass" bind:this={topElem}>
        <Search bind:searchValue {systems} {genres} />
        <CardPageSetter bind:cardPage />
        <Cards {systems} {searchValue} {cardPage} bind:height={cardsHeight} />
        {#if y > 100}
            <ScrollButton {topElem}/>
        {/if}
    </div>
</main>
