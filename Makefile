.PHONY: build run clean vet fmt lint lint-self test cover ci

GO ?= go
GOLANGCI_LINT ?= golangci-lint
BUILD_DIR ?= build
BIN_DIR ?= $(BUILD_DIR)/bin
COVERPROFILE ?= $(BUILD_DIR)/coverage.out
COVERHTML ?= $(BUILD_DIR)/coverage.html

build:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/bavovna-lint ./cmd/bavovna-lint
	$(GO) build -o $(BIN_DIR)/bavovna-lint-all ./cmd/bavovna-lint-all

vet:
	$(GO) vet ./...

fmt:
	@out=$$($(GOLANGCI_LINT) fmt --diff 2>&1); rc=$$?; \
	  if [ $$rc -ne 0 ] || [ -n "$$out" ]; then echo "$$out"; exit 1; fi

lint:
	$(GOLANGCI_LINT) run ./...

lint-self: build
	IGNORE="**/testdata/**,**/vendor/**" \
	  $(GO) vet -vettool=$(abspath $(BIN_DIR))/bavovna-lint-all ./...

test:
	$(GO) test -race -count=1 ./...

cover:
	mkdir -p $(BUILD_DIR)
	$(GO) test -race -count=1 -coverprofile=$(COVERPROFILE) ./...
	$(GO) tool cover -html=$(COVERPROFILE) -o $(COVERHTML)

clean:
	rm -rf $(BUILD_DIR)

ci: vet fmt lint lint-self test
