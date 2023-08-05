<script lang="ts">
  	import Card from './Card.svelte';
  	import LoadCard from './LoadCard.svelte';
  	import type TTRPGSystem from '../interfaces/TTRPGSystem';

	export let systems: TTRPGSystem[] = [];
	export let searchValue = '';
	export let cardPage = 0;

	export let height = 0;

</script>

<div bind:clientHeight={height} class="flex flex-row flex-wrap justify-center mt-6 h-full w-full gap-8">
	{#each systems.filter(system => searchValue === '' ? true : system.Title.toLowerCase() === searchValue.toLowerCase()) as system}
		{#await new Promise((r) => setTimeout(r, 500))}
			<LoadCard/>
		{:then _}
			<Card {system} cardPage={cardPage}/>
		{/await}
	{/each}
</div>