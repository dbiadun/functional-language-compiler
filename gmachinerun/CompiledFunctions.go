package main

import "machine"
import "time"

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

func applyTinyLED2(m *GMachine) {
	led := machine.LED2
	addr := m.allocNewNode(&NInt{n: int(led)})
	m.stack.put(addr)
}

func applyTinyBUTTON(m *GMachine) {
	button := machine.BUTTON
	addr := m.allocNewNode(&NInt{n: int(button)})
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

func applyTinyPinInput(m *GMachine) {
	input := machine.PinInput
	addr := m.allocNewNode(&NInt{n: int(input)})
	m.stack.put(addr)
}

func applyTinyPinInputPullup(m *GMachine) {
	inputPullup := machine.PinInputPullup
	addr := m.allocNewNode(&NInt{n: int(inputPullup)})
	m.stack.put(addr)
}

func applyTinyPinInputPulldown(m *GMachine) {
	inputPulldown := machine.PinInputPulldown
	addr := m.allocNewNode(&NInt{n: int(inputPulldown)})
	m.stack.put(addr)
}

func applyTinyPinOutput(m *GMachine) {
	output := machine.PinOutput
	addr := m.allocNewNode(&NInt{n: int(output)})
	m.stack.put(addr)
}

func applyTinyPinRising(m *GMachine) {
	rising := machine.PinRising
	addr := m.allocNewNode(&NInt{n: int(rising)})
	m.stack.put(addr)
}

func applyTinyPinFalling(m *GMachine) {
	falling := machine.PinFalling
	addr := m.allocNewNode(&NInt{n: int(falling)})
	m.stack.put(addr)
}

func applyTinyPinToggle(m *GMachine) {
	toggle := machine.PinToggle
	addr := m.allocNewNode(&NInt{n: int(toggle)})
	m.stack.put(addr)
}

func applyTinyGet(m *GMachine) {
	a := m.stack.get()

	if a == nil {
		errFatal("Empty stack while trying to apply tinyGet")
	}

	pinNum, ok := m.heap.get(a).(*NInt)

	if !ok {
		errFatal("tinyGet arg address not pointing to a proper node type")
	}

	state := machine.Pin(pinNum.n).Get()
	addr := m.allocNewNode(boolToNData(state))
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

func applyTinySetInterrupt(m *GMachine) {
	aPin := m.stack.get()
	aPinChange := m.stack.get()
	aFun := m.stack.get()
	m.interruptionCallbacks = append(m.interruptionCallbacks, aFun)

	if aPin == nil || aPinChange == nil || aFun == nil {
		errFatal("Empty stack while trying to apply tinySetInterrupt")
	}

	pinNum, ok1 := m.heap.get(aPin).(*NInt)
	pinChangeNum, ok2 := m.heap.get(aPinChange).(*NInt)
	fun := m.heap.get(aFun)

	if !ok1 || !ok2 || fun == nil {
		errFatal("tinySetInterrupt arg address not pointing to a proper node type")
	}

	pin := machine.Pin(pinNum.n)
	pinChange := machine.PinChange(pinChangeNum.n)
	//funIntAddr := aFun.a

	callback := func(p machine.Pin) {
		interruptionData := InterruptionData{aFun, int(p)}
		select {
		case m.interruptions <- interruptionData:
		default:
		}
	}

	pin.SetInterrupt(pinChange, callback)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}
