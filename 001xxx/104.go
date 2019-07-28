import "math"

func pathInZigZagTree(label int) []int {
	d := int(math.Log2(float64(label))) + 1
	o := make([]int, d)
	for 0 != label {
		o[d-1] = label
		if d > 1 {
			label = (2 << uint(d-1)) + (2 << uint(d-2)) - label - 1
		} else {
			label = 2 - label
		}
		label >>= 1
		d--
	}
	return o
}
