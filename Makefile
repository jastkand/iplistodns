GO := go
APP := extract_dns
SRC := extract_dns.go
BIN_DIR := bin
BIN := $(BIN_DIR)/$(APP)
INPUT ?= ip-list.json

# Support: make run ip-list.json
ifeq (run,$(firstword $(MAKECMDGOALS)))
ifneq ($(word 2,$(MAKECMDGOALS)),)
INPUT := $(word 2,$(MAKECMDGOALS))
$(eval $(INPUT):;@:)
endif
endif

.PHONY: help run run-test run-main build fmt clean

help:
	@echo "Common commands:"
	@echo "  make run                # Run with INPUT (default: ip-list.json)"
	@echo "  make run ip-list.json   # Run using positional file argument"
	@echo "  make run-test           # Run with ip-list-test.json"
	@echo "  make run-main           # Run with ip-list.json"
	@echo "  make build              # Build binary to bin/extract_dns"
	@echo "  make fmt                # Format Go files"
	@echo "  make clean              # Remove build artifacts"
	@echo "  make run INPUT=...      # Override input file"

run:
	$(GO) run $(SRC) -input $(INPUT)

run-test:
	$(GO) run $(SRC) -input ip-list-test.json

run-main: ip-list.json
	$(GO) run $(SRC) -input ip-list.json

build:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN) $(SRC)

fmt:
	$(GO) fmt ./...

clean:
	rm -rf $(BIN_DIR)

ip-list.json:
	curl -o ip-list.json https://russia.iplist.opencck.org/?format=json
