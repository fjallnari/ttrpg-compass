/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
	theme: {
		extend: {
			colors: {
				rich: '#12131C',
				raisin: '#242436',
				space: '#28293D',
				'off-gray': '#222227',
				goldenrod: '#B69047',
				eggshell: '#F4F1DE',
			},
			fontFamily: {
				italiana: ['Italiana', 'serif'],
				raleway: ['Raleway Variable', 'sans-serif'],
				cinzel: ['Cinzel Variable', 'serif'],
			}
		},
	},
	plugins: [],
}
