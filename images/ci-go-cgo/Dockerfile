FROM golang:1.27.0-alpine3.23@sha256:3747dcba41c8b0db3211fda4db61638b980e17ac5bb3c94460a975a9cfe19395

RUN apk add --no-cache bash build-base git make

ENV GOTOOLCHAIN=go1.27.0
ENV CGO_ENABLED=1

RUN test "$(go env GOVERSION)" = go1.27.0 && \
    test "$(go env GOTOOLCHAIN)" = go1.27.0 && \
    test "$(go env CGO_ENABLED)" = 1 && \
    printf 'package main\nimport "C"\nfunc main() {}\n' >/tmp/cgo.go && \
    go build -o /dev/null /tmp/cgo.go && \
    rm /tmp/cgo.go
