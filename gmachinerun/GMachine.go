package main

/////////////////////////////////////// G MACHINE //////////////////////////////////////////

type GMachine struct {
	instrQueue  *GInstrQueue
	stack       *GStack
	dump        *GDump
	heap        *GHeap
	globalMap   *GGlobalMap
	gcRound     int
	gcNodeCount int
	gcThreshold int
}

func NewGMachine() *GMachine {
	gMachine := new(GMachine)
	gMachine.instrQueue = new(GInstrQueue)
	gMachine.stack = new(GStack)
	gMachine.dump = new(GDump)
	gMachine.heap = newGHeap()
	gMachine.globalMap = newGGlobalMap()

	gMachine.addCompiledGlobals()

	return gMachine
}

func (m *GMachine) run() int {
	for {
		for m.instrQueue.size() > 0 {
			instr := m.instrQueue.get()
			instr.apply(m)
		}

		addr := m.stack.get()
		if addr == nil {
			return 1
		}

		res1 := m.heap.get(addr)
		if res1 == nil {
			return 1
		}

		_, ok := res1.(*NUnit)
		if ok {
			return 0
		} else {
			m.stack.put(addr)
			m.instrQueue.put(&IUnwind{})
		}
	}
}

func (m *GMachine) addGlobal(name string, arity int, instrs []GInstr) {
	code := &GCode{c: instrs}
	node := &NGlobal{argsNum: arity, code: code}
	addr := m.allocNewNode(node)
	m.globalMap.put(name, addr)
}

func (m *GMachine) addCompiledGlobals() {
	m.addGlobal("+", 2, m.createFunInstructions(2, &IBinOp{op: "+"}))
	m.addGlobal("-", 2, m.createFunInstructions(2, &IBinOp{op: "-"}))
	m.addGlobal("*", 2, m.createFunInstructions(2, &IBinOp{op: "*"}))
	m.addGlobal("/", 2, m.createFunInstructions(2, &IBinOp{op: "/"}))
	m.addGlobal("<", 2, m.createFunInstructions(2, &ICompOp{op: "<"}))
	m.addGlobal(">", 2, m.createFunInstructions(2, &ICompOp{op: ">"}))
	m.addGlobal("<=", 2, m.createFunInstructions(2, &ICompOp{op: "<="}))
	m.addGlobal(">=", 2, m.createFunInstructions(2, &ICompOp{op: ">="}))
	m.addGlobal("==", 2, m.createFunInstructions(2, &ICompOp{op: "=="}))
	m.addGlobal("/=", 2, m.createFunInstructions(2, &ICompOp{op: "/="}))
	m.addGlobal("&&", 2, m.createFunInstructions(2, &ILogOp{op: "&&"}))
	m.addGlobal("||", 2, m.createFunInstructions(2, &ILogOp{op: "||"}))

	m.addGlobal("return", 1, m.createIOFunInstructions(1, &IIOFun{fun: "return"}))
	m.addGlobal("putStr", 1, m.createIOFunInstructions(1, &IIOFun{fun: "putStr"}))
	m.addGlobal("putInt", 1, m.createIOFunInstructions(1, &IIOFun{fun: "putInt"}))
	m.addGlobal("getLine", 0, m.createIOFunInstructions(0, &IIOFun{fun: "getLine"}))

	m.addGlobal("tinySleep", 1, m.createIOFunInstructions(1, &IIOFun{fun: "tinySleep"}))

	m.addGlobal("tinyLED", 0, m.createIOFunInstructions(0, &IIOFun{fun: "tinyLED"}))

	m.addGlobal("tinyConfigure", 2, m.createIOFunInstructions(2, &IIOFun{fun: "tinyConfigure"}))
	m.addGlobal("tinyPinOutput", 0, m.createIOFunInstructions(0, &IIOFun{fun: "tinyPinOutput"}))

	m.addGlobal("tinyLow", 1, m.createIOFunInstructions(1, &IIOFun{fun: "tinyLow"}))
	m.addGlobal("tinyHigh", 1, m.createIOFunInstructions(1, &IIOFun{fun: "tinyHigh"}))
}

