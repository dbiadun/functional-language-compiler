package main

/////////////////////////////////////// G MACHINE //////////////////////////////////////////

type GMachine struct {
	instrQueue *GInstrQueue
	stack      *GStack
	dump       *GDump
	heap       *GHeap
}

func newGMachine() *GMachine {
	gMachine := new(GMachine)
	gMachine.instrQueue = new(GInstrQueue)
	gMachine.stack = new(GStack)
	gMachine.dump = new(GDump)
	gMachine.heap = newGHeap()

	return gMachine
}

/////////////////////////////////// INSTRUCTIONS QUEUE /////////////////////////////////////

type GInstrQueue struct {
	q []GInstr
}

func (q *GInstrQueue) put(instr GInstr) {
	q.q = append(q.q, instr)
}

func (q *GInstrQueue) getFirst() GInstr {
	n := len(q.q)
	first := q.q[n-1]
	q.q = q.q[:n-1]
	return first
}

///////////////////////////////////////// STACK ////////////////////////////////////////////

type GStack struct {
	s []*GAddr
}

func (s *GStack) put(a *GAddr) {
	s.s = append(s.s, a)
}

func (s *GStack) readTop() *GAddr {
	n := len(s.s)
	if n == 0 {
		return nil
	} else {
		return s.s[n-1]
	}
}

////////////////////////////////////////// DUMP ////////////////////////////////////////////

type GDump struct {
	d []*GStack
}

func (d *GDump) put(s *GStack) {
	d.d = append(d.d, s)
}

func (d *GDump) get() *GStack {
	n := len(d.d)
	if n == 0 {
		return nil
	}

	s := d.d[n-1]
	d.d = d.d[:n-1]
	return s
}

////////////////////////////////////////// HEAP ////////////////////////////////////////////

type GHeap struct {
	h       map[int]GNode
	nextNum int
}

func newGHeap() *GHeap {
	h := new(GHeap)
	h.h = make(map[int]GNode)

	return h
}

func (h *GHeap) newAddr() *GAddr {
	n := h.nextNum
	h.nextNum++

	return &GAddr{a: n}
}

func (h *GHeap) get(addr *GAddr) GNode {
	node, ok := h.h[addr.a]
	if ok {
		return node
	} else {
		return nil
	}
}

func (h *GHeap) put(addr *GAddr, node GNode) {
	num := addr.a
	h.h[num] = node
}

//////////////////////////////////////// HELPERS ///////////////////////////////////////////

type GAddr struct {
	a int
}

//////////////////////////////////////// NODES /////////////////////////////////////////////

type GNode interface {
	gNode()
}

type BaseGNode struct{}

func (*BaseGNode) gNode() {}

////////////////////////////////////// INSTRUCTIONS ////////////////////////////////////////

type GInstr interface {
	gInstr()
}

type BaseGInstr struct{}

func (*BaseGInstr) gInstr() {}
