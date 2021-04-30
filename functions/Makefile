.PHONY: all test clean

build:
	mkdir -p build
	go get ./...
	go build -o ./build/ ./...

test:
	go test ./...

clean:
	rm -f build/*