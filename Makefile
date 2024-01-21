build:
	go build -o ./bin/pumpkinx

run: build
	./bin/pumpkinx

test:
	go test -v ./...		
