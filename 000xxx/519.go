import "math/rand"

type Solution struct {
	nr, nc, l int
	m         map[int]int
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{n_rows, n_cols, n_rows * n_cols, map[int]int{}}
}

func (this *Solution) Flip() []int {
	r := rand.Intn(this.l)
	i := r
	this.l--
	if j, ok := this.m[r]; ok {
		r = j
	}
	if j, ok := this.m[this.l]; ok {
		this.m[i] = j
	} else {
		this.m[i] = this.l
	}
	return []int{r / this.nc, r % this.nc}
}

func (this *Solution) Reset() {
	this.l, this.m = this.nr*this.nc, map[int]int{}
}
