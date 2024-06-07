#!/bin/bash

# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o bin/har2sequence-linux-amd64 main.go

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o bin/har2sequence-linux-arm64 main.go

# macOS AMD64
GOOS=darwin GOARCH=amd64 go build -o bin/har2sequence-macos-amd64 main.go

# macOS ARM64
GOOS=darwin GOARCH=arm64 go build -o bin/har2sequence-macos-arm64 main.go

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -o bin/har2sequence-win-amd64.exe main.go

# Windows ARM64
GOOS=windows GOARCH=arm64 go build -o bin/har2sequence-win-arm64.exe main.go

echo "Build completed successfully."