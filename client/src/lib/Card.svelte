<script lang="ts">
    import Icon from "@iconify/svelte";
    import type TTRPGSystem from "../interfaces/TTRPGSystem";
    import StellarChart from "./StellarChart.svelte";
    import StellarIcons from "./StellarIcons.svelte";
    import { ASPECTS, int2roman } from "../util/util.ts"
	import { foundSystems, selectedSystem } from "../stores.ts";

    export let system: TTRPGSystem;

    export let cardPage = 0;

    let trackedMetric = '' as keyof TTRPGSystem;

    const api = 'http://localhost:5000/api/systems/similar';

    const getSimilarSystems = async () => {
        const res = await fetch(`${api}/${system.Id}`);
        
        if (res.status === 404) {
            return;
        }

        let data = await res.json() as TTRPGSystem[];

        selectedSystem.set(system);
        foundSystems.set([system, ...data]);
    }

    $: isNotSelected = $selectedSystem && $selectedSystem.Id != system.Id;

</script>

<!-- bg-opacity-95 bg-blend-color-burn bg-auto-100% bg-repeat-round bg-[url('../src/assets/marble_texture.jpg')] -->
<div class="card shadow-lg overflow-hidden bg-abyss-900 bg-opacity-75 backdrop-blur-md rounded h-[29rem] w-[22rem]">
    <div class="title flex flex-col justify-center items-center gap-2 text-center text-opacity-90">
        <h1 class="font-cinzel text-xl pt-2 px-2 {$selectedSystem?.Id === system.Id ? "text-goldenrod": ""}">
            {`${system.Title}`}
        </h1>
        {#if system.Edition && system.Edition !== ''}
            <h2 class="font-poiret-one italic text-gray-300 text-lg pb-2">
                {`${system.Edition}`}
            </h2>
        {/if}
    </div>
    <div class="delim flex justify-center items-center ">
        <img src="/delim_deco.svg" class="w-4/5" alt="line">
    </div>
    <div class="content flex justify-center items-center text-center {cardPage === 3 ? 'overflow-y-auto': ''} min-h-0 relative">
        {#if cardPage === 0}
            <div class="flex flex-col justify-evenly items-center gap-8">
                <p class="font-poiret-one text-lg px-4 py-2">
                    {system.Description}
                </p>
            </div>
        {:else if cardPage === 1}
            <StellarChart {system} config={{ wantIcons: false, wantRadial: true }} />
            <StellarIcons bind:trackedMetric />
        {:else if cardPage === 2}
            <div class="flex flex-col justify-center items-center gap-4 px-4 py-2">
                <div class="flex flex-col justify-center items-center gap-1">
                    <h3 class="font-cinzel text-xl font-medium">Family</h3>
                    <p class="font-poiret-one text-lg">
                        {system.Family.replaceAll('_', ' ')}
                    </p>
                </div>
                <div class="flex flex-col justify-center items-center gap-1">
                    <h3 class="font-cinzel text-xl font-medium">GM</h3>
                    <p class="font-poiret-one text-lg">
                        {system.Gm}
                    </p>
                </div>
                <div class="flex flex-col justify-center items-center gap-1">
                    <h3 class="font-cinzel text-xl font-medium">Publisher</h3>
                    <p class="font-poiret-one text-lg">
                        {system.Publisher}
                    </p>
                    <!-- <a class="font-poiret-one text-lg active:text-goldenrod transition-colors" href="{system.Url}">{system.Url}</a> -->
                </div>
            </div>
        {:else if cardPage === 3}
            <div class="flex flex-col justify-center {isNotSelected ? 'self-start': ''} gap-1 px-2">
                <div class="flex flex-col justify-center items-center gap-1 p-2">
                    <div class="flex flex-row justify-center items-center gap-2">
                        <h3 class="font-cinzel text-xl font-medium">Similar systems</h3>
                        <!--  -->
                        <button on:click={() => getSimilarSystems()}
                            class="active:text-goldenrod cursor-pointer transition-colors"
                        >
                            <Icon icon="mdi:compare" class="text-xl" />
                        </button>
                    </div>
                    {#if !$selectedSystem || $selectedSystem.Id == system.Id}
                        <p class="font-poiret-one text-lg">
                            {system.Similar.join(', ')}
                        </p>
                    {/if}
                </div>
                {#if isNotSelected}
                    {#if $selectedSystem?.Genre === system.Genre}
                        <div class="flex flex-row justify-center items-center gap-2">
                            <div class="text-goldenrod">
                                <Icon icon="mdi:plus-thick"/>
                            </div>
                            <div class="flex flex-row gap-1 text-lg">
                                <h3 class="font-cinzel text-goldenrod font-medium">same genre |</h3>
                                <h3 class="font-poiret-one">{system.Genre.replaceAll('_', ' ')}</h3>
                            </div>
                        </div>
                    {/if}
                    {#if $selectedSystem?.Family === system.Family}
                        <div class="flex flex-row justify-center items-center gap-2">
                            <div class="text-goldenrod">
                                <Icon icon="mdi:plus-thick"/>
                            </div>
                            <div class="flex flex-row gap-1 text-lg">
                                <h3 class="font-cinzel text-goldenrod font-medium">family |</h3>
                                <h3 class="font-poiret-one">{system.Family.replaceAll('_', ' ').toLowerCase()}</h3>
                            </div>
                        </div>
                    {/if}
                    {#each ASPECTS as aspect}
                        {#if $selectedSystem && $selectedSystem[aspect] === system[aspect]}
                            <div class="flex flex-row justify-center items-center gap-2">
                                <div class="text-goldenrod">
                                    <Icon icon="mdi:plus-thick"/>
                                </div>
                                <div class="flex flex-row gap-1 text-lg">
                                    <h3 class="font-cinzel text-goldenrod font-medium lowercase">{aspect} |</h3>
                                    <h3 class="font-poiret-one">{system[aspect]}</h3>
                                </div>
                            </div>
                        {/if}
                    {/each}
                {/if}
            </div>
        {/if}
    </div>
    <div class="nav flex justify-center items-center gap-4 text-opacity-90 text-2xl z-10">
        {#each Array(4) as _, index}
            <button class="{cardPage === index ? 'text-goldenrod': ''} cursor-pointer transition-colors" on:click={() => {cardPage = index}}>
                <Icon icon="mdi:rhombus" />
            </button>
        {/each}
    </div>
    <div class="genre bg-transparent bg-opacity-60 flex justify-center items-center">
        <h1 class="font-cinzel text-xl text-opacity-90 uppercase">
            {#if cardPage === 1}
                {trackedMetric ? `${trackedMetric}: ${int2roman(~~system[trackedMetric]) ?? ''}` : 'hover/click aspect'}
            {:else if cardPage === 2}
                details
            {:else if cardPage === 0}
                {system.Genre.replaceAll('_', ' ').toUpperCase()}
            {:else}
                similarity
            {/if}
        </h1>
    </div>
</div>

<style>
    .card {  display: grid;
        background-image: stretch;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        grid-template-rows: 1.25fr 0.25fr 6fr 1fr 1fr;
        gap: 0px 0px;
        grid-auto-flow: row;
        grid-template-areas:
            "title title title title"
            "delim delim delim delim"
            "content content content content"
            "nav nav nav nav"
            "genre genre genre genre";
    }

    .title { grid-area: title; }

    .content { grid-area: content; }

    .nav { grid-area: nav; }

    .genre { grid-area: genre; }

    .delim { grid-area: delim; }

</style>