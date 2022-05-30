mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
base_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

# Golang project related settings
APP=$(base_dir)
PROJECT_PACKAGE=github.com/lucassvieira/$(APP)
BUILD_DIRECTORY=build
LINTER_EXECUTABLE := golangci-lint
LINTER_PATH := $(GOPATH)/bin/$(LINTER_EXECUTABLE)

# Make commands
.PHONY: build
build:
	go build ${GOARGS} --tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIRECTORY}/app ./cmd/app

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
	go mod init $(PROJECT_PACKAGE)
	go mod tidy
	make install-tools
