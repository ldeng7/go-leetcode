type PhoneDirectory struct {
	n, i, ir int
	rec      []int
	state    []int
}

func Constructor(maxNumbers int) PhoneDirectory {
	return PhoneDirectory{maxNumbers, 0, -1, make([]int, maxNumbers), make([]int, maxNumbers)}
}

func (this *PhoneDirectory) Get() int {
	out := -1
	if this.i < this.n {
		out = this.i
		this.i++
	} else if this.ir >= 0 {
		out = this.rec[this.ir]
		this.ir--
	}
	if out != -1 {
		this.state[out] = 1
	}
	return out
}

func (this *PhoneDirectory) Check(number int) bool {
	return this.state[number] == 0
}

func (this *PhoneDirectory) Release(number int) {
	if this.state[number] == 0 {
		return
	}
	this.ir++
	this.rec[this.ir] = number
	this.state[number] = 0
}
