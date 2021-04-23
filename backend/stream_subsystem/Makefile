APP=stream_subsystem

.PHONY: test
test: clean
	go test ./...

.PHONY: test-without-io
test-without-io: clean
	go test ./inmem/.
	go test ./jwt/.
	go test ./stream/.
	go test .

.PHONY: clean
clean:
	go clean