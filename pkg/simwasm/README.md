# simwasm

- [simwasm](#simwasm)
  - [Compile this program to WebAssembly by running the following command](#compile-this-program-to-webassembly-by-running-the-following-command)
  - [Javascript Glue](#javascript-glue)
  - [Executing WebAssembly with Node.js](#executing-webassembly-with-nodejs)
  - [References](#references)

## Compile this program to WebAssembly by running the following command

```shell
GOARCH=wasm GOOS=js go build -o ../../web/static/js/simwasm.wasm
```

## Javascript Glue

WebAssembly is supposed to exist hand in hand with JavaScript. Hence some JavaScript glue code is needed to import the WebAssembly Module we just created and run it in the browser. This code is already available in the Go installation. Let’s go ahead and copy it to our assets directory.

```shell
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ~/Documents/webassembly/assets/
```

- `GOOS=js GOARCH=wasm` - flag to generate binary file for js/wasm architecture.
- `-o <path>/name.wasm` - flag to specify the destination and name of the binary file. Needless to say, the build should happen in the same directory as the golang file.

## Executing WebAssembly with Node.js

It’s possible to execute compiled WebAssembly modules using Node.js rather than
a browser, which can be useful for testing and automation.

First, make sure Node is installed and in your PATH.

Then, add $(go env GOROOT)/misc/wasm to your PATH. This will allow go run and go
test find go_js_wasm_exec in a PATH search and use it to just work for js/wasm:

```shell
$ export PATH="$PATH:$(go env GOROOT)/misc/wasm"
$ GOOS=js GOARCH=wasm go run .
Hello, WebAssembly!
$ GOOS=js GOARCH=wasm go test
PASS
ok  	example.org/my/pkg	0.800s
```

Finally, the wrapper may also be used to directly execute a Go Wasm binary:

```shell

$ GOOS=js GOARCH=wasm go build -o mybin .
$ $(go env GOROOT)/misc/wasm/go_js_wasm_exec ./mybin
Hello, WebAssembly!
$ GOOS=js GOARCH=wasm go test -c
$ $(go env GOROOT)/misc/wasm/go_js_wasm_exec ./pkg.test
PASS
ok  	example.org/my/pkg	0.800s
```

## References

- <https://eremeev.ca/posts/golang-wasm-working-with-form/>
- <https://golangbot.com/webassembly-using-go/>
- <https://github.com/golang/go/wiki/WebAssembly>
- <https://github.com/markfarnan/go-canvas>
