build:
	#todo: impl build steps
	rm -f $(find -type f -regex '.*_templ.*')
	templ generate
	go build -o ./tmp/main ./src
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
		github.com/markphelps/optional \
		github.com/joho/godotenv \
		github.com/mattn/go-sqlite3 \
        github.com/libsql/libsql-client-go
codegen:
	templ generate
