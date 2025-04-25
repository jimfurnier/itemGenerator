#!/usr/bin/env bash

set -e

pushd "$(git rev-parse --show-toplevel)"

go mod tidy
go fmt ./...
go build -o build/generate cmd/main.go
