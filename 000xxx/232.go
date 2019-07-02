type stackElemType = int

type ArrayStack struct {
	arr []stackElemType
}

func (s *ArrayStack) Init() *ArrayStack {
	s.arr = []stackElemType{}
	return s
}

func (s *ArrayStack) Empty() bool {
	return len(s.arr) == 0
}

func (s *ArrayStack) Top() *stackElemType {
	if e := len(s.arr) - 1; e >= 0 {
		return &s.arr[e]
	}
	return nil
}

func (s *ArrayStack) Push(item stackElemType) {
	s.arr = append(s.arr, item)
}

func (s *ArrayStack) Pop() *stackElemType {
	if e := len(s.arr) - 1; e >= 0 {
		t := s.arr[e]
		s.arr = s.arr[:e]
		return &t
	}
	return nil
}

type MyQueue struct {
	s, st ArrayStack
}

func Constructor() MyQueue {
	this := &MyQueue{}
	this.s.Init()
	this.st.Init()
	return *this
}

func (this *MyQueue) Push(x int) {
	this.s.Push(x)
}

func (this *MyQueue) Pop() int {
	o := -1
	for !this.s.Empty() {
		x := this.s.Pop()
		o = *x
		if !this.s.Empty() {
			this.st.Push(o)
		}
	}
	for !this.st.Empty() {
		x := this.st.Pop()
		this.s.Push(*x)
	}
	return o
}

func (this *MyQueue) Peek() int {
	o := -1
	for !this.s.Empty() {
		x := this.s.Pop()
		o = *x
		this.st.Push(o)
	}
	for !this.st.Empty() {
		x := this.st.Pop()
		this.s.Push(*x)
	}
	return o
}

func (this *MyQueue) Empty() bool {
	return this.s.Empty()
}
