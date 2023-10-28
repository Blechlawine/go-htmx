import { defineConfig, presetUno } from "unocss";

export default defineConfig({
	content: {
		filesystem: ["src/views/**/*.templ"],
	},
	cli: {
		entry: {
			patterns: ["src/views/**/*.templ"],
			outFile: "static/uno.css",
		},
	},
	presets: [presetUno()],
});
