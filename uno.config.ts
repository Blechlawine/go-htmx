import { defineConfig, presetUno } from "unocss";

export default defineConfig({
	content: {
		filesystem: ["views/**/*.html"],
	},
	cli: {
		entry: {
			patterns: ["views/**/*.html"],
			outFile: "static/uno.css",
		},
	},
	presets: [presetUno()],
});
