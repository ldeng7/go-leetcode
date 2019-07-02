type queueElemType = int

type ArrayQueue struct {
	arr []queueElemType
	i   int
}

func (q *ArrayQueue) Init() *ArrayQueue {
	q.arr = []queueElemType{}
	return q
}

func (q *ArrayQueue) Empty() bool {
	return len(q.arr)-q.i == 0
}

func (q *ArrayQueue) Top() *queueElemType {
	if len(q.arr)-q.i != 0 {
		return &q.arr[q.i]
	}
	return nil
}

func (q *ArrayQueue) Push(item queueElemType) {
	if len(q.arr) <= 32 || q.i <= (len(q.arr)>>1) {
		q.arr = append(q.arr, item)
	} else {
		arr := make([]queueElemType, len(q.arr)-q.i+1)
		copy(arr, q.arr[q.i:])
		arr[len(arr)-1] = item
		q.arr = arr
		q.i = 0
	}
}

func (q *ArrayQueue) Pop() *queueElemType {
	if len(q.arr)-q.i != 0 {
		e := q.arr[q.i]
		q.i++
		return &e
	}
	return nil
}

type MyStack struct {
	q, qt ArrayQueue
}

func Constructor() MyStack {
	this := &MyStack{}
	this.q.Init()
	this.qt.Init()
	return *this
}

func (this *MyStack) Push(x int) {
	this.q.Push(x)
}

func (this *MyStack) Pop() int {
	o := -1
	for !this.q.Empty() {
		x := this.q.Pop()
		o = *x
		if !this.q.Empty() {
			this.qt.Push(o)
		}
	}
	for !this.qt.Empty() {
		x := this.qt.Pop()
		this.q.Push(*x)
	}
	return o
}

func (this *MyStack) Top() int {
	o := -1
	for !this.q.Empty() {
		x := this.q.Pop()
		o = *x
		this.qt.Push(o)
	}
	for !this.qt.Empty() {
		x := this.qt.Pop()
		this.q.Push(*x)
	}
	return o
}

func (this *MyStack) Empty() bool {
	return this.q.Empty()
}
