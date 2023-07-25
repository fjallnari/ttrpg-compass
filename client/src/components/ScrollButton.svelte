<script lang="ts">
    import Icon from "@iconify/svelte";
    import { fade } from "svelte/transition";

    export let topElem: HTMLImageElement;

</script>

<button class="md:w-12 md:h-12 w-10 h-10 md:bottom-10 md:right-10 bottom-4 right-4 fixed flex justify-center items-center"
    transition:fade
    on:click={() => topElem.scrollIntoView({behavior: 'smooth'})}
>
    <svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linecap:square;stroke-miterlimit:1.5" viewBox="0 0 152 177">
        <path d="M115.089 46.01 76.68 84.419a1.133 1.133 0 0 1-1.6 0L36.67 46.01a1.134 1.134 0 0 1 0-1.601L75.08 6a1.133 1.133 0 0 1 1.6 0l38.409 38.409a1.132 1.132 0 0 1 0 1.601Zm-7.318-1.848a.922.922 0 0 0 0-1.302l-31.24-31.24a.922.922 0 0 0-1.302 0l-31.24 31.24a.922.922 0 0 0 0 1.302l31.24 31.24a.922.922 0 0 0 1.302 0l31.24-31.24Z" style="fill:#b69047"/>
        <path d="M5.408 75.672 75.88 5.408l70.263 70.264" style="fill:none;stroke:#b69047;stroke-width:6px"/>
        <path d="M75.88 85.219v87.297" style="fill:none;stroke:#b69047;stroke-width:6px;stroke-linejoin:bevel"/>
        <path d="M36.339 45.678v87.297M115.42 45.678v87.297" style="fill:none;stroke:#b69047;stroke-width:6px;stroke-linecap:butt;stroke-linejoin:bevel"/>
        <path d="M95.89 61.235a.58.58 0 0 0 0-.821L76.186 40.71a.58.58 0 0 0-.821 0L55.661 60.414a.58.58 0 0 0 0 .821l19.704 19.704a.58.58 0 0 0 .821 0L95.89 61.235Z" style="fill:#b69047"/>
    </svg>
    <!-- <svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2" viewBox="0 0 140 135">
        <path d="M139.909 66.646a1.87 1.87 0 0 1-.514 1.871l-67.988 65.236c-.782.75-2.051.75-2.833 0L.586 68.517a1.868 1.868 0 0 1 0-2.718L68.574.563c.254-.244.558-.408.881-.493.349-.093.72-.093 1.069-.001.323.085.629.25.883.494l67.988 65.236c.254.244.425.537.514.847ZM101.542 89.27l24.899-23.891a1.521 1.521 0 0 0 0-2.211l-55.298-53.06a1.68 1.68 0 0 0-2.305 0L13.54 63.168a1.521 1.521 0 0 0 0 2.211l24.715 23.714L69.08 59.515a1.06 1.06 0 0 1 1.453 0l31.009 29.755Z" style="fill:#b69047"/>
    </svg> -->
</button>