<script lang="ts">
  import Icon from "@iconify/svelte";
    import type TTRPGSystem from "../interfaces/TTRPGSystem";
  import Chip from "./Chip.svelte";

    export let system: TTRPGSystem;

    export let navPage = 0;

</script>

<div class="card shadow-lg overflow-hidden bg-abyss-900 bg-opacity-95 bg-blend-color-burn bg-auto-100% bg-repeat-round bg-[url('../src/assets/marble_texture.jpg')] h-96 w-72">
    <div class="title flex justify-center items-center text-center text-opacity-90">
        <h1 class="font-italiana font-bold text-2xl p-2">
            {`${system.Title} ${system.Edition}`}
        </h1>
    </div>
    <div class="delim flex justify-center items-center ">
        <img src="../src/assets/delim_deco.svg" class="w-4/5" alt="line">
    </div>
    <div class="content flex justify-center items-center text-center">
        {#if navPage === 0}
            <p class="font-poiret-one text-lg px-4 py-2">
                {system.Description}
            </p>
        {:else if navPage === 1}
            <p class="font-poiret-one text-lg px-4 py-2">
                
            </p>
        {:else}
            <div class="flex flex-col justify-center items-center gap-2">
                {#each [system.Type, system.Url, system.Gm] as attrib}
                    <p class="font-poiret-one text-lg">
                        {attrib}
                    </p>
                {/each}
            </div>
        {/if}
    </div>
    <div class="nav flex justify-center items-center gap-4 text-opacity-90 text-xl">
        {#each Array(3) as _, index}
            <button class="{navPage === index ? 'text-goldenrod': ''} cursor-pointer" on:click={() => {navPage = index}}>
                <Icon icon="mdi:rhombus" />
            </button>
        {/each}
    </div>
    <div class="genre bg-transparent bg-opacity-60 flex justify-center items-center">
        <h1 class="font-poiret-one text-xl text-opacity-90 font-bold">
            {system.Genre.toUpperCase()}
        </h1>
    </div>
</div>

<style>
    .card {  display: grid;
        background-image: stretch;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        grid-template-rows: 1.25fr 0.25fr 4fr 1fr 1fr;
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