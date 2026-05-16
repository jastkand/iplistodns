GO := go
APP := extract_dns
SRC := extract_dns.go
BIN_DIR := bin
BIN := $(BIN_DIR)/$(APP)

.PHONY: help test build fmt clean refresh

help:
	@echo "Common commands:"
	@echo "  make test               # Run Go tests"
	@echo "  make build              # Build binary to bin/extract_dns"
	@echo "  make fmt                # Format Go files"
	@echo "  make clean              # Remove build artifacts"
	@echo "  make refresh            # Build, download IP list, and parse to parsed.txt"

test:
	$(GO) test ./...

build:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN) $(SRC)

fmt:
	$(GO) fmt ./...

clean:
	rm -rf $(BIN_DIR)

ip-list.json:
	curl -o ip-list.json https://russia.iplist.opencck.org/?format=json

refresh: build ip-list.json
	$(BIN) -input ip-list.json > parsed.txt
