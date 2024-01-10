#! bin/bash
GOOS=js GOARCH=wasm go build -C game -o ../web/game.wasm