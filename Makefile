.PHONY: build
build:
	@go build ./...

coverage.out:
	@go test -race -cover -covermode=atomic -coverprofile ./$@ ./...

.PHONY: test
test: coverage.out

.PHONY: cover
cover: coverage.out
	@echo ""
	@go tool cover -func $<

builder: builder_stamp

builder_stamp: Dockerfile
	@docker build -t fc-task-driver-builder:latest .
	@touch $@

# For any given target, append "-in-docker" to it to run the build
# recipe in a container, e.g. instead of:
# $ make build
# you can use
# $ make build-in-docker
%-in-docker: builder_stamp
	docker run --rm \
		--mount type=bind,source=$(shell pwd),target=/src \
		fc-task-driver-builder:latest $(subst -in-docker,,$@)

