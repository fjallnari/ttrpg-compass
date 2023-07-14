<script lang="ts">
    import Icon from "@iconify/svelte";
    import type TTRPGSystem from "../interfaces/TTRPGSystem";

    export let system: TTRPGSystem;

</script>

<!-- <div class="max-w-md w-1/2 box-border h-max grid grid-cols-2 grid-rows-8 gap-2 border-2 border-white">
    <div class="text-center w-full col-span-2 col-start-1 row-start-1 p-2 whitespace-nowrap">
        <h1>{system.Title}</h1>
    </div>
    <div class="col-span-2 col-start-1 row-start-2 bg-amber-400 p-2 h-5">

    </div>
    <div class="row-span-4 col-start-1 row-start-3 bg-red-500 h-20 p-2 w-full">

    </div>
    <div class="row-span-4 col-start-2 row-start-3 bg-green-500 h-fit p-2 w-full">
        <p>{system.Description}</p>
    </div>
    <div class="col-span-2 row-span-2 col-start-1 row-start-7 bg-blue-500 p-2 h-5">

    </div>

</div> -->

<!-- <div class="max-w-md w-1/2 text-center box-border h-full border-2 border-white grid grid-flow-row grid-cols-6 grid-rows-7 gap-2">
    <div class="flex justify-center items-center p-2 whitespace-nowrap w-full col-span-3 row-span-2">
        <h1>{system.Title}</h1>
    </div>
    <div class="bg-amber-400 p-2 h-fit col-span-3 row-span-3 col-start-1 row-start-3">
        <p>{system.Description}</p>
    </div>
    <div class="bg-red-500 w-full col-span-3 row-span-2 col-start-1 row-start-6">12</div>
    <div class="flex flex-col flex-1 h-full bg-blue-500 col-span-3 row-span-7 col-start-4 row-start-1">
        <div>a</div>
    </div>
</div> -->
<div class="max-w-md w-1/2 text-center h-full shadow-lg rounded-md text-eggshell bg-abyss-600 border-4 border-double border-goldenrod">
    <div class="container">
        <div class="title flex justify-center items-center m-0 shadow-inner font-italiana font-bold text-2xl bg-abyss-800 border-r-2 border-b-2 border-goldenrod">
            <h1>{system.Title}</h1>
        </div>
        <div class="info flex flex-col justify-center items-center border-r-2 font-cinzel border-goldenrod shadow-inner">
            <p>{system.Type.toUpperCase()}</p>
            <p>{system.Gm.toUpperCase()}</p>
        </div>
        <div class="genre border-t-2 border-r-2 p-2 text-base font-cinzel font-semi-bold border-goldenrod bg-abyss-800 shadow-inner">
            <p>{system.Genre.toUpperCase()}</p>
        </div>

        <div class="metrics h-full flex flex-col justify-center items-center p-2 shadow-inner">
            <p class="italic">
                {system.Description}
            </p>
            <div class="flex justify-center items-center gap-2 mt-auto">
                <Icon icon="mdi:rhombus" class="text-goldenrod cursor-pointer"/>
                <Icon icon="mdi:rhombus" class="cursor-pointer"/>
            </div>
        </div>
    </div>
</div>


<style>
    .container {  display: grid;
        grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
        grid-template-rows: 1fr 1fr 1fr 1fr 1fr;
        gap: 0px 0px;
        grid-auto-flow: row;
        grid-template-areas:
            "title title title metrics metrics metrics"
            "title title title metrics metrics metrics"
            "info info info metrics metrics metrics"
            "info info info metrics metrics metrics"
            "genre genre genre metrics metrics metrics";
    }

    .title { grid-area: title; }

    .genre { grid-area: genre; }

    .info { grid-area: info; }

    .metrics { grid-area: metrics; }

</style>