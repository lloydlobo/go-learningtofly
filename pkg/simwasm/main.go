// # Errorlog
//
//   - 20240224: syscall/js isn't recognized by linter but gets build via GOOS=js
//     and GOARCH=wasm
//     If you are using an IDE like VS Code, there might be a configuration issue
//     that is preventing the IDE from recognizing the syscall/js package. You can
//     try setting the build constraints explicitly in your IDE settings to GOOS=js
//     and GOARCH=wasm. This will tell the IDE to build your program for WASM and
//     allow it to recognize the syscall/js package.
//
//   - 20240225: wasm magic word error when cmd Go files use "package cmd" instead
//     of "package main"
package main

import (
	"fmt"

	// Package js gives access to the WebAssembly host environment when using the js/wasm architecture. Its API is based on JavaScript semantics.
	// This package is EXPERIMENTAL. Its current scope is only to allow tests to run, but not yet to provide a comprehensive API for users. It is exempt from the Go compatibility promise.
	"syscall/js"
)

func Add(a, b int) int {
	fmt.Println("[GO] Log file.go", a, b)
	return a + b
}

func FuncName(M int, durations []int) int {
	fmt.Println("[GO] Log file.go", M, durations)
	return 10
}

// TinyGo supports a //export <name> or alias //go:export <name> comment
// directive that does what you're looking for.
//
// The standard Go compiler has an [ongoing open
// discussion](https://github.com/golang/go/issues/25612) about replicating TinyGo
// feature. The tl;dr seems to be that you can achieve this by setting funcs to the
// JS global namespace, with the js.Global().Set(...)
//
// Source: https://stackoverflow.com/a/67983256

func main() {
	fmt.Println("simwasm/main says: Hello, world!")

	// Mount Go function on the JavaScript global object (usually `window`, or `global`)
	// Note: Return interface{} or any
	// Reference: https://stackoverflow.com/a/76082718
	{
		js.Global().Set("Add", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) != 2 {
				fmt.Println("invalid number of args")
				return nil
			}
			if args[0].Type() != js.TypeObject {
				fmt.Println("the first argument should be a number")
			}
			if args[1].Type() != js.TypeObject {
				fmt.Println("the second argument should be a number")
			}
			return Add(args[0].Int(), args[1].Int())
		}))

		js.Global().Set("FuncName", js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) != 2 {
				fmt.Println("invalid number of args")
				return nil
			}
			if args[0].Type() != js.TypeNumber {
				fmt.Println("the first argument should be a number")
				return nil
			}

			arg := args[1]
			if arg.Type() != js.TypeObject {
				fmt.Println("the second argument should be an array")
				return nil
			}

			durations := make([]int, arg.Length())
			for i := 0; i < len(durations); i++ {
				item := arg.Index(i)
				if item.Type() != js.TypeNumber {
					fmt.Printf("the item at index %d should be a number\n", i)
					return nil
				}
				durations[i] = item.Int()
			}

			// Call the actual func.
			return FuncName(args[0].Int(), durations)
		}))
	}

	document := js.Global().Get("document")
	document.Call("getElementById", "loading").Get("style").Call("setProperty", "display", "none")

	// ---
	// | document := js.Global().Get("document")
	// | document.Call("getElementById", "loading").Get("style").Call("setProperty", "display", "none")
	// | document.Call("getElementById", "calc").Get("style").Call("setProperty", "display", "flex") // set it to block
	// v

	// |	dom.Hide("loading")
	// |	dom.Show("calc", "inline-flex")
	// |
	// |	if __DEBUG_hardcoded := false; __DEBUG_hardcoded {
	// |		if __DEBUG_invalidValue := true; __DEBUG_invalidValue {
	// |			dom.SetValue("first-number", "value", "X")
	// |		} else {
	// |			dom.SetValue("first-number", "value", "7")
	// |		}
	// |		dom.SetValue("second-number", "value", "5")
	// |		dom.SetValue("result", "value", "0")
	// |
	// |		calculate()
	// |	} else {
	// |		dom.SetValue("first-number", "value", "0")
	// |		dom.SetValue("second-number", "value", "0")
	// |		dom.SetValue("result", "value", "0")
	// |
	// |		if __DEBUG_tinygo := true; __DEBUG_tinygo {
	// |			dom.AddEventListenerTinygo("first-number", "input", calculateTinygo)
	// |			dom.AddEventListenerTinygo("second-number", "input", calculateTinygo)
	// |		} else {
	// |			dom.AddEventListener("first-number", "input", calculate)
	// |			dom.AddEventListener("second-number", "input", calculate)
	// |		}
	// |	}
	// |
	// ch := make(chan struct{})
	// <-ch // OR use <-make(chan bool)
	// ^
	// | Uncaught Error: Go program has already exited
	// | 	at globalThis.Go._resume (wasm_exec.js:543:11)
	// | 	at HTMLInputElement.<anonymous> (wasm_exec.js:556:8)
	// |
	// | Program executes and exits, so when we try to invoke my event handler,
	// | there's no process to do that. So, we need some way to keep the program
	// | running. we'll do this by creating a channel, waiting for an event on it, but
	// | never sending one.
	// ---
	// | OR
	// | Prevent the program from exiting.
	// | Note: the exported func should be released if you don't need it any more,
	// | and let the program exit after then. To simplify this demo, this is
	// | omitted. See https://pkg.go.dev/syscall/js#Func.Release for more information.
	// v
	select {}
}

