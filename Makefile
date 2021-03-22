.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o bin/kpl ./cmd

.PHONY: clean
clean:
	rm -r ./bin