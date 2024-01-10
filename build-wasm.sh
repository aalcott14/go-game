#! bin/bash
cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/
GOOS=js GOARCH=wasm go build -C game -o ../web/game.wasm