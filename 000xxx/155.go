type MinStack struct {
	buf []int
	min int
}

func Constructor() MinStack {
	s := MinStack{
		buf: []int{},
	}
	return s
}

func (this *MinStack) Push(x int) {
	this.buf = append(this.buf, x)
	if 1 == len(this.buf) || x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if 0 == len(this.buf) {
		return
	}
	top := this.buf[len(this.buf)-1]
	this.buf = this.buf[:len(this.buf)-1]
	if this.min == top {
		if 0 == len(this.buf) {
			this.min = 0
		} else {
			this.min = this.buf[0]
			for i := 1; i < len(this.buf); i++ {
				if this.buf[i] < this.min {
					this.min = this.buf[i]
				}
			}
		}
	}

}

func (this *MinStack) Top() int {
	if 0 == len(this.buf) {
		return 0
	}
	return this.buf[len(this.buf)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
