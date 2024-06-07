#!/bin/bash

# Define the output directories
output_dirs=(
  "bin/linux/amd64"
  "bin/linux/arm64"
  "bin/win/amd64"
  "bin/win/arm64"
  "bin/macos/amd64"
  "bin/macos/arm64"
)

# Create the output directories if they don't exist
for dir in "${output_dirs[@]}"; do
  mkdir -p "$dir"
done

# Build for 64-bit Linux
GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/har2sequence

# Build for ARM64 Linux
GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/har2sequence

# Build for 64-bit Windows
GOOS=windows GOARCH=amd64 go build -o bin/win/amd64/har2sequence.exe

# Build for ARM64 Windows
GOOS=windows GOARCH=arm64 go build -o bin/win/arm64/har2sequence.exe

# Build for 64-bit macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o bin/macos/amd64/har2sequence

# Build for ARM64 macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o bin/macos/arm64/har2sequence

echo "Build completed successfully."