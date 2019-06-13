type FreqStack struct {
	mf  map[int]int
	ms  map[int][]int
	max int
}

func Constructor() FreqStack {
	return FreqStack{map[int]int{}, map[int][]int{}, 0}
}

func (this *FreqStack) Push(x int) {
	f := this.mf[x] + 1
	this.mf[x] = f
	if f > this.max {
		this.max = f
	}
	if s, ok := this.ms[f]; ok {
		this.ms[f] = append(s, x)
	} else {
		this.ms[f] = []int{x}
	}
}

func (this *FreqStack) Pop() int {
	s := this.ms[this.max]
	x := s[len(s)-1]
	s = s[:len(s)-1]
	this.ms[this.max] = s
	this.mf[x]--
	if 0 == len(s) {
		this.max--
	}
	return x
}
