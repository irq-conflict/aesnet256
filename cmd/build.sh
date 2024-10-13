#!/bin/bash

echo "Building for amd64 Linux"
export GOOS=linux
go build -trimpath -ldflags="-s -w" -o aesnet256.amd64
echo "Building for arm64 linux"
echo "Building for amd64 Windows"
export GOOS=windows
go build -trimpath -ldflags="-s -w" -o aesnet256.exe
export GOOS=linux
export GOARCH=arm64
go build -trimpath -ldflags="-s -w" -o aesnet256.arm64
