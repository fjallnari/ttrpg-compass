<script lang="ts">
    import { createEventDispatcher } from 'svelte';

    export let options: string[] = [];
    export let title: string = '';
    export let value: string = '';

    const dispatch = createEventDispatcher();

</script>

<div class="flex flex-row justify-center items-center gap-2 w-36 lg:w-1/3">
    <h3>{title}</h3>
    <select bind:value 
        on:change={() => dispatch('search')} 
        class="block w-full p-2.5
        text-egshell text-sm border-b-abyss-900 border-solid border-b-2
        rounded-t bg-abyss-800 shadow backdrop-blur-md
        focus:ring-blue-500 focus:border-goldenrod focus:outline-none"
    >
        <option value={'*'}>any</option>
        {#each options.sort() as option}
            <option value={option}>{option.replaceAll('_', ' ')}</option>
        {/each}
    </select>
</div>