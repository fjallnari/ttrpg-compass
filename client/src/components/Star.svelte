<script lang="ts">
  	import type TTRPGSystem from "../interfaces/TTRPGSystem";

	export let system: TTRPGSystem;

	type CoordMap = Record<number, string>;

	let N: CoordMap = {
		1: "82.307l5.182 5.182",
		2: "75.331l5.182 12.158",
		3: "61.3l5.182 26.189",
		4: "46.165l5.182 41.324",
		5: "27.258l5.182 60.231"
	}

	let NE: CoordMap = {
		1: "h7.329v7.329l",
		2: " 11.996-4.667-4.667 11.996 ",
		3: " 22.996-15.667-15.667 22.996 ",
		4: " 33.996-26.667-26.667 33.996 ",
		5: " 44.996-37.667-37.667 44.996 "
	}

	let E: CoordMap = {
		1: "5.182 5.182-5.182",
		2: "12 5.182-12",
		3: "26 5.182-26",
		4: "41 5.182-41",
		5: "60 5.182-60",
	}

	let SE: CoordMap = {
		1: "v7.329h-7.329L100",
		2: " 5.645 12.974-12.974-5.645-5.182 5.182",
		3: " 20.111 27.44-27.44-20.111-5.182 5.182"
	}

	let S: CoordMap = {
		1: "117.693l-5.182-5.182",
		2: "124.145l-5.182-11.634",
		3: "138.145l-5.182-25.634",
	}

	let SW: CoordMap = {
		1: "h-7.329v-7.329",
		2: "-26.076 18.747 18.747-26.076"
	}

	let W: CoordMap = {
		1: "L82.307 100l5.182",
		2: "L82.307 100l5.182"
	}

	let NW: CoordMap = {
		1: "v-7.329h7.329",
		2: " 59.197 59.197l35.621 28.292",
	}
	
</script>

<svg viewBox="-25 -40 250 250" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2">
	<path style="fill:none" d="M0 0h200v200H0z"/>
	<path d="M87.489 94.818{NW[2]}L100 {N[system.Complexity]}{NE[system.Progression]}{E[system.Balance]} 5.182{SE[1]} {S[3]}{SW[2]}{W[2]}-5.182Zm19.466 5.324c.079-.078.079-.206 0-.284l-6.813-6.813c-.078-.079-.206-.079-.284 0l-6.813 6.813c-.079.078-.079.206 0 .284l6.813 6.813c.078.079.206.079.284 0l6.813-6.813Z" style="fill:#b69047"/>
</svg>

