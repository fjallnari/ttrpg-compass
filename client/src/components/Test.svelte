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
<div class="flex flex-col justify-center items-center w-full">
	<button
		class="cursor-pointer transition ease-in duration-100 active:text-amber-500 text-amber-50 bg-sky-950 rounded-md drop-shadow-lg py-6 px-2 box-content"
		on:click={() => clickButton()}
		on:keyup={() => {}}
	>Get Cairn data</button>
	<div>
		<p>{JSON.stringify(lala)}</p>
	</div>
</div>