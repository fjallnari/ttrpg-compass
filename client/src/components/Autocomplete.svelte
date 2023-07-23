<script lang="ts">

    export let value = '';
    export let systems: string[] = [];

    let suggestion = '';
    let systemFound: string | undefined = undefined;

    $: {
        if (value === '') {
            suggestion = '';
        } else {
            systemFound = systems.find((system) => system.toLowerCase().startsWith(value.toLowerCase()));

            if (systemFound) {
                suggestion = `${value}${systemFound.slice(value.length)}`
            } else {
                suggestion = '';
            }
        }
    }

    const onKeypress = (e: any) => {
        if (e.key === 'Enter' && suggestion !== '') {
            value = suggestion;
            suggestion = '';
        }
    }

</script>

<div id="container" class="w-full font-poiret-one relative font-semibold">
    <div id="complete" class="absolute top-0 left-0 cursor-text text-[#999]">{suggestion}</div>
    <input bind:value on:keypress={onKeypress} type="search" id="input" class="bg-inherit border-none outline-none z-10 w-full" placeholder="Search TTRPG systems..." required>
</div>