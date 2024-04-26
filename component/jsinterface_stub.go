//go:build !js && !wasm

package component

type JSInterface struct{}

func (jsi *JSInterface) CallFunction(funcName string, args ...interface{}) {
	// Eine Stub-Implementierung, die nichts tut.
	// fmt.Println("Skip JSInterface CallFunction")
}
