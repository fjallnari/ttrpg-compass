<script lang="ts">
	import Cards from "$lib/Cards.svelte";
	import Search from "$lib/Search.svelte";
	import type { PageData } from "./$types";
	import ScrollButton from "$lib/ScrollButton.svelte";
	import type TTRPGSystem from "../interfaces/TTRPGSystem";
	import CardPageSetter from "$lib/CardPageSetter.svelte";

    export let data: PageData;

    let cardPage = 0;
	let topElem: HTMLImageElement;
	let y = 0;
    let searchValue = '';

    $: systems = data.systems as TTRPGSystem[];

    // $: if (y == 1000) {
    //     console.log('scrolling');
    // }

    // function handleListScroll(
    //   e: UIEvent & {
    //     currentTarget: EventTarget & HTMLUListElement
    //   },
    // ) {
    //   if (
    //     e.currentTarget.clientHeight + e.currentTarget.scrollTop >=
	// 	e.currentTarget.scrollHeight - 300
    //   ) {
	// 	console.log('load more');
    //   }
    // }

</script>

<svelte:window bind:scrollY={y} />

<main class="flex flex-col flex-wrap justify-center items-center w-11/12 p-6 m-auto">
    <div class="flex flex-col justify-center items-center w-full gap-8 relative">
        <img src="/favicon_bgless.svg" class="w-48" alt="compass" bind:this={topElem}>
        <Search bind:searchValue systems={systems} />
        <CardPageSetter bind:cardPage />
        <Cards {systems} {searchValue} {cardPage}/>
        {#if y > 100}
            <ScrollButton {topElem}/>
        {/if}
    </div>
</main>