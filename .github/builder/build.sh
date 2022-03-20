#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o .bin/axolotl-amd64.exe .
GOOS=windows GOARCH=386 go build -o .bin/axolotl-386.exe .
# 64-bit
GOOS=darwin GOARCH=amd64 go build -o .bin/axolotl-amd64-darwin .
# 32-bit
GOOS=darwin GOARCH=386 go build -o .bin/axolotl-386-darwin .
# 64-bit
GOOS=linux GOARCH=amd64 go build -o .bin/axolotl-amd64-linux .
# 32-bit
GOOS=linux GOARCH=386 go build -o .bin/axolotl-386-linux .