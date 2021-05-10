APP=content_subsystem

.PHONY: test
test: clean
	go test ./...

.PHONY: test-without-io
test-without-io: clean
	go test .
	go test ./inmem/.
	go test ./content/.

.PHONY: clean
clean:
	go clean