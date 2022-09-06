package main

import (
	"machine"
	"time"
)

func main() {
	gMachine := NewGMachine()
	_ = gMachine

	instrs := []GInstr{
		&IPushInt{n: 3},
		&IPushInt{n: 20},
		&IPushInt{n: 27},
		&IBinOp{op: '-'},
		&IPushInt{n: 5},
		&IBinOp{op: '+'},
		&IBinOp{op: '/'},
	} // ((27 - 20) + 5) / 3 = 4

	gMachine.instrQueue.putN(instrs)
	res := gMachine.run()

	const TIME = 500

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for i := 0; i < res; i++ {
		led.Low()
		time.Sleep(time.Millisecond * TIME)

		led.High()
		time.Sleep(time.Millisecond * TIME)
	}
}
