<script lang="ts">
    import Icon from "@iconify/svelte";
    import type TTRPGSystem from "../interfaces/TTRPGSystem";
    import StellarChart from "./StellarChart.svelte";
    import StellarIcons from "./StellarIcons.svelte";
    import { int2roman } from "../util/util.ts"

    export let system: TTRPGSystem;

    export let cardPage = 0;

    let trackedMetric = '' as keyof TTRPGSystem;

</script>

<!-- bg-opacity-95 bg-blend-color-burn bg-auto-100% bg-repeat-round bg-[url('../src/assets/marble_texture.jpg')] -->
<div class="card shadow-lg overflow-hidden bg-abyss-900 bg-opacity-75 backdrop-blur-md rounded h-[29rem] w-[22rem]">
    <div class="title flex flex-col justify-center items-center gap-2 text-center text-opacity-90">
        <h1 class="font-cinzel text-xl pt-2 px-2">
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
    <div class="content flex justify-center items-center text-center min-h-0 relative">
        {#if cardPage === 0}
            <div class="flex flex-col justify-evenly items-center gap-8">
<!--                 <div class="flex justify-center items-center gap-4 font-poiret-one font-semibold text-lg">
                    <p>{system.Gm}</p>
                </div> -->
                <p class="font-poiret-one text-lg px-4 py-2">
                    {system.Description}
                </p>
                <!-- <div class="flex justify-center items-center gap-4 font-poiret-one font-bold text-lg">
                    <p>{system.Type.toUpperCase()}</p>
                </div> -->
            </div>
        {:else if cardPage === 1}
            <StellarChart {system} config={{ wantIcons: false, wantRadial: true }} />
            <StellarIcons bind:trackedMetric />
        {:else}
            <div class="flex flex-col justify-center items-center gap-4 px-4 py-2">
                <div class="flex flex-col justify-center items-center gap-1">
                    <h3 class="font-cinzel text-xl font-medium">GM?</h3>
                    <p class="font-poiret-one text-lg">
                        {system.Gm}
                    </p>
                </div>
                <div class="flex flex-col justify-center items-center gap-1">
                    <h3 class="font-cinzel text-xl font-medium">Similar systems</h3>
                    <p class="font-poiret-one text-lg">
                        {system.Similar.join(', ')}
                    </p>
                </div>
                <div class="flex flex-col justify-center items-center gap-1">
                    <h3 class="font-cinzel text-xl font-medium">URL</h3>
                    <a class="font-poiret-one text-lg active:text-goldenrod transition-colors" href="{system.Url}">{system.Url}</a>
                </div>
            </div>
        {/if}
    </div>
    <div class="nav flex justify-center items-center gap-4 text-opacity-90 text-2xl z-10">
        {#each Array(3) as _, index}
            <button class="{cardPage === index ? 'text-goldenrod': ''} cursor-pointer transition-colors" on:click={() => {cardPage = index}}>
                <Icon icon="mdi:rhombus" />
            </button>
        {/each}
    </div>
    <div class="genre bg-transparent bg-opacity-60 flex justify-center items-center">
        <h1 class="font-cinzel text-xl text-opacity-90 ">
            {#if cardPage === 1}
                {trackedMetric ? `${trackedMetric}: ${int2roman(~~system[trackedMetric]) ?? ''}` : 'hover/click aspect'}
            {:else if cardPage === 2}
                {system.Type.toUpperCase()}
            {:else}
                {system.Genre.toUpperCase()}
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