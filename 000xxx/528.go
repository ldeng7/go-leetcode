import (
	"math/rand"
	"sort"
)

type Solution struct {
	sums []int
	m    int
}

func Constructor(w []int) Solution {
	sums := make([]int, len(w))
	sums[0] = w[0]
	for i := 1; i < len(w); i++ {
		sums[i] = sums[i-1] + w[i]
	}
	return Solution{sums, sums[len(w)-1]}
}

func (this *Solution) PickIndex() int {
	s := rand.Intn(this.m)
	return sort.Search(len(this.sums), func(i int) bool {
		return this.sums[i] > s
	})
}
