#!/bin/bash

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/crypto-ticker-linux-64 ./cmd/*
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/crypto-ticker-darwin-64 ./cmd/*

