#!/bin/bash

set -o errexit
set -o nounset
set -o xtrace

export GOLANG="${HOME}/.go/go1.21.3"

PATH="${GOLANG}/bin:$PATH"

go build ./cmd/tsproxy
