#!/usr/bin/env sh
set -e
VERSION="${1:-v0.1.1}"
go build -ldflags "-s -w -X github.com/Ahmed9ali/Pascal/internal/version.Version=$VERSION" -o pascal .
echo "Built pascal $VERSION"
