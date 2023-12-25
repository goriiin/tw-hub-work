import type { Config } from 'tailwindcss'

const config: Config = {
	content: [
		'./src/pages/**/*.{js,ts,jsx,tsx,mdx}',
		'./src/components/**/*.{js,ts,jsx,tsx,mdx}',
		'./src/app/**/*.{js,ts,jsx,tsx,mdx}',
	],
	theme: {
		extend: {
			backgroundColor: (theme) => ({
				...theme('colors'),
				text: '#dfdac4',
				background: '#252422',
				primary: '#EB5E28',
				secondary: '#413e3a',
				accent: '#F5AD8C',
			}),
			textColor: (theme) => ({
				...theme('colors'),
				text: '#dfdac4',
				background: '#252422',
				primary: '#EB5E28',
				secondary: '#413e3a',
				accent: '#F5AD8C',
			}),
		},
	},
	plugins: [],
}
export default config
