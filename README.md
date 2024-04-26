# WASM and Go

## Compiling by yourself

### Compiling your game

```shell
env GOOS=js GOARCH=wasm go build -o mygame.wasm github.com/lootensen/wasm-demo-game
```

### Copying wasm_exec.js to execute the WebAssembly binary

```shell
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
```


## GH Pages

[Zu Github Pages (TWA)](https://lootensen.github.io/wasm-game-for-twa/)