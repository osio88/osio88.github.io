build:
# build wasm
	GOARCH=wasm GOOS=js go build -o web/app.wasm src/cmd/soop/main.go

	go build -o bin/soop src/cmd/soop/main.go

static:
	cp -r web/. docs/web/.

run:
	bin/soop

clean:
# remove old docs
	rm -rf docs
	rm -rf bin