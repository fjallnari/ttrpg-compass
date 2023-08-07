<script lang="ts">
	import Icon from "@iconify/svelte";
	import type TTRPGSystem from "../interfaces/TTRPGSystem";
	import Autocomplete from "./Autocomplete.svelte";
	import { slide } from "svelte/transition";

    export let systems: TTRPGSystem[] = [];
    export let genres: string[] = [];
    export let searchValue: string = "";

    let filtersMenuActive = false;
    let filters = {
        genre: ''
    };
    

</script>

<div class="flex flex-col items-center w-72 md:w-3/5 lg:w-1/3 
    p-2 border-b-[3px] border-transparent border-b-goldenrod border-solid 
    rounded-t bg-abyss-900 bg-opacity-75 
    focus-within:bg-opacity-50 focus-within:shadow-xl shadow backdrop-blur-md"
>
    <div class="flex items-center flex-row w-full gap-2">
        <button on:click={() => {}}>
            <Icon 
                icon="mdi:search" 
                class="text-2xl active:text-goldenrod"
            />
        </button>
        <Autocomplete bind:value={searchValue} systems={systems.map(sys => sys.Title)}/>
        <button on:click={() => filtersMenuActive = !filtersMenuActive}>
            <Icon 
                icon="mdi:filter-multiple" 
                class="text-xl active:text-goldenrod {filtersMenuActive ? "text-goldenrod": ""}"
            />
        </button>
    </div>
    {#if filtersMenuActive}
        <div transition:slide class="h-20 w-full m-2 p-2 font-poiret-one">
            Here be dragons
        </div>
    {/if}
</div>