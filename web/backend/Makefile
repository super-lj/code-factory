.PHONY: build clean
.DEFAULT_GOAL := build

build: clean
	mkdir build
	go build -o ./build/main web-backend/cmd/bff

clean:
	rm -rf build