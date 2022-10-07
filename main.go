package main

import "functional-language-compiler/compiler"

func main() {
	cmp := compiler.NewCompiler()
	cmp.Compile()
}
