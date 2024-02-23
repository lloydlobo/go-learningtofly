// Reference:
// 	- https://golangbot.com/webassembly-using-go/
// 	- https://eremeev.ca/posts/golang-wasm-working-with-form/#create-a-golang-file

package simwasm

import (
	"encoding/json"
	"fmt"
	"syscall"
	"syscall/js"
)

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

// jsonWrapper exposes a function from Go to Javascript.
func jsonWrapper() js.Func {
	fmt.Printf("syscall.AF_INET: %v\n", syscall.AF_INET)
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no. off arguments passed"
		}

		inputJSON := args[0].String()
		fmt.Printf("input %s\n", &inputJSON)

		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}

		return pretty
	})

	return jsonFunc
}

// import (
// 	// "syscall/js" - is the package that allows us to communicate with browsers
// 	// using JavaScript APIs. There are lots of packages on the internet, but
// 	// the core of all of them is "syscall/js".
// 	"syscall/js"
// 	// ^
// 	// | $ go vet
// 	// | package simwasm
// 	// |         imports syscall/js: build constraints exclude all Go files in C:\Program Files\Go\src\syscall\js
// )
//
// func main() {
// 	document := js.Global().Get("document")
// 	document.Call("getElementById", "calculate")
// 	<-make(chan bool)
// }
