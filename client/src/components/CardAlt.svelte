<script lang="ts">
  import Icon from "@iconify/svelte";
    import type TTRPGSystem from "../interfaces/TTRPGSystem";

    export let system: TTRPGSystem;

</script>

<div class="card shadow-lg overflow-hidden bg-abyss-900 bg-opacity-95 bg-blend-color-burn bg-auto-100% bg-repeat-round bg-[url('../src/assets/marble_texture.jpg')] h-96 w-72">
    <div class="title flex justify-center items-center text-center text-opacity-90">
        <h1 class="font-italiana font-bold text-2xl p-2">
            {system.Title}
        </h1>
    </div>
    <div class="content flex justify-center items-center text-center">
        <p class="font-poiret-one text-lg px-4">
            {system.Description}
        </p>
    </div>
    <div class="nav flex justify-center items-center gap-4 text-opacity-90 text-xl">
        <Icon icon="mdi:rhombus" class="text-goldenrod cursor-pointer"/>
        <Icon icon="mdi:rhombus" class="cursor-pointer"/>
        <Icon icon="mdi:rhombus" class="cursor-pointer"/>
    </div>
    <div class="genre bg-abyss-800 shadow-inner bg-opacity-60 flex justify-center items-center">
        <h1 class="font-poiret-one text-xl text-opacity-90 font-bold">
            {system.Genre.toUpperCase()}
        </h1>
    </div>
</div>

<style>
    .card {  display: grid;
        background-image: stretch;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        grid-template-rows: 1.5fr 4fr 1fr 1fr;
        gap: 0px 0px;
        grid-auto-flow: row;
        grid-template-areas:
            "title title title title"
            "content content content content"
            "nav nav nav nav"
            "genre genre genre genre";
    }

    .title { grid-area: title; }

    .content { grid-area: content; }

    .nav { grid-area: nav; }

    .genre { grid-area: genre; }

</style>