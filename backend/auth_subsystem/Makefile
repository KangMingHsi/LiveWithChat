APP=auth_subsystem

.PHONY: test
test: clean
	go test ./...

.PHONY: test-without-io
test-without-io: clean
	go test ./inmem/.
	go test ./jwt/.
	go test ./auth/.
	go test .

.PHONY: clean
clean:
	go clean

.PHONY: service-up
service-up:
	docker-compose up -d

.PHONY: service-down
service-down:
	docker-compose down