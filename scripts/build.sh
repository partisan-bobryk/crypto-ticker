#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/crypto-ticker-x64-linux ./cmd/*
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./bin/crypto-ticker-x64-darwin ./cmd/*

