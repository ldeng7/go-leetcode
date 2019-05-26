type ListNode struct {
	Val  int
	prev *ListNode
	next *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

func (l *List) PushBack(n *ListNode) {
	n.next = nil
	if l.tail != nil {
		l.tail.next = n
	} else {
		l.head = n
	}
	n.prev = l.tail
	l.tail = n
}

func (l *List) PopBack() *ListNode {
	if l.tail == nil {
		return nil
	}
	r := l.tail
	if r.prev != nil {
		l.tail = r.prev
	} else {
		l.head, l.tail = nil, nil
	}
	return r
}

type Stack interface {
	Empty() bool
	Top() (int, bool)
	Push(item int)
	Pop() (int, bool)
}

type ListStack struct {
	l *List
}

func NewListStack() Stack {
	return &ListStack{&List{}}
}

func (s *ListStack) Empty() bool {
	return nil == s.l.head
}

func (s *ListStack) Top() (int, bool) {
	if nil != s.l.tail {
		return s.l.tail.Val, true
	}
	return 0, false
}

func (s *ListStack) Push(item int) {
	s.l.PushBack(&ListNode{Val: item})
}

func (s *ListStack) Pop() (int, bool) {
	tail := s.l.PopBack()
	if nil != tail {
		return tail.Val, true
	}
	return 0, false
}

type MaxStack struct {
	s, sm Stack
}

func Constructor() MaxStack {
	return MaxStack{NewListStack(), NewListStack()}
}

func (this *MaxStack) Push(x int) {
	this.s.Push(x)
	if t, ok := this.sm.Top(); !ok || t <= x {
		this.sm.Push(x)
	}
}

func (this *MaxStack) Pop() int {
	t, ok := this.s.Pop()
	if ok {
		if tm, _ := this.sm.Top(); tm == t {
			this.sm.Pop()
		}
	}
	return t
}

func (this *MaxStack) Top() int {
	t, _ := this.s.Top()
	return t
}

func (this *MaxStack) PeekMax() int {
	t, _ := this.sm.Top()
	return t
}

func (this *MaxStack) PopMax() int {
	t, ok := this.sm.Top()
	if !ok {
		return 0
	}
	st := NewListStack()
	for {
		t, _ := this.s.Top()
		tm, _ := this.sm.Top()
		if t == tm {
			break
		}
		st.Push(t)
		this.s.Pop()
	}
	this.s.Pop()
	this.sm.Pop()
	for !st.Empty() {
		t, _ := st.Top()
		this.Push(t)
		st.Pop()
	}
	return t
}
