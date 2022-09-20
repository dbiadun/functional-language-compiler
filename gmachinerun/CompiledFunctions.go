package main

import "machine"
import "time"

func applyReturn() {
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

//func applyGetLine(m *GMachine) {
//	var line string
//	fmt.Scanln(&line)
//
//	addr := m.allocNewNode(&NString{s: line})
//	m.stack.put(addr)
//}

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

func applyTinyLED(m *GMachine) {
	led := machine.LED
	addr := m.allocNewNode(&NInt{n: int(led)})
	m.stack.put(addr)
}

func applyTinyConfigure(m *GMachine) {
	a1 := m.stack.get()
	a2 := m.stack.get()

	if a1 == nil || a2 == nil {
		errFatal("Empty stack while trying to apply tinyConfigure")
	}

	pinNum, ok1 := m.heap.get(a1).(*NInt)
	modeNum, ok2 := m.heap.get(a2).(*NInt)

	if !ok1 || !ok2 {
		errFatal("tinyConfigure arg address not pointing to a proper node type")
	}

	pin := machine.Pin(pinNum.n)
	mode := machine.PinMode(modeNum.n)

	pin.Configure(machine.PinConfig{Mode: mode})

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func applyTinyPinOutput(m *GMachine) {
	output := machine.PinOutput
	addr := m.allocNewNode(&NInt{n: int(output)})
	m.stack.put(addr)
}

func applyTinyLow(m *GMachine) {
	a := m.stack.get()

	if a == nil {
		errFatal("Empty stack while trying to apply tinyLow")
	}

	pinNum, ok := m.heap.get(a).(*NInt)

	if !ok {
		errFatal("tinyLow arg address not pointing to a proper node type")
	}

	pin := machine.Pin(pinNum.n)

	pin.Low()

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func applyTinyHigh(m *GMachine) {
	a := m.stack.get()

	if a == nil {
		errFatal("Empty stack while trying to apply tinyHigh")
	}

	pinNum, ok := m.heap.get(a).(*NInt)

	if !ok {
		errFatal("tinyHigh arg address not pointing to a proper node type")
	}

	pin := machine.Pin(pinNum.n)

	pin.High()

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}
