#!/bin/bash
mkdir -p output
GOOS=linux go build -o output/dump-tool dump-tool/cmd/main.go
docker build -t martynwin/sn-dump-collector-go:$1 -f docker/Dockerfile.go .

