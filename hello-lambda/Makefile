build-linux:
	mkdir -p out
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o out/main main.go

package-linux: build-linux
	zip -j out/hello-lambda_linux.zip out/main

build-docker-lambda:
	docker build -t hello-lambda .

clean:
	rm -rf out
