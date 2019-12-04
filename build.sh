#!/usr/bin/env bash

go mod download

env GOOS=linux GOARCH=amd64 go build -o hash-cracker_linux-amd64
env GOOS=windows GOARCH=amd64 go build -o hash-cracker_windows-amd64.exe
env GOOS=darwin GOARCH=amd64 go build -o hash-cracker_darwin-amd64
