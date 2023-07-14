<script lang="ts">
	import Icon from '@iconify/svelte';
    import Card from './Card.svelte';
  import CardAlt from './CardAlt.svelte';

	//export let icon: string;
	//export let width: string = '2em';
	//export let height: string = width;

	export let systems: any[] = [];

</script>
<div class="flex flex-col flex-wrap justify-center items-center w-11/12 text-eggshell">
	<Icon icon="mdi:compass-rose" style="font-size: 64px;"/>
	<div class="flex flex-row flex-wrap justify-center mt-6 h-full w-full gap-8">
		{#each systems as system}
			<CardAlt {system}/>
		{/each}
	</div>
</div>