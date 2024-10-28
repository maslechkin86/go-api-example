CMD	    = ./cmd
GO          = go

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

$(BIN):
	@mkdir -p $@

#Tools
mockery_version=v2.43.2
ginkgo_version=v2.19.0
oapicodegen_version=v2.4.1

.PHONY: tools-install
tools-install: ; $(info $(M) installing tools...) @ ## Install source linting / generation / test tools
	$Q $(GO) install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(oapicodegen_version)
	$Q $(GO) install github.com/vektra/mockery/v2@$(mockery_version)
	$Q $(GO) install github.com/onsi/ginkgo/v2/ginkgo@$(ginkgo_version)

.PHONY: gen
gen: tools-install
gen: ; $(info $(M) generating source files...) @ ## Run all source code generation tools
	$Q rm -rf $(shell find . -type d -name "mocks") && $(GO) generate ./...

# Run tests
.PHONY: test
test: ; @ ## Run tests
	@echo "Running tests..."
	go test -v ./...

.PHONY: go-run
go-run: ; @ ## Run the binary
	@echo " > Running..."
	go run -tags development ./cmd/user-app


.PHONY: help
help:
	$Q grep -hE '^[ a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'