func (m *GMachine) createFunInstructions(arity int, funInstr GInstr) []GInstr {
	var instrs []GInstr
	for i := 0; i < arity; i++ {
		instrs = append(instrs, &IPush{n: arity - 1})
		instrs = append(instrs, &IEval{})
	}
	instrs = append(instrs, funInstr)
	instrs = append(instrs, &IUpdate{n: arity})
	instrs = append(instrs, &IPop{n: arity})
	instrs = append(instrs, &IUnwind{})

	return instrs
}

func (m *GMachine) createIOFunInstructions(arity int, funInstr GInstr) []GInstr {
	var instrs []GInstr
	for i := 0; i < arity; i++ {
		instrs = append(instrs, &IPush{n: arity - 1})
		instrs = append(instrs, &IEval{})
	}
	instrs = append(instrs, funInstr)

	// We can't update the node corresponding to an IO function, because it may need to evaluate more than once
	instrs = append(instrs, &ISlide{n: arity + 1})

	// We need it to prevent from exiting the program
	instrs = append(instrs, &IUnwind{})

	return instrs
}

func (m *GMachine) allocNewNode(node GNode) *GAddr {
	addr := m.heap.putNew(node)
	m.gcNodeCount++
	m.gcRunIfNeeded(addr)

	return addr
}

//func (m *GMachine) printStack() {
//	stack := m.stack.s
//	if stack == nil {
//		fmt.Println("ajdlfajf")
//		return
//	}
//
//	for _, addr := range stack {
//		node := m.heap.get(addr)
//		if node == nil {
//			fmt.Print("Nil, ")
//			continue
//		}
//		fmt.Print(fmt.Sprintf("%s, ", node))
//	}
//	fmt.Println()
//}

func errFatal(s string) {
	panic("Runtime error: " + s)
}

/////////////////////////////////// GARBAGE COLLECTION /////////////////////////////////////

func (m *GMachine) gcRunIfNeeded(newNodeToVisit *GAddr) {
	if m.gcNodeCount > m.gcThreshold {
		m.gcRun(newNodeToVisit)
		m.gcThreshold = 2 * m.gcNodeCount
	}
}

func (m *GMachine) gcRun(newNodeToVisit *GAddr) {
	m.gcRound++
	m.gcMark(newNodeToVisit)
	m.gcSweep()
}

func (m *GMachine) gcMark(newNodeToVisit *GAddr) {
	m.gcVisitStack(m.stack)
	for _, dumpEntry := range m.dump.d {
		m.gcVisitStack(dumpEntry.s)
	}
	for _, addr := range m.globalMap.m {
		m.gcVisitNode(addr)
	}
	m.gcVisitNode(newNodeToVisit)
}

func (m *GMachine) gcVisitStack(stack *GStack) {
	if stack.s == nil {
		return
	}

	for _, addr := range stack.s {
		m.gcVisitNode(addr)
	}
}

func (m *GMachine) gcVisitNode(addr *GAddr) {
	node := m.heap.get(addr)
	if node != nil {
		node.setGcRound(m.gcRound)

		switch node := node.(type) {
		case *NInt:
		case *NApp:
			m.gcVisitNode(node.fun)
			m.gcVisitNode(node.arg)
		case *NGlobal:
		case *NInd:
			m.gcVisitNode(node.a)
		case *NData:
			for _, arg := range node.args {
				m.gcVisitNode(arg)
			}
		}
	}
}

func (m *GMachine) gcSweep() {
	for addr, node := range m.heap.h {
		if node == nil || node.getGcRound() != m.gcRound {
			m.gcNodeCount--
			m.heap.h[addr] = nil
			delete(m.heap.h, addr)
		}
	}
}

/////////////////////////////////// INSTRUCTIONS QUEUE /////////////////////////////////////

type GInstrQueue struct {
	q []GInstr
}

func (q *GInstrQueue) put(instr GInstr) {
	q.q = append(q.q, instr)
}

func (q *GInstrQueue) putN(instrs []GInstr) {
	n := len(instrs)
	for i := n - 1; i >= 0; i-- {
		q.put(instrs[i])
	}
}

