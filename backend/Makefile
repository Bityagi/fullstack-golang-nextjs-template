PROJECT_NAME                := platform-microservice-boilerplate
OPENAPI_GENERATOR_VERSION   := 6.3.0
SPEC                        := ./spec/openapi/openapi-v1.0.yaml
SPEC_TEMPLATE               := ./spec/openapi/openapi.yaml
SPECTRAL_RULES              := ./spec/openapi/.spectral.yaml
GOLANG_CI     				:= /golangci.yaml
ROUTER                      := chi
VERSION                     := v1.0
GEN_CODE_DIR                := build/generated
BINARY_DIR                  := ./build/bin
REPORT_DIR                  := ./build/reports
GIT_HOST                    := build
GIT_USER                    := code
GIT_REPO                    := spec

.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: server
server: ## Run the Application
	go run ./src/main.go

##@ Open API Spec Commands
lint-spec: ## Lint your spec
	@echo "Linting the spec found at " $(SPEC)
	docker run --rm -v  ${PWD}:/local stoplight/spectral lint /local/$(SPEC) --ruleset /local/$(SPECTRAL_RULES)

mock-data: ## Mock data using your spec
	docker run --init --rm -v $(PWD):/local -p 4010:4010 stoplight/prism:4 mock -h 0.0.0.0 /local/$(SPEC) -d true

##@ Code Cleanliness
lint-code: lint-src-code lint-build-code

lint-src-code: ## Lint your code in the SRC Folder
	docker run --rm -v $(PWD):/local \
 		-w /local golangci/golangci-lint:latest \
 		golangci-lint run --fix --config /local$(GOLANG_CI) /local/src

lint-build-code: ## Lint the generated Code in the Build Folder
	docker run --rm -v $(PWD):/local \
		-w /local golangci/golangci-lint:latest \
		golangci-lint run --fix --config /local$(GOLANG_CI) ./local/$(GEN_CODE_DIR)/src

##@ Auto-Generate Code using Spec
openApiGenerate: bundle-spec ## Generate controllers and modelers using OpenAPI Spec
	@echo "Generating controllers and models using spec " $(SPEC)
	@echo "Generated code can be found at " $(GEN_CODE_DIR)
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-i /local/$(SPEC) \
		-g go-server \
		-o /local/$(GEN_CODE_DIR) \
		--git-host $(GIT_HOST) --git-user-id $(GIT_USER) --git-repo-id $(GIT_REPO) \
		--global-property=apiDocs=true,modelDocs=true,router=$(ROUTER) \
		--additional-properties=router=$(ROUTER),featureCORS=true,packageVersion=$(VERSION),sourceFolder=src,onlyInterfaces=true,outputAsLibrary=false

##@ Bundle Spec
bundle-spec: ## Bundle your spec
	@echo "Bundling the spec found at " $(SPEC)
	docker run -v $(PWD):/project wework/speccy \
		resolve $(SPEC_TEMPLATE) -o $(SPEC)

##@ Test Coverage
coverage: ## Test Coverage
	mkdir -p $(REPORT_DIR)
	cd ./src && go test -coverprofile=$(REPORT_DIR)/coverage.out -coverpkg=./... ./... && \
		go tool cover -html=$(REPORT_DIR)/coverage.out -o $(REPORT_DIR)/coverage.html

##@ Git Best Practices
install-git-hooks: ## Install Githooks
	@echo "Installing Git hooks"
