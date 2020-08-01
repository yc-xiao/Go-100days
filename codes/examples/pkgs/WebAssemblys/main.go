package main

import "syscall/js"

/*
	0. WebAssemblys　新的编码方式，可用于浏览器解析，更快速。 作为js性能的补充
	1. 先编写js逻辑，在将go文件编译成main.wasm
	2. 在html文件内引用go内置的js文件和.wasm并用js去解析.wasm文件.
 */

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")
}