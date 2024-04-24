//go:build js && wasm

package main

import (
	"github.com/lootensen/wasm-demo-game/app"
	"github.com/lootensen/wasm-demo-game/component"
)

func main() {
	app.RunApp(&component.JSInterface{})
}
