BUILD_TARGET?=./app

ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
	export GOPATH
endif
    
.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: install build ## Install dependencies & build

.PHONY: run
run: build ## Build & run application
	@echo "Starting..."
	./${BUILD_TARGET}

.PHONY: install
install: ## Install dependencies
	@echo "Fetching dependencies..."
	@go mod download

.PHONY: install-air
install-air: ## Install Air
	@echo "Installing Air..."
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b ${GOPATH}/bin

.PHONY: build
build: ## Build Application
	@echo "Building binary..."
	@go build -ldflags "-s -w" -o ${BUILD_TARGET} ./cmd/app/start.go

.PHONY: dev
dev: ## Running development
	@echo "Running..."
	@go run cmd/quick-start.go

.PHONY: air
air: ## Running development with hot reload using Air
	@exec "${GOPATH}/bin/air"

.PHONY: proto
proto: ## Generate grpc pb from .proto
	@rm -rf api/grpc/*.go
	@protoc -I=./proto --go_out=api/grpc/ --go_opt=paths=source_relative \
         --go-grpc_out=api/grpc/ --go-grpc_opt=paths=source_relative \
         --proto_path=proto \
         proto/*.proto