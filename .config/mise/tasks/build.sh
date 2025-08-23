#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

mkdir --parents --verbose 'dist'
go build -o 'dist' ./...
