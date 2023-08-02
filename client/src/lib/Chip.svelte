<script lang="ts">
    export let value: string = '???';
    export let color: string = 'abyss-800';

</script>

<chip class='flex justify-center items-center flex-row gap-1 p-2 bg-{color} flex-wrap shadow-md rounded border-goldenrod border-solid border'>
    {value}
</chip>