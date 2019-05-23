type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

func (l *List) PushBack(n *ListNode) {
	n.Next = nil
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	n.Prev = l.Tail
	l.Tail = n
}

func (l *List) PopBack() *ListNode {
	if l.Tail == nil {
		return nil
	}
	r := l.Tail
	if r.Prev != nil {
		l.Tail = r.Prev
	} else {
		l.Head, l.Tail = nil, nil
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
	return nil == s.l.Head
}

func (s *ListStack) Top() (int, bool) {
	if nil != s.l.Tail {
		return s.l.Tail.Val, true
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
