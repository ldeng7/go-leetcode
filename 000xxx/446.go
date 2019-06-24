import "math"

func numberOfArithmeticSlices(A []int) int {
	o, l := 0, len(A)
	t := make([]map[int]int, l)
	for i := 0; i < l; i++ {
		t[i] = map[int]int{}
		mi := t[i]
		for j := 0; j < i; j++ {
			d, mj := A[i]-A[j], t[j]
			if d > math.MaxInt32 || d < math.MinInt32 {
				continue
			}
			mi[d]++
			if _, ok := mj[d]; ok {
				o += mj[d]
				mi[d] += mj[d]
			}
		}
	}
	return o
}
