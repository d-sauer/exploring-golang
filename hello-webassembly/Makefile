build:
	go build

build-wasm:
	GOOS=js GOARCH=wasm go build -o ./bin/main.wasm
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./bin
	cp index.html ./bin

clean:
	rm -rf ./bin

run:
	go run main.go

test:
	go run main.go