func (q *GInstrQueue) get() GInstr {
	n := len(q.q)
	first := q.q[n-1]
	q.q = q.q[:n-1]
	return first
}

func (q *GInstrQueue) size() int {
	return len(q.q)
}

///////////////////////////////////////// STACK ////////////////////////////////////////////

type GStack struct {
	s []*GAddr
}

func (s *GStack) put(a *GAddr) {
	s.s = append(s.s, a)
}

func (s *GStack) putN(addrs []*GAddr) {
	n := len(addrs)
	for i := n - 1; i >= 0; i-- {
		s.put(addrs[i])
	}
}

func (s *GStack) readTop() *GAddr {
	n := len(s.s)
	if n == 0 {
		return nil
	} else {
		return s.s[n-1]
	}
}

func (s *GStack) readNthFromTop(n int) *GAddr {
	l := len(s.s)
	if l < n+1 {
		return nil
	}

	addr := s.s[l-1-n]
	return addr
}

func (s *GStack) get() *GAddr {
	n := len(s.s)
	if n == 0 {
		return nil
	} else {
		addr := s.s[n-1]
		s.s = s.s[:n-1]
		return addr
	}
}

func (s *GStack) popN(n int) {
	l := len(s.s)
	if l < n {
		errFatal("Too few addresses to pop on the stack")
	}
	s.s = s.s[:l-n]
}

func (s *GStack) getN(n int) []*GAddr {
	if len(s.s) < n {
		errFatal("Stack overflow while trying to getN")
	}

	l := len(s.s)
	var res []*GAddr
	for i := 1; i <= n; i++ {
		res = append(res, s.s[l-i])
	}
	s.s = s.s[:l-n]

	return res
}

func (s *GStack) pushNthFromTop(n int) {
	l := len(s.s)
	if l < n+1 {
		errFatal("Trying to push address from index out of stack's bounds")
	} else {
		addr := s.s[l-1-n]
		s.s = append(s.s, addr)
	}
}

func (s *GStack) slide(n int) {
	l := len(s.s)
	if l < n+1 {
		errFatal("Stack too small to perform Slide instruction")
	}
	a := s.s[l-1]
	s.s[l-1-n] = a
	s.s = s.s[:l-n]
}

////////////////////////////////////////// DUMP ////////////////////////////////////////////

type GDump struct {
	d []*GDumpEntry
}

type GDumpEntry struct {
	q *GInstrQueue
	s *GStack
}

func (d *GDump) put(q *GInstrQueue, s *GStack) {
	d.d = append(d.d, &GDumpEntry{q, s})
}

