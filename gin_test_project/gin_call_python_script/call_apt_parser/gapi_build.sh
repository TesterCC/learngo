#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" gapi.go
echo "[Dev] finish build gapi.\n"