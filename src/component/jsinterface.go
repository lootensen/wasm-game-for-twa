//go:build js && wasm

package component

import (
	"fmt"
	"strings"
	"syscall/js"
)

type JSInterface struct{}

func NewJSInterface() *JSInterface {
	return &JSInterface{}
}

func (jsi *JSInterface) CallFunction(funcName string, args ...interface{}) {

	splittedFuncName := splitFuncName(funcName)
	// root jsFunc
	jsFunc := js.Global()
	// now iterates through the chained functions and returns them
	// this way we end up with an invokable function
	for _, part := range splittedFuncName {
		jsFunc = jsFunc.Get(part)
		if jsFunc.IsUndefined() {
			fmt.Printf("JS function or property '%s' not defined\n", part)
			return
		}
	}

	// we check if jsFunc is actually a function
	if jsFunc.Type() == js.TypeFunction {

		// we initialize a new interface array with the length of args
		// we loop over the array and use the js.ValueOf function to map the JS types to the arguments
		jsArgs := make([]interface{}, len(args))
		for i, arg := range args {
			jsArgs[i] = js.ValueOf(arg)
		}
		jsFunc.Invoke(jsArgs...)
	} else {
		fmt.Println("JS function not defined", funcName)
	}
}

func splitFuncName(funcName string) []string {
	return strings.Split(funcName, ".")
}
