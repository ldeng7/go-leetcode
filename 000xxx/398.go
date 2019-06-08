import "math/rand"

type Solution struct {
	ar []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	c, o := 0, -1
	for i, e := range this.ar {
		if e != target {
			continue
		}
		c++
		if rand.Int()%c == 0 {
			o = i
		}
	}
	return o
}
