<script lang="ts">
	import Icon from '@iconify/svelte';
  	import Card from './Card.svelte';
  	import LoadCard from './LoadCard.svelte';

	//export let icon: string;
	//export let width: string = '2em';
	//export let height: string = width;

	export let systems: any[] = [];
	let navPage = 0;

</script>
<div class="flex flex-col flex-wrap justify-center items-center w-11/12 text-eggshell gap-8">
	<Icon icon="mdi:compass-rose" class="text-goldenrod text-8xl"/>
	<div class="flex items-center w-4/5 md:w-3/5 lg:w-1/3 h-10 border-b-2 border-transparent border-b-goldenrod border-solid rounded-t bg-abyss-900 bg-opacity-75 hover:bg-opacity-50 backdrop-blur-md">
		<p class="p-2 font-italiana font-semibold">Search systems...</p>
	</div>
	<div class=" flex justify-center items-center gap-4 text-opacity-90">
        {#each Array(3) as _, index}
            <button class="{navPage === index ? 'text-goldenrod': ''} cursor-pointer text-4xl" on:click={() => {navPage = index}}>
                <Icon icon="mdi:rhombus" />
            </button>
        {/each}
    </div>
	<div class="flex flex-row flex-wrap justify-center mt-6 h-full w-full gap-8">
		{#each systems as system}
			{#await new Promise((r) => setTimeout(r, 1000))}
				<LoadCard/>
			{:then _}
				<Card {system} navPage={navPage}/>
			{/await}
		{/each}
	</div>
</div>