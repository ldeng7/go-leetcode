import "math/rand"

type Solution struct {
	ar []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.ar
}

func (this *Solution) Shuffle() []int {
	l := len(this.ar)
	o := make([]int, l)
	copy(o, this.ar)
	for i := 0; i < l; i++ {
		j := i + rand.Int()%(l-i)
		o[i], o[j] = o[j], o[i]
	}
	return o
}
