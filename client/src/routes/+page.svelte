<script lang="ts">
	import Cards from "$lib/Cards.svelte";
	import type { PageData } from "./$types";

    export let data: PageData;


</script>

<main class="flex flex-col flex-wrap justify-center items-center w-11/12 p-6 m-auto">
    <Cards bind:systems={data.systems}/>
</main>