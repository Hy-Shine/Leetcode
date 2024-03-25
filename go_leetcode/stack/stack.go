package stack

type nodeStack struct {
	value any
	next  *nodeStack
}

func NewStack() *nodeStack {
	return &nodeStack{}
}

func (node *nodeStack) Push(val any) {
	node.next = &nodeStack{value: val, next: node.next}
}

func (node *nodeStack) Pop() (val any) {
	if node.next != nil {
		val = node.next.value
		node.next = node.next.next
	}
	return
}

func (node *nodeStack) Value() any {
	return node.next.value
}

func (node *nodeStack) IsEmpty() bool {
	return node.next == nil
}

// ----------------------------------------------------------

type stackS[K comparable] struct {
	list []K
}

func NewStackS[K comparable]() *stackS[K] {
	return &stackS[K]{list: make([]K, 0)}
}

func (ss *stackS[K]) IsEmpty() bool {
	return len(ss.list) == 0
}

func (ss *stackS[K]) Pop() K {
	if ss.list == nil {
		ss.list = make([]K, 0)
	}

	var val K
	if !ss.IsEmpty() {
		val = ss.list[len(ss.list)-1]
		ss.list = ss.list[:len(ss.list)-1]
	}
	return val
}

func (ss *stackS[K]) Push(val K) {
	if ss.list == nil {
		ss.list = make([]K, 0)
	}
	ss.list = append(ss.list, val)
}

func (ss *stackS[K]) Value() (K, bool) {
	if !ss.IsEmpty() {
		return ss.list[len(ss.list)-1], true
	}

	var k K
	return k, false
}
