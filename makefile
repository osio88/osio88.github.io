build:
	# remove old docs
	rm -rf docs

	# build wasm
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/soop/main.go

	go build -o bin/soop cmd/soop/main.go

run:
	bin/soop
	cp -r web/. docs/web/.