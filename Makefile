templ:
	@templ generate --watch --proxy="http://localhost:8080" --open-browser=false
tailwind:
	npx @tailwindcss/cli -i static/css/input.css -o static/css/styles.css --watch