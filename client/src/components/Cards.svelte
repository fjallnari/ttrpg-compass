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
	<div class=" flex justify-center items-center gap-4 text-opacity-90">
        {#each Array(3) as _, index}
            <button class="{navPage === index ? 'text-goldenrod': ''} cursor-pointer text-4xl" on:click={() => {navPage = index}}>
                <Icon icon="mdi:rhombus" />
            </button>
        {/each}
    </div>
	<div class="flex flex-row flex-wrap justify-center mt-6 h-full w-full gap-8">
		{#await new Promise((r) => setTimeout(r, 10000))}
			<LoadCard/>
		{:then _}
			<Card system={systems[0]} navPage={navPage}/>	
		{/await}
		{#each systems as system}
			<Card {system} navPage={navPage}/>
		{/each}
	</div>
</div>