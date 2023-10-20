build:
	#todo: impl build steps
test:
	#todo: impl test steps
watch:
	air
deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go get -u github.com/gofiber/fiber/v2 \
		github.com/fasthttp/router \
		github.com/a-h/templ \
		github.com/markphelps/optional
codegen:
	templ generate
