build:
	@GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./bin/arenatv
	@go build -o docs/arenatv ./bin/arenatv

run: build
	@cd docs && ./arenatv

clean:
	@go clean ./...
	@-rm docs/arenatv