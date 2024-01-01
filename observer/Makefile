BUILD_DIR ?= $(CURDIR)/build

build: go.sum
	@echo "--> Building..."
	go build -mod=readonly -o $(BUILD_DIR)/ ./...

start:
	${MAKE} build
	echo '\n'|./build/observer ./config.toml


.PHONY: build start