func (d *GDump) get() (*GInstrQueue, *GStack) {
	n := len(d.d)
	if n == 0 {
		return nil, nil
	}

	backup := d.d[n-1]
	d.d = d.d[:n-1]
	return backup.q, backup.s
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

// All the following functions assume that the heap is initialized

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

func (h *GHeap) putNew(node GNode) *GAddr {
	addr := h.newAddr()
	h.put(addr, node)

	return addr
}

/////////////////////////////////////// GLOBAL MAP /////////////////////////////////////////

type GGlobalMap struct {
	m map[string]*GAddr
}

func newGGlobalMap() *GGlobalMap {
	m := new(GGlobalMap)
	m.m = make(map[string]*GAddr)

	return m
}

// All the following functions assume that the map is initialized

func (m *GGlobalMap) get(s string) *GAddr {
	addr, ok := m.m[s]
	if ok {
		return addr
	} else {
		return nil
	}
}

func (m *GGlobalMap) put(s string, addr *GAddr) {
	m.m[s] = addr
}

//////////////////////////////////////// HELPERS ///////////////////////////////////////////

type GAddr struct {
	a int
}

type GCode struct {
	c []GInstr
}

//////////////////////////////////////// NODES /////////////////////////////////////////////

type GNode interface {
	//fmt.Stringer
	gNode()
	getGcRound() int
	setGcRound(int)
}

type BaseGNode struct {
	gcRound int
}

func (*BaseGNode) gNode() {}

func (n *BaseGNode) getGcRound() int {
	return n.gcRound
}

func (n *BaseGNode) setGcRound(round int) {
	n.gcRound = round
}

type NInt struct {
	BaseGNode
	n int
}

type NString struct {
	BaseGNode
	s string
}

type NUnit struct {
	BaseGNode
}

type NApp struct {
	BaseGNode
	fun *GAddr
	arg *GAddr
}

type NGlobal struct {
	BaseGNode
	argsNum int
	code    *GCode
}

type NInd struct {
	BaseGNode
	a *GAddr
}

type NData struct {
	BaseGNode
	tag  int
	args []*GAddr
}

const (
	boolFalseTag = 0
	boolTrueTag  = 1
)

//func (n *NInt) String() string {
//	return "NInt " + strconv.Itoa(n.n)
//}
//
//func (n *NString) String() string {
//	return "NString " + n.s
//}
//
//func (n *NUnit) String() string {
//	return "NUnit"
//}
//
//func (n *NApp) String() string {
//	return "NApp"
//}
//
//func (n *NGlobal) String() string {
//	return "NGlobal " + strconv.Itoa(n.argsNum)
//}
//
//func (n *NInd) String() string {
//	return "NInd"
//}
//
//func (n *NData) String() string {
//	return "NData " + strconv.Itoa(n.tag)
//}

////////////////////////////////////// INSTRUCTIONS ////////////////////////////////////////

type GInstr interface {
	gInstr()
	apply(m *GMachine)
}

type BaseGInstr struct{}

func (*BaseGInstr) gInstr()         {}
func (*BaseGInstr) apply(*GMachine) {}

type IPushInt struct {
	BaseGInstr
	n int
}

func (i *IPushInt) apply(m *GMachine) {
	node := &NInt{n: i.n}
	addr := m.allocNewNode(node)
	m.stack.put(addr)
}

type IPushString struct {
	BaseGInstr
	s string
}

func (i *IPushString) apply(m *GMachine) {
	node := &NString{s: i.s}
	addr := m.allocNewNode(node)
	m.stack.put(addr)
}

type IPushUnit struct {
	BaseGInstr
}

func (i *IPushUnit) apply(m *GMachine) {
	node := &NUnit{}
	addr := m.allocNewNode(node)
	m.stack.put(addr)
}

// TODO: Handle partial application

type IPushGlobal struct {
	BaseGInstr
	f string
}

func (i *IPushGlobal) apply(m *GMachine) {
	addr := m.globalMap.get(i.f)
	if addr == nil {
		errFatal("Function '" + i.f + "' not defined")
	}

	m.stack.put(addr)
}

type IPush struct {
	BaseGInstr
	n int
}

func (i *IPush) apply(m *GMachine) {
	m.stack.pushNthFromTop(i.n)
}

type IMkApp struct {
	BaseGInstr
}

func (i *IMkApp) apply(m *GMachine) {
	funA := m.stack.get()
	argA := m.stack.get()
	if funA == nil || argA == nil {
		errFatal("Not enough values for MkApp on the stack")
	}

	node := &NApp{fun: funA, arg: argA}
	addr := m.allocNewNode(node)
	m.stack.put(addr)
}

type IUnwind struct {
	BaseGInstr
}

func (i *IUnwind) apply(m *GMachine) {
	addr := m.stack.readTop()
	if addr == nil {
		errFatal("No address to unwind on the stack")
	}

	node := m.heap.get(addr)
	if node == nil {
		errFatal("No node to unwind on the heap")
	}

	switch x := node.(type) {
	case *NApp:
		_ = x
		i.applyApp(m)
	case *NGlobal:
		i.applyGlobal(m)
	case *NInd:
		i.applyInd(m)
	case *NInt:
		i.applyReturn(m)
	case *NString:
		i.applyReturn(m)
	case *NUnit:
		i.applyReturn(m)
	case *NData:
		i.applyReturn(m)
	}
}

func (i *IUnwind) applyApp(m *GMachine) {
	addr := m.stack.readTop()
	node := m.heap.get(addr).(*NApp)

	m.stack.put(node.fun)
	m.instrQueue.put(new(IUnwind))
}

// TODO: make sure that partial application works

func (i *IUnwind) applyGlobal(m *GMachine) {
	addr := m.stack.get()
	node := m.heap.get(addr).(*NGlobal)
	n := node.argsNum

	var appAddrs []*GAddr
	var lastAddr *GAddr = nil

	for j := 0; j < n; j++ {
		appAddr := m.stack.get()
		if appAddr == nil {
			errFatal("Not enough parameters passed to a function")
		}

		appNode, ok := m.heap.get(appAddr).(*NApp)
		if appNode == nil || !ok {
			errFatal("Not enough parameters passed to a function")
		}

		appAddrs = append(appAddrs, appNode.arg)
		lastAddr = appAddr
	}

	if n > 0 {
		m.stack.put(lastAddr)
		m.stack.putN(appAddrs)
	} else {
		m.stack.put(addr)
	}

	// TODO: check if cleaning the queue before isn't needed
	m.instrQueue.putN(node.code.c)
}

func (i *IUnwind) applyInd(m *GMachine) {
	addr := m.stack.get()
	ind := m.heap.get(addr).(*NInd)

	m.stack.put(ind.a)
	m.instrQueue.put(new(IUnwind))
}

func (i *IUnwind) applyReturn(m *GMachine) {
	queue, stack := m.dump.get()
	if queue == nil || stack == nil {
		// return to get the whole program result
		return
	}
	addr := m.stack.get()

	m.instrQueue = queue
	m.stack = stack
	m.stack.put(addr)
}

type IUpdate struct {
	BaseGInstr
	n int
}

func (i *IUpdate) apply(m *GMachine) {
	if i.n < 0 {
		return
	}
	addr := m.stack.get()
	addrToReplace := m.stack.readNthFromTop(i.n)
	if addr == nil || addrToReplace == nil {
		errFatal("Address index out of stack's bounds when updating")
	}

	node := &NInd{a: addr}
	m.heap.put(addrToReplace, node)
}

type IPack struct {
	BaseGInstr
	t int
	n int
}

func (i *IPack) apply(m *GMachine) {
	args := m.stack.getN(i.n)
	node := &NData{tag: i.t, args: args}
	addr := m.allocNewNode(node)
	m.stack.put(addr)
}

type ISplit struct {
	BaseGInstr
}

func (i *ISplit) apply(m *GMachine) {
	addr := m.stack.get()
	if addr == nil {
		errFatal("No constructor addr to unpack on the stack")
	}

	node := m.heap.get(addr)
	if node == nil {
		errFatal("No node under given address when looking for a constructor")
	}
	data, ok := node.(*NData)

	if ok {
		m.stack.putN(data.args)
	} else {
		errFatal("The value on stack to unpack is not a constructor")
	}
}

// Handle literals in case patterns

type IJump struct {
	BaseGInstr
	m map[int][]GInstr
}

func (i *IJump) apply(m *GMachine) {
	addr := m.stack.readTop()
	if addr == nil {
		errFatal("Empty stack while trying to apply Jump")
	}

	node := m.heap.get(addr)
	if addr == nil {
		errFatal("Address not on heap while trying to apply Jump")
	}

	data, ok := node.(*NData)
	if !ok {
		errFatal("Not a constructor at the given address while trying to apply Jump")
	}

	code, ok := i.m[data.tag]
	if ok {
		m.instrQueue.putN(code)
	} else {
		// -1 is the default tag
		code, ok := i.m[-1]
		if ok {
			m.instrQueue.putN(code)
		} else {
			errFatal("Missing pattern in case expression")
		}
	}
}

type ISlide struct {
	BaseGInstr
	n int
}

func (i *ISlide) apply(m *GMachine) {
	m.stack.slide(i.n)
}

type IBinOp struct {
	BaseGInstr
	op string
}

func (i *IBinOp) apply(m *GMachine) {
	a0 := m.stack.get()
	a1 := m.stack.get()
	if a0 == nil || a1 == nil {
		errFatal("Empty stack while trying to apply BinOp")
	}

	node0, ok0 := m.heap.get(a0).(*NInt)
	node1, ok1 := m.heap.get(a1).(*NInt)
	if !ok0 || !ok1 {
		errFatal("BinOp arg address not pointing to an NInt")
	}

	n0 := node0.n
	n1 := node1.n

	var res int
	switch i.op {
	case "+":
		res = n0 + n1
	case "-":
		res = n0 - n1
	case "*":
		res = n0 * n1
	case "/":
		res = n0 / n1
	}

	resNode := &NInt{n: res}
	addr := m.allocNewNode(resNode)
	m.stack.put(addr)
}

type ICompOp struct {
	BaseGInstr
	op string
}

func (i *ICompOp) apply(m *GMachine) {
	a0 := m.stack.get()
	a1 := m.stack.get()
	if a0 == nil || a1 == nil {
		errFatal("Empty stack while trying to apply CompOp")
	}

	node0, ok0 := m.heap.get(a0).(*NInt)
	node1, ok1 := m.heap.get(a1).(*NInt)
	if !ok0 || !ok1 {
		errFatal("CompOp arg address not pointing to an NInt")
	}

	n0 := node0.n
	n1 := node1.n

	var res bool
	switch i.op {
	case "<":
		res = n0 < n1
	case ">":
		res = n0 > n1
	case "<=":
		res = n0 <= n1
	case ">=":
		res = n0 >= n1
	case "==":
		res = n0 == n1
	case "/=":
		res = n0 != n1
	}

	resNode := &NData{args: []*GAddr{}}
	if res {
		resNode.tag = boolTrueTag
	} else {
		resNode.tag = boolFalseTag
	}

	addr := m.allocNewNode(resNode)
	m.stack.put(addr)
}

type ILogOp struct {
	BaseGInstr
	op string
}

func (i *ILogOp) apply(m *GMachine) {
	a0 := m.stack.get()
	a1 := m.stack.get()
	if a0 == nil || a1 == nil {
		errFatal("Empty stack while trying to apply LogOp")
	}

	node0, ok0 := m.heap.get(a0).(*NData)
	node1, ok1 := m.heap.get(a1).(*NData)
	if !ok0 || !ok1 {
		errFatal("LogOp arg address not pointing to an NData")
	}

	var n0, n1 bool
	if node0.tag == boolTrueTag {
		n0 = true
	}
	if node1.tag == boolTrueTag {
		n1 = true
	}

	var res bool
	switch i.op {
	case "&&":
		res = n0 && n1
	case "||":
		res = n0 || n1
	}

	resNode := &NData{args: []*GAddr{}}
	if res {
		resNode.tag = boolTrueTag
	} else {
		resNode.tag = boolFalseTag
	}

	addr := m.allocNewNode(resNode)
	m.stack.put(addr)
}

type IEval struct {
	BaseGInstr
}

func (i *IEval) apply(m *GMachine) {
	addr := m.stack.get()
	if addr == nil {
		errFatal("No address to eval on stack")
	}

	m.dump.put(m.instrQueue, m.stack)
	m.instrQueue = new(GInstrQueue)
	m.stack = new(GStack)

	m.instrQueue.put(new(IUnwind))
	m.stack.put(addr)
}

type IAlloc struct {
	BaseGInstr
	n int
}

func (i *IAlloc) apply(m *GMachine) {
	for j := 0; j < i.n; j++ {
		node := &NInd{a: nil}
		addr := m.allocNewNode(node)
		m.stack.put(addr)
	}
}

type IPop struct {
	BaseGInstr
	n int
}

func (i *IPop) apply(m *GMachine) {
	m.stack.popN(i.n)
}

type IIOFun struct {
	BaseGInstr
	fun string
}

func (i *IIOFun) apply(m *GMachine) {
	switch i.fun {
	case "return":
		applyReturn()
	case "putStr":
		applyPutStr(m)
	case "putInt":
		applyPutInt(m)
		//case "getLine":
		//	applyGetLine(m)
	case "tinySleep":
		applyTinySleep(m)
	case "tinyLED":
		applyTinyLED(m)
	case "tinyConfigure":
		applyTinyConfigure(m)
	case "tinyPinOutput":
		applyTinyPinOutput(m)
	case "tinyLow":
		applyTinyLow(m)
	case "tinyHigh":
		applyTinyHigh(m)
	}
}