// | func calculate() {
// | 	firstNum, firstNumErr := strconv.Atoi(dom.GetString("first-number", "value"))
// | 	secondNum, secondNumErr := strconv.Atoi(dom.GetString("second-number", "value"))
// |
// | 	if firstNumErr == nil && secondNumErr == nil { // is nil
// | 		dom.SetValue("result", "value", strconv.Itoa(firstNum+secondNum))
// | 		dom.AddClass("result", "error")
// | 	} else { // is not nil
// | 		dom.SetValue("result", "value", "ERR")
// | 		dom.AddClass("result", "error")
// | 	}
// | }
// |
// | func calculateTinygo(_ js.Value, _ []js.Value) interface{} {
// | 	firstNum, firstNumErr := strconv.Atoi(dom.GetString("first-number", "value"))
// | 	secondNum, secondNumErr := strconv.Atoi(dom.GetString("second-number", "value"))
// |
// | 	if firstNumErr == nil && secondNumErr == nil { // is nil
// | 		dom.SetValue("result", "value", strconv.Itoa(firstNum+secondNum))
// | 		dom.AddClass("result", "error")
// | 	} else { // is not nil
// | 		dom.SetValue("result", "value", "ERR")
// | 		dom.AddClass("result", "error")
// | 	}
// |
// | 	return nil
// | }

// package simwasm

// import (
// 	"encoding/json"
// 	"fmt"
// 	"syscall"
// 	"syscall/js"
// )

// func main() {
// 	fmt.Println("Go Web Assembly")
// 	js.Global().Set("formatJSON", jsonWrapper())
// }

// func prettyJson(input string) (string, error) {
// 	var raw any
// 	if err := json.Unmarshal([]byte(input), &raw); err != nil {
// 		return "", err
// 	}
// 	pretty, err := json.MarshalIndent(raw, "", "  ")
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(pretty), nil
// }

// // jsonWrapper exposes a function from Go to Javascript.
// func jsonWrapper() js.Func {
// 	fmt.Printf("syscall.AF_INET: %v\n", syscall.AF_INET)
// 	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
// 		if len(args) != 1 {
// 			return "Invalid no. off arguments passed"
// 		}

// 		inputJSON := args[0].String()
// 		fmt.Printf("input %s\n", &inputJSON)

// 		pretty, err := prettyJson(inputJSON)
// 		if err != nil {
// 			fmt.Printf("unable to convert to json %s\n", err)
// 			return err.Error()
// 		}

// 		return pretty
// 	})

// 	return jsonFunc
// }
