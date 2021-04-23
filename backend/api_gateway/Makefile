APP=api_gateway

.PHONY: test
test: clean
	go test ./...

.PHONY: test-without-io
test-without-io: clean
	go test ./...

.PHONY: clean
clean:
	go clean