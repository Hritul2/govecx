.PHONY: test bench vet fmt lint example cover

test:
	go test -race ./...

bench:
	go test -bench=. -benchmem -run=^$$ ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w .

lint:
	staticcheck ./...

example:
	go run ./examples/basic

cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
