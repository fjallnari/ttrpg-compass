<script lang="ts">
    import Icon from "@iconify/svelte";
    import { fade } from "svelte/transition";

    export let topElem: HTMLImageElement;

</script>

<button class="w-12 h-12 bottom-10 right-10 fixed flex justify-center items-center "
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
</button>