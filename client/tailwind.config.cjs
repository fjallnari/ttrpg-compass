/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
	theme: {
		extend: {
			colors: {
				abyss: {
					900: '#111a20',
					800: '#172225',
					700: '#1C2B35',
					600: '#233643',
					500: '#2A4150',
					400: '#314C5E'
				},
				
				rich: '#12131C',
				raisin: '#242436',
				space: '#28293D',
				'off-gray': '#222227',
				goldenrod: '#B69047',
				gold: '#865903',
				eggshell: '#F4F1DE',
				'off-white': '#E5E7EB',
			},
			fontFamily: {
				italiana: ['Italiana', 'serif'],
				raleway: ['Raleway Variable', 'sans-serif'],
				cinzel: ['Cinzel Variable', 'serif'],
				'poiret-one': ['Poiret One', 'sans-serif'],
			},
			backgroundSize: {
				'auto-100%': 'auto 100%',
			  }
		},
	},
	plugins: [],
}
