type MyCircularQueue struct {
	l, c, h, t int
	arr        []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{k, 0, k - 1, 0, make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.c != this.l {
		this.arr[this.t] = value
		this.t++
		if this.t == this.l {
			this.t = 0
		}
		this.c++
		return true
	}
	return false
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.c != 0 {
		this.h++
		if this.h == this.l {
			this.h = 0
		}
		this.c--
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if this.c != 0 {
		i := this.h + 1
		if i == this.l {
			i = 0
		}
		return this.arr[i]
	}
	return -1
}

func (this *MyCircularQueue) Rear() int {
	if this.c != 0 {
		i := this.t - 1
		if i == -1 {
			i = this.l - 1
		}
		return this.arr[i]
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.c == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.c == this.l
}
