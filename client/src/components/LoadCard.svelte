<script lang="ts">
    import { Rive } from "@rive-app/canvas";
    import { onDestroy, onMount } from "svelte";

    let canvas: HTMLCanvasElement, riveInstance: Rive;

    onMount(() => {
        riveInstance = new Rive({
            // Hosted .riv asset, or a local one. Uncomment line 8 to try with
            // a local basketball.riv asset, or add your own!
            //src: "https://cdn.rive.app/animations/vehicles.riv",
            src: "../src/assets/load_circle.riv",
            canvas: canvas,
            stateMachines: "load",
            autoplay: true,
            onLoad: () => {
                riveInstance.resizeDrawingSurfaceToCanvas();
            }
        });
    });

    onDestroy(() => {
        riveInstance?.cleanup();
    });

</script>

<div class="flex justify-center items-center shadow-lg overflow-hidden bg-abyss-900 bg-opacity-75 backdrop-blur-md rounded-md h-[28rem] w-[22rem]">
    <canvas bind:this={canvas} class="w-1/2 h-1/2"></canvas>
</div>