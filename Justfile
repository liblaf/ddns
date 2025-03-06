default:

build:
  go build -o bin/ ./...

clean:
  @ rm --force --recursive --verbose bin/
