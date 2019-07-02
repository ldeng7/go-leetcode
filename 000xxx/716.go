type MaxStack struct {
	s, sm []int
}

func Constructor() MaxStack {
	return MaxStack{make([]int, 0, 32), make([]int, 0, 32)}
}

func (this *MaxStack) Push(x int) {
	this.s = append(this.s, x)
	if em := len(this.sm) - 1; em < 0 || this.sm[em] <= x {
		this.sm = append(this.sm, x)
	}
}

func (this *MaxStack) Pop() int {
	if e := len(this.s) - 1; e >= 0 {
		t := this.s[e]
		this.s = this.s[:e]
		if em := len(this.sm) - 1; this.sm[em] == t {
			this.sm = this.sm[:em]
		}
		return t
	}
	return 0
}

func (this *MaxStack) Top() int {
	if e := len(this.s) - 1; e >= 0 {
		return this.s[e]
	}
	return 0
}

func (this *MaxStack) PeekMax() int {
	if em := len(this.sm) - 1; em >= 0 {
		return this.sm[em]
	}
	return 0
}

func (this *MaxStack) PopMax() int {
	em := len(this.sm) - 1
	if em < 0 {
		return 0
	}
	tm, s := this.sm[em], []int{}
	for {
		e := len(this.s) - 1
		t := this.s[e]
		if t == tm {
			break
		}
		s = append(s, t)
		this.s = this.s[:e]
	}
	if e := len(this.s) - 1; e >= 0 {
		this.s = this.s[:e]
	}
	if em := len(this.sm) - 1; em >= 0 {
		this.sm = this.sm[:em]
	}
	for i := len(s) - 1; i >= 0; i-- {
		this.Push(s[i])
	}
	return tm
}
