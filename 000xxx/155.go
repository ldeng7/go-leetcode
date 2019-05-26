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

type MinStack struct {
	s, sm Stack
}

func Constructor() MinStack {
	return MinStack{NewListStack(), NewListStack()}
}

func (this *MinStack) Push(x int) {
	this.s.Push(x)
	if t, ok := this.sm.Top(); !ok || t >= x {
		this.sm.Push(x)
	}
}

func (this *MinStack) Pop() {
	if t, ok := this.s.Pop(); ok {
		if tm, _ := this.sm.Top(); tm == t {
			this.sm.Pop()
		}
	}
}

func (this *MinStack) Top() int {
	t, _ := this.s.Top()
	return t
}

func (this *MinStack) GetMin() int {
	t, _ := this.sm.Top()
	return t
}
