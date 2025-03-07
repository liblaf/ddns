export GOARCH := `go env GOARCH`
export GOEXE := `go env GOEXE`
export GOOS := `go env GOOS`

default:

build:
    go build -o bin/ ./...

dist:
    go build -o "./bin/ddns-{{ GOARCH }}-{{ GOOS }}{{ GOEXE }}" ./cmd/ddns/main.go

clean:
    @ rm --force --recursive --verbose bin/
