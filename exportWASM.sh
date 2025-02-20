#!/bin/bash

GOARCH=wasm GOOS=js go build -o game.wasm main.go
sudo cp /usr/local/go/misc/wasm/wasm_exec.js .

mv game.wasm docs/
mv wasm_exec.js docs/