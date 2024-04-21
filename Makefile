.PHONY: build start clean

build:
	sam build --no-cached

start: build
	sam local start-api

clean:
	go mod tidy