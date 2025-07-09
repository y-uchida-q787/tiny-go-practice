//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("WASM Go Initialized")

	c := make(chan struct{}, 0)

	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		name := args[0].String()
		fmt.Println("name:", name)
		return fmt.Sprintf("Hello, Nanashino %s!", name)
	})

	js.Global().Set("getGreeting", cb)
	<-c
}
