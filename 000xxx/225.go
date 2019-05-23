type Queue interface {
	Empty() bool
	Top() (int, bool)
	Push(item int)
	Pop() (int, bool)
}

type ArrayQueue struct {
	arr []int
	i   int
}

func NewArrayQueue() Queue {
	return &ArrayQueue{[]int{}, 0}
}

func (q *ArrayQueue) Empty() bool {
	return len(q.arr)-q.i == 0
}

func (q *ArrayQueue) Top() (int, bool) {
	if len(q.arr)-q.i != 0 {
		return q.arr[q.i], true
	}
	return 0, false
}

func (q *ArrayQueue) Push(item int) {
	if len(q.arr) <= 32 || q.i <= (len(q.arr)>>1) {
		q.arr = append(q.arr, item)
	} else {
		arr := make([]int, len(q.arr)-q.i+1)
		copy(arr, q.arr[q.i:])
		arr[len(arr)-1] = item
		q.arr = arr
		q.i = 0
	}
}

func (q *ArrayQueue) Pop() (int, bool) {
	if len(q.arr)-q.i != 0 {
		item := q.arr[q.i]
		q.i++
		return item, true
	}
	return 0, false
}

type MyStack struct {
	q, qt Queue
}

func Constructor() MyStack {
	return MyStack{NewArrayQueue(), NewArrayQueue()}
}

func (this *MyStack) Push(x int) {
	this.q.Push(x)
}

func (this *MyStack) Pop() int {
	x := -1
	for !this.q.Empty() {
		x, _ = this.q.Pop()
		if !this.q.Empty() {
			this.qt.Push(x)
		}
	}
	for !this.qt.Empty() {
		y, _ := this.qt.Pop()
		this.q.Push(y)
	}
	return x
}

func (this *MyStack) Top() int {
	x := -1
	for !this.q.Empty() {
		x, _ = this.q.Pop()
		this.qt.Push(x)
	}
	for !this.qt.Empty() {
		y, _ := this.qt.Pop()
		this.q.Push(y)
	}
	return x
}

func (this *MyStack) Empty() bool {
	return this.q.Empty()
}
