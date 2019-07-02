type MinStack struct {
	s, sm []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0, 32), make([]int, 0, 32)}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	if em := len(this.sm) - 1; em < 0 || this.sm[em] >= x {
		this.sm = append(this.sm, x)
	}
}

func (this *MinStack) Pop() {
	if e := len(this.s) - 1; e >= 0 {
		t := this.s[e]
		this.s = this.s[:e]
		if em := len(this.sm) - 1; this.sm[em] == t {
			this.sm = this.sm[:em]
		}
	}
}

func (this *MinStack) Top() int {
	if e := len(this.s) - 1; e >= 0 {
		return this.s[e]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if em := len(this.sm) - 1; em >= 0 {
		return this.sm[em]
	}
	return 0
}
