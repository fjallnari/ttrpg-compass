<script lang="ts">

	//export let icon: string;
	//export let width: string = '2em';
	//export let height: string = width;

	$: lala = "lala";

	const clickButton = async () => {
		console.log("Button clicked");
		lala = await fetch("http://localhost:5000/api/system/cairn").then(res => res.json());
	}

</script>

<button
	class="flex justify-center items-center w-full cursor-pointer transition ease-in duration-100 active:text-sky-300"
	on:click={() => clickButton()}
	on:keyup={() => {}}
>Get Cairn data</button>

<div>
	<p>{JSON.stringify(lala)}</p>
</div>