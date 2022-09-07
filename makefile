build:
	@GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./bin/arenatv
	@go build -o docs/arenatv ./bin/arenatv

run: build
	@cd docs && ./arenatv
# 	@cd docs && ./arenatv local

# build-github: build
# 	@cd docs && ./arenatv github

clean:
	@go clean ./...
	@-rm docs/arenatv