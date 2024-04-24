//go:build !js && !wasm

package main

import (
	"github.com/lootensen/wasm-demo-game/app"
)

func main() {
	app.RunApp(nil)
}
