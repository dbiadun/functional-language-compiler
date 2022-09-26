package main

import "machine"
import "time"
import "device/nrf"
import "runtime/interrupt"

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

	callback := func(p machine.Pin) {
		interruptionData := InterruptionData{aFun, int(p)}
		select {
		case interruptions <- interruptionData:
		default:
		}
	}

	pin.SetInterrupt(pinChange, callback)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func applyTinySetTimer(m *GMachine) {
	aTimerNum := m.stack.get()
	aDuration := m.stack.get()
	aAction := m.stack.get()

	if aTimerNum == nil || aDuration == nil || aAction == nil {
		errFatal("Empty stack while trying to apply tinySetTimer")
	}

	timerNum, ok1 := m.heap.get(aTimerNum).(*NInt)
	duration, ok2 := m.heap.get(aDuration).(*NInt)
	action := m.heap.get(aAction)

	if !ok1 || !ok2 || action == nil {
		errFatal("tinySetTimer arg address not pointing to a proper node type")
	}

	timerCallbackActions[timerNum.n] = aAction
	m.interruptionCallbacks = append(m.interruptionCallbacks, aAction)

	timer := getTimer(timerNum.n)

	timer.SetTASKS_STOP(nrf.TIMER_TASKS_STOP_TASKS_STOP_Trigger)

	timer.SetPRESCALER(4)                                                   // 16Mhz / 2^4 = 1Mhz = 1_000_000 ticks per second
	timer.SetBITMODE(nrf.TIMER_BITMODE_BITMODE_32Bit)                       // 32 bits for timer width
	timer.SetCC(0, uint32(duration.n))                                      // One capture every `duration` microseconds
	timer.SetSHORTS_COMPARE0_CLEAR(nrf.TIMER_SHORTS_COMPARE0_CLEAR_Enabled) // Set clearing the timer after hitting wanted value
	timer.SetINTENSET_COMPARE0(nrf.TIMER_INTENSET_COMPARE0_Set)             // Enable interrupt

	intr := newTimerInterrupt(timerNum.n)
	intr.SetPriority(0xc0)
	intr.Enable()

	timer.SetTASKS_CLEAR(nrf.TIMER_TASKS_CLEAR_TASKS_CLEAR_Trigger)
	timer.SetTASKS_START(nrf.TIMER_TASKS_START_TASKS_START_Trigger)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func newTimerInterrupt(timerNum int) interrupt.Interrupt {
	var intr interrupt.Interrupt

	switch timerNum {
	case 1:
		intr = interrupt.New(nrf.IRQ_TIMER1, func(intr interrupt.Interrupt) {
			nrf.TIMER1.SetEVENTS_COMPARE(0, nrf.TIMER_EVENTS_COMPARE_EVENTS_COMPARE_NotGenerated) // Remove bit triggering the interrupt
			interruptions <- InterruptionData{timerCallbackActions[1], -1}
		})
	case 2:
		intr = interrupt.New(nrf.IRQ_TIMER2, func(intr interrupt.Interrupt) {
			nrf.TIMER2.SetEVENTS_COMPARE(0, nrf.TIMER_EVENTS_COMPARE_EVENTS_COMPARE_NotGenerated) // Remove bit triggering the interrupt
			interruptions <- InterruptionData{timerCallbackActions[2], -1}
		})
	case 3:
		intr = interrupt.New(nrf.IRQ_TIMER3, func(intr interrupt.Interrupt) {
			nrf.TIMER3.SetEVENTS_COMPARE(0, nrf.TIMER_EVENTS_COMPARE_EVENTS_COMPARE_NotGenerated) // Remove bit triggering the interrupt
			interruptions <- InterruptionData{timerCallbackActions[3], -1}
		})
	case 4:
		intr = interrupt.New(nrf.IRQ_TIMER4, func(intr interrupt.Interrupt) {
			nrf.TIMER4.SetEVENTS_COMPARE(0, nrf.TIMER_EVENTS_COMPARE_EVENTS_COMPARE_NotGenerated) // Remove bit triggering the interrupt
			interruptions <- InterruptionData{timerCallbackActions[4], -1}
		})
	default:
		intr = interrupt.New(nrf.IRQ_TIMER0, func(intr interrupt.Interrupt) {
			nrf.TIMER0.SetEVENTS_COMPARE(0, nrf.TIMER_EVENTS_COMPARE_EVENTS_COMPARE_NotGenerated) // Remove bit triggering the interrupt
			interruptions <- InterruptionData{timerCallbackActions[0], -1}
		})
	}

	return intr
}

func applyTinyStopTimer(m *GMachine) {
	aTimerNum := m.stack.get()

	if aTimerNum == nil {
		errFatal("Empty stack while trying to apply tinyStopTimer")
	}

	timerNum, ok1 := m.heap.get(aTimerNum).(*NInt)

	if !ok1 {
		errFatal("tinyStopTimer arg address not pointing to a proper node type")
	}

	timer := getTimer(timerNum.n)

	timer.SetTASKS_STOP(nrf.TIMER_TASKS_STOP_TASKS_STOP_Trigger)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func applyTinyStartTimer(m *GMachine) {
	aTimerNum := m.stack.get()

	if aTimerNum == nil {
		errFatal("Empty stack while trying to apply tinyStartTimer")
	}

	timerNum, ok1 := m.heap.get(aTimerNum).(*NInt)

	if !ok1 {
		errFatal("tinyStartTimer arg address not pointing to a proper node type")
	}

	timer := getTimer(timerNum.n)

	timer.SetTASKS_START(nrf.TIMER_TASKS_START_TASKS_START_Trigger)

	addr := m.allocNewNode(&NUnit{})
	m.stack.put(addr)
}

func getTimer(timerNum int) *nrf.TIMER_Type {
	var timer *nrf.TIMER_Type

	switch timerNum {
	case 1:
		timer = nrf.TIMER1
	case 2:
		timer = nrf.TIMER2
	case 3:
		timer = nrf.TIMER3
	case 4:
		timer = nrf.TIMER4
	default:
		timer = nrf.TIMER0
	}

	return timer
}
