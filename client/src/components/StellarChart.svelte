<script lang="ts">
  	import type TTRPGSystem from "../interfaces/TTRPGSystem";

	export let system: TTRPGSystem;
	export let starColor: string = "#b69047";
	export let radialColor: string = "#5e6163";

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
		1: "v7.329h-7.329",
		2: " 9.545 16.874-16.874-9.545",
		3: "  17.627 24.956-24.956-17.627",
		4: " 27.403 34.732-34.732-27.403",
		5: " 43.197 50.526-50.526-43.197"
	}

	let S: CoordMap = {
		1: "117.693l-5.182-5.182",
		2: "131.114l-5.182-18.603",
		3: "141.969l-5.182-29.458",
		4: "156.441l-5.182-43.93",
		5: "177.378l-5.182-64.867"
	}

	let SW: CoordMap = {
		1: "h-7.329v-7.329",
		2: "-14 7 7-14",
		3: "-26.076 18.747 18.747-26.076",
		4: "-38 31 31-38",
		5: "-50 43 43-50",
	}

	let W: CoordMap = {
		1: "L82.307 100l5.182-5.182",
		2: "L68.9 100.015l18.589-5.197",
		3: "L58.031 100l29.458-5.182",
		4: "l-43.197-5.167 43.197-5.197",
		5: "l-65.023-5.167 65.023-5.197"
	}

	let NW: CoordMap = {
		1: "v-7.329h7.329",
		2: " 82 82l13 6",
		3: " 70 70l25 17.5",
		4: " 59.197 59.197l35.621 28.292",
		5: " 45 45l49.5 42.5",
	}
	
</script>
<!-- -25 -40 250 250 -->
<svg viewBox="-50 -60 300 300" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2">
	<path style="fill:none" d="M0 0h200v200H0z"/>
	<path d="m99.966 68.915 21.99 9.109 9.109 21.991-9.109 21.99-21.99 9.109-21.991-9.109-9.109-21.99 9.109-21.991 21.991-9.109Z" style="fill:none;stroke:{radialColor};stroke-width:.53px;stroke-dasharray:1.6,2.13,0,0"/>
  	<path d="m99.966 58.031 29.676 12.292L141.935 100l-12.293 29.677-29.676 12.292-29.677-12.292L57.997 100l12.292-29.677 29.677-12.292Z" style="fill:none;stroke:{radialColor};stroke-width:.53px;stroke-dasharray:1.6,2.13,0,0"/>
  	<path d="m99.966 44.292 39.391 16.317L155.674 100l-16.317 39.391-39.391 16.317-39.392-16.317L44.258 100l16.316-39.391 39.392-16.317Z" style="fill:none;stroke:{radialColor};stroke-width:.53px;stroke-dasharray:1.6,2.13,0,0"/>
  	<path d="m99.966 22.466 54.824 22.709L177.499 100l-22.709 54.825-54.824 22.709-54.825-22.709L22.432 100l22.709-54.825 54.825-22.709Z" style="fill:none;stroke:{radialColor};stroke-width:.53px;stroke-dasharray:1.6,2.13,0,0"/>
	<path d="M87.489 94.818{NW[system.Versatility]}L100 {N[system.Complexity]}{NE[system.Progression]}{E[system.Balance]} 5.182{SE[system.Narrative]}L100 {S[system.Combat]}{SW[system.Exploration]}{W[system.Customization]}Zm19.466 5.324c.079-.078.079-.206 0-.284l-6.813-6.813c-.078-.079-.206-.079-.284 0l-6.813 6.813c-.079.078-.079.206 0 .284l6.813 6.813c.078.079.206.079.284 0l6.813-6.813Z" style="fill:{starColor}"/>
</svg>
