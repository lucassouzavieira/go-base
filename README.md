[![test](https://github.com/lucassouzavieira/go-base/actions/workflows/tests.yml/badge.svg)](https://github.com/lucassouzavieira/go-base/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucassouzavieira/go-base)](https://goreportcard.com/report/github.com/lucassouzavieira/go-base)
# Standard Go Project Layout

## Overview

This is a basic layout for Go application projects. It's based on [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout/blob/master/README.md) recommendations.

## Make commands
- `make init`  
Inits the repository removing undesirable files and update dependencies. 

- `make build`  
Builds the application. Executables are put into `build` directory

- `make docker-build`  
Builds the docker image

- `make proto`  
Generate the protobuf stubs from proto definitions

- `make test`  
Run all project tests
