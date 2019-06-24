type RLEIterator struct {
	ar []int
	i  int
}

func Constructor(A []int) RLEIterator {
	return RLEIterator{A, 0}
}

func (this *RLEIterator) Next(n int) int {
	l := len(this.ar)
	for ; this.i < l; this.i += 2 {
		if this.ar[this.i]-n >= 0 {
			this.ar[this.i] -= n
			return this.ar[this.i+1]
		} else {
			n -= this.ar[this.i]
		}
	}
	return -1
}
