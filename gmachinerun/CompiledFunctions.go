package main

import "time"

var timerCallbackActions [5]*GAddr

func applyReturn(_ *GMachine) {
	return
}

func applyPutStr(m *GMachine) {
	a := m.stack.get()
	if a == nil {
		errFatal("Empty stack while trying to apply putStr")
	}

	node, ok := m.heap.get(a).(*NString)
	if !ok {
		errFatal("putStr arg address not pointing to an NString")
	}

	print(node.s)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func applyPutInt(m *GMachine) {
	a := m.stack.get()
	if a == nil {
		errFatal("Empty stack while trying to apply putInt")
	}

	node, ok := m.heap.get(a).(*NInt)
	if !ok {
		errFatal("putStr arg address not pointing to an NInt")
	}

	print(node.n)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func applyStateSetInt(m *GMachine) {
	aID := m.stack.get()
	aVal := m.stack.get()
	if aID == nil || aVal == nil {
		errFatal("Empty stack while trying to apply stateSetInt")
	}

	id, ok := m.heap.get(aID).(*NString)
	val, ok := m.heap.get(aVal).(*NInt)
	if !ok {
		errFatal("stateSetInt arg address not pointing to a proper node")
	}

	m.state.setInt(id.s, val.n)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func applyStateGetInt(m *GMachine) {
	aID := m.stack.get()
	if aID == nil {
		errFatal("Empty stack while trying to apply stateGetInt")
	}

	id, ok := m.heap.get(aID).(*NString)
	if !ok {
		errFatal("stateGetInt arg address not pointing to a proper node")
	}

	val := m.state.getInt(id.s)

	addr := m.allocNewNode(&NInt{n: val})
	m.stack.put(addr)
}

func applyTinySleep(m *GMachine) {
	a := m.stack.get()

	if a == nil {
		errFatal("Empty stack while trying to apply tinyHigh")
	}

	milliseconds, ok := m.heap.get(a).(*NInt)

	if !ok {
		errFatal("tinyHigh arg address not pointing to a proper node type")
	}

	time.Sleep(time.Duration(int64(time.Millisecond) * int64(milliseconds.n)))

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}
