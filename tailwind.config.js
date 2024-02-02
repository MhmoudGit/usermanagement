/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./templates/**/*.{html,js,templ}'],
	theme: {
		extend: {
			fontFamily: {
				Sixtyfour: ['Sixtyfour', 'sans-serif'],
			},
		},
	},
	plugins: [],
}
