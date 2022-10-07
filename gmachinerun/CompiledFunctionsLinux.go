//go:build lin
// +build lin

package main

import "fmt"

func addPlatformDependentGlobals(m *GMachine) {
	m.addGlobal("getLine", 0, m.createIOFunInstructions(0, &IIOFun{fun: applyGetLine}))
}

func applyGetLine(m *GMachine) {
	var line string
	fmt.Scanln(&line)

	addr := m.allocNewNode(&NString{s: line})
	m.stack.put(addr)
}
