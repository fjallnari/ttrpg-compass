<script lang="ts">
	import Icon from '@iconify/svelte';
  	import Card from './Card.svelte';
  	import LoadCard from './LoadCard.svelte';
	import { fade, slide } from 'svelte/transition';
	import Autocomplete from './Autocomplete.svelte';
  	import type TTRPGSystem from '../interfaces/TTRPGSystem';
    import ScrollButton from './ScrollButton.svelte';

	//export let icon: string;
	//export let width: string = '2em';
	//export let height: string = width;

	export let systems: TTRPGSystem[] = [];

	let navPage = 0;
	let filtersMenuActive = false;
	let searchValue = '';
	let topElem: HTMLImageElement;
	let y = 0;

</script>

<svelte:window bind:scrollY={y}/>

<div class="flex flex-col justify-center items-center w-full gap-8 relative">
	<!-- <Icon icon="mdi:compass-rose" class="text-goldenrod text-8xl"/> -->
	<img src="/favicon_bgless.svg" class="w-48" alt="compass" bind:this={topElem}>
	<div class="flex flex-col items-center w-72 md:w-3/5 lg:w-1/3 p-2 border-b-[3px] border-transparent border-b-goldenrod border-solid rounded-t bg-abyss-900 bg-opacity-75 focus-within:bg-opacity-50 focus-within:shadow-xl shadow backdrop-blur-md">
		<div class="flex items-center flex-row w-full gap-2">
			<button on:click={() => {}}>
				<Icon icon="mdi:search" class="text-2xl active:text-goldenrod"/>
			</button>
			<Autocomplete bind:value={searchValue} systems={systems.map(sys => sys.Title)}/>
			<button on:click={() => filtersMenuActive = !filtersMenuActive}>
				<Icon icon="mdi:filter-multiple" class="text-xl active:text-goldenrod {filtersMenuActive ? "text-goldenrod": ""}"/>
			</button>
		</div>
		{#if filtersMenuActive}
			<div transition:slide class="h-20 w-full m-2 p-2 font-poiret-one">
				Here be dragons
			</div>
		{/if}
	</div>
	<div class="flex justify-center items-center gap-4 text-opacity-90">
        {#each Array(3) as _, index}
            <button class="{navPage === index ? 'text-goldenrod': ''} cursor-pointer text-4xl transition-colors" on:click={() => {navPage = index}}>
                <Icon icon="mdi:rhombus" />
            </button>
        {/each}
    </div>
	<div class="flex flex-row flex-wrap justify-center mt-6 h-full w-full gap-8">
		{#each systems.filter(system => searchValue === '' ? true : system.Title.toLowerCase() === searchValue.toLowerCase()) as system}
			{#await new Promise((r) => setTimeout(r, 1000))}
				<LoadCard/>
			{:then _}
				<Card {system} navPage={navPage}/>
			{/await}
		{/each}
	</div>
	{#if y > 100}
		<ScrollButton {topElem}/>
	{/if}
</div>