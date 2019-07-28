import "math/rand"

type Solution struct {
	rects [][]int
}

func Constructor(rects [][]int) Solution {
	return Solution{rects}
}

func (this *Solution) Pick() []int {
	sum := 0
	var rect []int
	for _, r := range this.rects {
		a := (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		sum += a
		if rand.Intn(sum) < a {
			rect = r
		}
	}
	y := rand.Intn(rect[2]-rect[0]+1) + rect[0]
	x := rand.Intn(rect[3]-rect[1]+1) + rect[1]
	return []int{y, x}
}
