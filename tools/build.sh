#!/bin/bash

set -o errexit
set -o nounset
set -o xtrace

cd "$(dirname "$0")/.."

go clean
go build ./cmd/tsproxy
