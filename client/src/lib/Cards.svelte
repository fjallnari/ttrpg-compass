<script lang="ts">
  	import Card from './Card.svelte';
  	import LoadCard from './LoadCard.svelte';
  	import type TTRPGSystem from '../interfaces/TTRPGSystem';
	import { foundSystems } from '../stores';
	import type Filters from '../interfaces/Filters';

	export let systems: TTRPGSystem[] = [];
	export let cardPage = 0;
	export let filters: Filters;

	export let height = 0;

</script>

<div bind:clientHeight={height} class="flex flex-row flex-wrap justify-center mt-6 h-full w-full gap-8">
	{#each $foundSystems.length == 0 ? systems : $foundSystems as system}
		{#await new Promise((r) => setTimeout(r, 400))}
			<LoadCard/>
		{:then _}
			<Card {system} cardPage={cardPage} bind:filters/>
		{/await}
	{/each}
</div>