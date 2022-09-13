package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// spew config
	//spew.Config.DisablePointerAddresses = true
	spew.Config.Indent = "  "

	compiler := newCompiler()
	compiler.compile("input.hs")

	//spew.Dump(typeChecker)
	//spew.Dump(ast)
}
