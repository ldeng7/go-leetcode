type FrontMiddleBackQueue struct {
	q []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{make([]int, 0, 16)}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	q := make([]int, len(this.q)+1)
	q[0] = val
	copy(q[1:], this.q)
	this.q = q
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if l := len(this.q); 0 != l {
		m := l >> 1
		q := append(this.q, 0)
		copy(q[m+1:], q[m:])
		q[m] = val
		this.q = q
	} else {
		this.q = append(this.q, val)
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.q = append(this.q, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if 0 != len(this.q) {
		val := this.q[0]
		this.q = this.q[1:]
		return val
	}
	return -1
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if l := len(this.q); 0 != l {
		m := (l - 1) >> 1
		q := this.q
		val := q[m]
		copy(q[m:], q[m+1:])
		this.q = q[:l-1]
		return val
	}
	return -1
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if l := len(this.q); 0 != l {
		val := this.q[l-1]
		this.q = this.q[:l-1]
		return val
	}
	return -1
}
