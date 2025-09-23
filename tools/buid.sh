#!/bin/bash

set -o errexit
set -o nounset
set -o xtrace

export GOLANG="${HOME}/.go/go1.22.0"

PATH="${GOLANG}/bin:$PATH"

go clean
go get tailscale.com/tsnet
go build ./cmd/tsproxy
