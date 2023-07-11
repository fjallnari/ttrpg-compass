<script lang="ts">
	import Icon from '@iconify/svelte';

	//export let icon: string;
	//export let width: string = '2em';
	//export let height: string = width;

	export let systems: any[] = [];

</script>
<div class="flex flex-col flex-wrap justify-center items-center w-11/12 text-amber-50">
	<Icon icon="mdi:compass-rose" style="font-size: 64px;" />
	<div class="flex flex-row flex-wrap justify-center mt-6 h-full gap-4">
		{#each systems as system}
			<p class="w-1/4">
				{JSON.stringify(system)}
			</p>
		{/each}
	</div>
</div>