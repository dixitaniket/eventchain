BUILD_DIR ?= $(CURDIR)
start:
	ignite chain serve --reset-once

update-abi:
	cd contracts && npx hardhat export-abi
	abigen --abi ${BUILD_DIR}/contracts/abi/contracts/TestEvent.sol/TestEvent.json --pkg event --type TestEvent --out ./observer/event/contract.go

.PHONY: start update-abi