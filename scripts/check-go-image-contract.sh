#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $0 <image> <cgo-enabled: 0|1>" >&2
  exit 2
fi

image=$1
cgo_enabled=$2

docker run --rm "$image" sh -euc '
  test "$(go env GOVERSION)" = go1.27.0
  test "$(go env GOTOOLCHAIN)" = go1.27.0
  test "$(go env CGO_ENABLED)" = "$1"
  command -v bash
  command -v git
  command -v make
  if [ "$1" = 1 ]; then
    command -v gcc
    printf '\''package main\nimport "C"\nfunc main() {}\n'\'' >/tmp/cgo.go
    go build -o /dev/null /tmp/cgo.go
    rm /tmp/cgo.go
  else
    ! command -v gcc
  fi
' _ "$cgo_enabled"
