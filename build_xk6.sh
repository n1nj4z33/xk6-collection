#!/usr/bin/env bash

set -o errexit

here="$(realpath "$(dirname "$0")")"

export GOPATH="${GOPATH:-$(go env GOPATH)}"
export GOBIN="${GOPATH}"/bin
export PATH="${GOBIN}:${PATH}"
export GOPROXY=""

go install go.k6.io/xk6/cmd/xk6@v0.8.1

build_version="$(date -u +"%FT%T%z")/$(git describe --tags --always --long --dirty)"
export XK6_BUILD_FLAGS="-ldflags='-w -s -X go.k6.io/k6/lib/consts.VersionDetails=${build_version}'"
export K6_VERSION=v0.42.0

xk6 build \
    --with xk6-collection="${here}"/collection \
    --output "${here}"/build/k6 \
    "$@"
