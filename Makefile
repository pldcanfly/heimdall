build:
	go build -o bin/heimdall

run: build
	./bin/heimdall
