.PHONY: run clean vet fmt lint test cover ci

GO ?= go
GOLANGCI_LINT ?= golangci-lint
BUILD_DIR ?= build
BIN_DIR ?= $(BUILD_DIR)/bin
COVERPROFILE ?= $(BUILD_DIR)/coverage.out
COVERHTML ?= $(BUILD_DIR)/coverage.html
CUSTOM_GCL ?= $(BIN_DIR)/custom-gcl

vet:
	$(GO) vet ./...

fmt:
	@out=$$($(GOLANGCI_LINT) fmt --diff 2>&1); rc=$$?; \
	  if [ $$rc -ne 0 ] || [ -n "$$out" ]; then echo "$$out"; exit 1; fi

$(CUSTOM_GCL):
	mkdir -p $(BIN_DIR)
	$(GOLANGCI_LINT) custom

lint: $(CUSTOM_GCL)
	$(CUSTOM_GCL) run ./...

test:
	$(GO) test -race -count=1 ./...

cover:
	mkdir -p $(BUILD_DIR)
	$(GO) test -race -count=1 -coverprofile=$(COVERPROFILE) ./...
	$(GO) tool cover -html=$(COVERPROFILE) -o $(COVERHTML)

clean:
	rm -rf $(BUILD_DIR)

ci: vet fmt lint test
