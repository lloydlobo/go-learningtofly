# simwasm

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

## References

- <https://eremeev.ca/posts/golang-wasm-working-with-form/>
- <https://golangbot.com/webassembly-using-go/>
