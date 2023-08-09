<script lang="ts">
	import Icon from "@iconify/svelte";
	import type TTRPGSystem from "../interfaces/TTRPGSystem";
	import Autocomplete from "./Autocomplete.svelte";
	import { slide } from "svelte/transition";
	import { foundSystems } from "../stores";

    export let genres: string[] = [];
    export let searchValue: string = "";
    export let suggestion: string = "";

    const api = 'http://localhost:5000/api/systems/search';

    let filtersMenuActive = false;
    let filters = {
        genre: '*'
    };

    const searchSystems = async () => {
        const res = await fetch(`${api}/title:${searchValue}/genre:${filters.genre}`);
        if (res.status === 404) {
            return;
        }
        let data = await res.json() as TTRPGSystem[];

        if (data.length === 0) {
            suggestion = '';
            return;
        }

        if (data) {
            suggestion = `${searchValue}${data[0].Title.slice(searchValue.length)}`
        } else {
            suggestion = '';
        }
        foundSystems.set(data);
    }
    

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
        <Autocomplete 
            bind:value={searchValue}
            bind:suggestion
            on:search={searchSystems}
        />
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
            <select id="countries" bind:value={filters.genre} class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
                <option value={'*'}>any</option>
                {#each genres.sort() as genre}
                    <option value={genre}>{genre.replaceAll('_', ' ')}</option>
                {/each}
            </select>
        </div>
    {/if}
</div>