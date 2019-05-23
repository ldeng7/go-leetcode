type Stack interface {
	Empty() bool
	Top() (int, bool)
	Push(item int)
	Pop() (int, bool)
}

type ArrayStack struct {
	arr []int
}

func NewArrayStack() Stack {
	return &ArrayStack{[]int{}}
}

func (s *ArrayStack) Empty() bool {
	return len(s.arr) == 0
}

func (s *ArrayStack) Top() (int, bool) {
	if len(s.arr) != 0 {
		return s.arr[len(s.arr)-1], true
	}
	return 0, false
}

func (s *ArrayStack) Push(item int) {
	s.arr = append(s.arr, item)
}

func (s *ArrayStack) Pop() (int, bool) {
	if len(s.arr) != 0 {
		item := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return item, true
	}
	return 0, false
}

type MyQueue struct {
	s, st Stack
}

func Constructor() MyQueue {
	return MyQueue{NewArrayStack(), NewArrayStack()}
}

func (this *MyQueue) Push(x int) {
	this.s.Push(x)
}

func (this *MyQueue) Pop() int {
	x := -1
	for !this.s.Empty() {
		x, _ = this.s.Pop()
		if !this.s.Empty() {
			this.st.Push(x)
		}
	}
	for !this.st.Empty() {
		y, _ := this.st.Pop()
		this.s.Push(y)
	}
	return x
}

func (this *MyQueue) Peek() int {
	x := -1
	for !this.s.Empty() {
		x, _ = this.s.Pop()
		this.st.Push(x)
	}
	for !this.st.Empty() {
		y, _ := this.st.Pop()
		this.s.Push(y)
	}
	return x
}

func (this *MyQueue) Empty() bool {
	return this.s.Empty()
}
