mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
base_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

# Golang project related settings
APP=$(base_dir)
APP_DIR=app
PROJECT_PACKAGE=github.com/lucassvieira/$(APP)
BUILD_DIRECTORY=build
LINTER_EXECUTABLE := golangci-lint
LINTER_PATH := $(GOPATH)/bin/$(LINTER_EXECUTABLE)
BUILD_ENV :=
BUILD_ENV += CGO_ENABLED=0

# Docker 
DOCKER_REPOSITORY=localhost
DOCKER_IMAGE=$(APP)
DOCKER_NAMESPACE=local
DOCKER_REPOSITORY=$(DOCKER_REPOSITORY)/$(DOCKER_NAMESPACE)/$(DOCKER_IMAGE)

# Make commands
.PHONY: build
build:
	$(BUILD_ENV) go build -o build/$(APP) -a ./cmd/$(APP_DIR)

docker-build:
	docker build -t $(DOCKER_REPOSITORY):local . --build-arg APP=$(APP) --build-arg GITHUB_TOKEN=$(GITHUB_TOKEN)

proto:
	protoc --go-grpc_out=:. internal/grpc/schema/*.proto

.PHONY: lint
lint:
	$(LINTER_PATH) run ./...

# Utilitary commands
install-tools:
	go get github.com/google/wire/cmd/wire
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

.PHONY: init
init:
	find . -name '*.keep' | xargs rm
	rm -f go.mod go.sum
	go mod init $(PROJECT_PACKAGE)
	go mod tidy
	make install-tools