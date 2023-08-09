<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import { foundSystems } from "../stores";

    export let value = '';
    export let suggestion = '';

    const dispatch = createEventDispatcher();

    $: {
        if (value === '') {
            suggestion = '';
            foundSystems.set([]);
        } else if (value.length >= 2) {
            dispatch('search');
        }
    }

    const onKeypress = (e: any) => {
        if (e.key === 'Enter' && suggestion !== '') {
            value = suggestion;
            suggestion = '';
        }
    }

</script>

<div id="container" class="w-full italic font-poiret-one relative text-gray-300 text-lg">
    <div id="complete" class="absolute top-0 left-0 cursor-text text-[#999] select-none">{suggestion}</div>
    <input bind:value on:keypress={onKeypress} type="search" id="input" class="bg-inherit border-none outline-none z-10 w-full italic" placeholder="Search TTRPG systems..." required>
</div>