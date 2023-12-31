BUILD_DIR ?= $(CURDIR)/build

build: go.sum
	@echo "--> Building..."
	go build -mod=readonly -o $(BUILD_DIR)/ ./...

generate:
	go generate ./....

.PHONY: build generate