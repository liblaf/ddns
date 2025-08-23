#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

GOARCH="$(go env GOARCH)"
GOEXE="$(go env GOEXE)"
GOOS="$(go env GOOS)"

mkdir --parents --verbose 'dist'
go build -o "dist/ddns-$GOOS-$GOARCH$GOEXE" 'cmd/ddns/main.go'
