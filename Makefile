mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
base_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

# Golang project related settings
APP=$(base_dir)
APP_DIR=app
PROJECT_PACKAGE=github.com/lucassvieira/$(APP)
WORK_DIR=/go/src/$(PROJECT_PACKAGE)
BUILD_DIRECTORY=build
LINTER_EXECUTABLE := golangci-lint
LINTER_PATH := $(GOPATH)/bin/$(LINTER_EXECUTABLE)
BUILD_ENV :=
BUILD_ENV += CGO_ENABLED=0
SSH_KEY=~/.ssh/id_ed25519

# Docker 
DOCKER_REPOSITORY=localhost
DOCKER_IMAGE=$(APP)
DOCKER_TAG=1.0
DOCKER_NAMESPACE=lucassvieira
DOCKER_IMAGE_NAME=$(DOCKER_NAMESPACE)/$(DOCKER_IMAGE)

# Make commands
.PHONY: build
build:
	$(BUILD_ENV) go build -o build/$(APP) -a ./cmd/$(APP_DIR)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) . --build-arg APP=$(APP) \
	--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
	--build-arg WORK_DIR=$(WORK_DIR) \
	--build-arg SSH_KEY=$(SSH_KEY) \
	--build-arg APP_DIR=$(APP_DIR) 

proto:
	protoc --go_out=:. --go-grpc_out=:. internal/grpc/schema/*.proto

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