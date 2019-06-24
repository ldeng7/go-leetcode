import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxSubarraySumCircular(A []int) int {
	ma := math.MinInt64
	for _, a := range A {
		ma = max(ma, a)
	}
	if ma <= 0 {
		return ma
	}
	mi, mi1 := math.MaxInt32, math.MaxInt32
	ma, ma1 := math.MinInt32, math.MinInt32
	s := 0
	for _, a := range A {
		s += a
		mi1 = min(a, mi1+a)
		mi = min(mi, mi1)
		ma1 = max(a, ma1+a)
		ma = max(ma, ma1)
	}
	return max(ma, s-mi)
}
