type CustomStack struct {
	m  int
	ar []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{maxSize, make([]int, 0, maxSize)}
}

func (this *CustomStack) Push(x int) {
	if len(this.ar) != this.m {
		this.ar = append(this.ar, x)
	}
}

func (this *CustomStack) Pop() int {
	ar := this.ar
	l := len(ar)
	if 0 != l {
		o := ar[l-1]
		this.ar = ar[:l-1]
		return o
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	ar := this.ar
	l := len(ar)
	for i := 0; i < k && i < l; i++ {
		ar[i] += val
	}
}
