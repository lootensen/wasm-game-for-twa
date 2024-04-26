//go:build !js && !wasm

package main

import (
	"github.com/lootensen/wasm-demo-game/src/app"
	"github.com/lootensen/wasm-demo-game/src/component"
)

func main() {
	app.RunApp(&component.JSInterface{})
}
