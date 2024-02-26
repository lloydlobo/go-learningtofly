// # Errorlog
//
//   - 20240224: syscall/js isn't recognized by linter but gets build via GOOS=js
//     and GOARCH=wasm
//   - 20240225: wasm magic word error when cmd Go files use "package cmd" instead
//     of "package main"
package main

import (
	"fmt"
	// "wasmgo/internal/dom"
)

func main() {
	fmt.Println("simwasm/main says: Hello, world!")

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
	// |	ch := make(chan struct{})
	// |	<-ch // OR use <-make(chan bool)
	// |	// ^
	// |	// | Uncaught Error: Go program has already exited
	// |	// | 	at globalThis.Go._resume (wasm_exec.js:543:11)
	// |	// | 	at HTMLInputElement.<anonymous> (wasm_exec.js:556:8)
	// |	// |
	// |	// | Program executes and exits, so when we try to invoke my event handler,
	// |	// | there's no process to do that. So, we need some way to keep the program
	// |	// | running. we'll do this by creating a channel, waiting for an event on it, but
	// |	// | never sending one.
	// |	// ---
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
