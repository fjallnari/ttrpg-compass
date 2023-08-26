<script lang="ts">
	import Icon from "@iconify/svelte";
	import type TTRPGSystem from "../interfaces/TTRPGSystem";
	import Autocomplete from "./Autocomplete.svelte";
	import { slide } from "svelte/transition";
	import { foundSystems, selectedSystem, serverURL } from "../stores";
	import FilterSelect from "./FilterSelect.svelte";
	import type Filters from "../interfaces/Filters";

    export let genres: string[] = [];
    export let families: string[] = [];
    export let searchValue: string = "";
    export let suggestion: string = "";
    export let filters: Filters;

    let filtersMenuActive = false;

    const searchSystems = async () => {
        const sanitizedSearchValue = searchValue === "" ? "*" : searchValue.replace(' ', '_');
        const sanitizedFamily = filters.family === "*" ? "*" : filters.family.replace(' ', '_');
        const res = await fetch(`${$serverURL}/api/systems/search/title:${sanitizedSearchValue}/genre:${filters.genre}/family:${sanitizedFamily}`);
        
        if (res.status === 404) {
            //console.log('no systems found');
            //foundSystems.set([]);
            return;
        }

        let data = await res.json() as TTRPGSystem[];

        if (data.length === 0) {
            filters = {
                genre: '*',
                family: '*',
            };
            suggestion = '';
            foundSystems.set([]);
            return;
        }

        if (data && sanitizedSearchValue != "*") {
            suggestion = `${searchValue}${data[0].Title.slice(searchValue.length)}`;
        } else {
            suggestion = '';
        }

        selectedSystem.set(undefined);
        foundSystems.set(data);
    }
    

</script>

<div class="flex flex-col items-center w-72 md:w-3/5 lg:w-1/3 
    p-2 border-b-[3px] border-transparent border-b-goldenrod border-solid 
    rounded-t bg-abyss-900 bg-opacity-75 
    focus-within:bg-opacity-50 focus-within:shadow-xl shadow backdrop-blur-md"
>
    <div class="flex items-center flex-row w-full gap-2">
        <Icon 
            icon="mdi:search" 
            class="text-xl"
        />
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
        <div transition:slide class="flex flex-row flex-wrap justify-center items-center gap-5 h-fit w-full m-2 p-2 font-poiret-one">
            <FilterSelect 
                bind:value={filters.genre} 
                title="Genre:" 
                options={genres}
                on:search={searchSystems}
            />

            <FilterSelect 
                bind:value={filters.family} 
                title="Family:" 
                options={families}
                on:search={searchSystems}
            />
        </div>
    {/if}
</div>