<script lang="ts">
    import type TTRPGSystem from "../interfaces/TTRPGSystem";

    export let system: TTRPGSystem;

</script>

<div class="max-w-md w-1/2 box-border h-max grid grid-cols-2 grid-rows-8 gap-2 border-2 border-white">
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

</div>